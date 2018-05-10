///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

package drivers

import (
	"reflect"
	"time"

	ewrapper "github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/vmware/dispatch/pkg/controller"
	"github.com/vmware/dispatch/pkg/entity-store"
	"github.com/vmware/dispatch/pkg/event-manager/drivers/entities"
	"github.com/vmware/dispatch/pkg/trace"
)

// EntityHandler handles driver entity operations
type EntityHandler struct {
	store   entitystore.EntityStore
	backend Backend
}

// NewEntityHandler creates new instance of EntityHandler
func NewEntityHandler(store entitystore.EntityStore, backend Backend) *EntityHandler {
	return &EntityHandler{
		store:   store,
		backend: backend,
	}
}

// Type returns Entity Handler type
func (h *EntityHandler) Type() reflect.Type {
	defer trace.Trace("")()

	return reflect.TypeOf(&entities.Driver{})
}

// Add adds new driver to the store, and executes its deployment.
func (h *EntityHandler) Add(obj entitystore.Entity) (err error) {
	defer trace.Tracef("name %s", obj.GetName())()

	driver := obj.(*entities.Driver)
	defer func() { h.store.Update(driver.GetRevision(), driver) }()

	// deploy the deployment in k8s cluster

	if err := h.backend.Deploy(driver); err != nil {
		translateErrorToEntityState(driver, err)
		return ewrapper.Wrap(err, "error deploying driver")
	}

	driver.Status = entitystore.StatusREADY

	log.Infof("%s-driver %s has been deployed on k8s", driver.Type, driver.Name)

	return nil
}

// Update updates the driver by updating the deployment
func (h *EntityHandler) Update(obj entitystore.Entity) (err error) {
	defer trace.Tracef("name %s", obj.GetName())()

	driver := obj.(*entities.Driver)
	defer func() { h.store.Update(driver.GetRevision(), driver) }()

	if err := h.backend.Update(driver); err != nil {
		translateErrorToEntityState(driver, err)
		return ewrapper.Wrap(err, "error updating driver")
	}

	driver.Status = entitystore.StatusREADY

	log.Infof("%s-driver %s has been updated", driver.Type, driver.Name)

	return nil
}

// Delete deletes the driver from the backend
func (h *EntityHandler) Delete(obj entitystore.Entity) error {
	defer trace.Tracef("name '%s'", obj.GetName())()

	driver := obj.(*entities.Driver)

	// delete the deployment from k8s cluster
	err := h.backend.Delete(driver)
	if err != nil {
		translateErrorToEntityState(driver, err)
		h.store.Update(driver.GetRevision(), driver)
		return ewrapper.Wrap(err, "error deleting driver")
	}

	if err := h.store.Delete(driver.OrganizationID, driver.Name, driver); err != nil {
		return ewrapper.Wrap(err, "store error when deleting driver")
	}
	log.Infof("driver %s deleted from k8s and the entity store", driver.Name)
	return nil
}

// Sync Executes sync loop
func (h *EntityHandler) Sync(organizationID string, resyncPeriod time.Duration) ([]entitystore.Entity, error) {
	defer trace.Trace("")()

	// list entity filter
	filter := entitystore.FilterEverything().Add(
		entitystore.FilterStat{
			Scope:   entitystore.FilterScopeField,
			Subject: "ModifiedTime",
			Verb:    entitystore.FilterVerbBefore,
			Object:  time.Now().Add(-resyncPeriod),
		},
		entitystore.FilterStat{
			Scope:   entitystore.FilterScopeField,
			Subject: "Status",
			Verb:    entitystore.FilterVerbIn,
			Object: []entitystore.Status{
				entitystore.StatusCREATING, entitystore.StatusUPDATING, entitystore.StatusDELETING,
				entitystore.StatusERROR,
				entitystore.StatusREADY,
			},
		})
	syncingEntities, err := controller.DefaultSync(h.store, h.Type(), organizationID, resyncPeriod, filter)
	if err != nil {
		return nil, err
	}

	// sync with k8s, pending entities might change status during sync
	k := h.backend.(*k8sBackend)
	var pendingEntities []entitystore.Entity
	for _, e := range syncingEntities {
		driver := e.(*entities.Driver)

		switch e.GetStatus() {
		case entitystore.StatusREADY:
			// 'real' ready entity will not be processed
			var k8serr error
			fullname := getDriverFullName(driver)
			deployment, err := k.clientset.AppsV1beta1().Deployments(k.config.DriverNamespace).Get(fullname, v1.GetOptions{})

			if err != nil {
				if k8serrors.IsNotFound(err) {
					k8serr = &EventdriverErrorDeploymentNotFound{
						Err: ewrapper.Wrapf(err, "k8s: deployment=%s not found", fullname),
					}
				} else {
					k8serr = &EventdriverErrorUnknown{
						Err: ewrapper.Wrapf(err, "k8s: deployment=%s unexpected error", fullname),
					}
				}
			} else if deployment.Status.AvailableReplicas == 0 {
				k8serr = &EventdriverErrorDeploymentNotAvaialble{
					Err: ewrapper.Errorf("k8s: deployment=%s not available", fullname),
				}
			}

			if k8serr != nil {
				translateErrorToEntityState(driver, k8serr)
				pendingEntities = append(pendingEntities, e)
			}
		default:
			pendingEntities = append(pendingEntities, e)
		}
	}

	return pendingEntities, nil
}

// Error handles error state
func (h *EntityHandler) Error(obj entitystore.Entity) error {
	defer trace.Tracef("")()

	driver := obj.(*entities.Driver)
	log.Debugf("%s error: reason are: %s", driver.GetName(), driver.GetReason())

	recover := false
	var err error
	// try to recover/handle error state
	switch driver.GetReason()[0] {
	case errReasonDeploymentNotFound:
		// TODO
		log.Debug("deployment not found, re-try creating")
		driver.SetStatus(entitystore.StatusCREATING)
		driver.SetReason(entitystore.Reason{})
		h.store.Update(driver.GetRevision(), driver)
		recover = true

	case errReasonDeploymentAlreadyExists:
		// TODO
		log.Debug("deployment already exist, first delete deployment, re-try creating")
		err = h.backend.Delete(driver)
		driver.SetStatus(entitystore.StatusCREATING)
		driver.SetReason(entitystore.Reason{})
		h.store.UpdateWithError(driver, err)
		recover = true

	case errReasonDeploymentNotAvaialble:
		// TODO
		log.Debug("deployment not available")

		k := h.backend.(*k8sBackend)
		deployment, err := k.clientset.AppsV1beta1().Deployments(k.config.DriverNamespace).Get(getDriverFullName(driver), v1.GetOptions{})
		if err != nil && k8serrors.IsNotFound(err) {
			translateErrorToEntityState(driver, &EventdriverErrorDeploymentNotFound{
				Err: ewrapper.Wrapf(err, "k8s: deployment=%s not found", getDriverFullName(driver)),
			})
		}
		if err == nil && deployment.Status.AvailableReplicas > 0 {
			log.Debugf("%s is back from DeploymentNotAvailable", driver.GetName())
			driver.SetStatus(entitystore.StatusREADY)
			driver.SetReason(entitystore.Reason{})
			recover = true
		}

		h.store.Update(driver.GetRevision(), driver)

	default:
		// TODO
		log.Debug("unknown error")
		err = ewrapper.New("unknonw error")
	}

	if recover {
		log.Debugf("%s recovered", driver.Name)
	} else {
		log.Debugf("%s failed to recover", driver.Name)
	}

	return err
}

func translateErrorToEntityState(driver *entities.Driver, e error) {
	if e == nil {
		return
	}
	log.Debugf("put driver to error state %s", driver.GetName())
	reason := []string{
		ErrorTypeToErrorReason[reflect.TypeOf(e)],
		e.Error(),
	}
	driver.SetReason(reason)
	driver.SetStatus(entitystore.StatusERROR)
}
