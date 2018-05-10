///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

package drivers

import "reflect"

const (
	errReasonDeploymentNotFound      = "DeploymentNotFound"
	errReasonDeploymentAlreadyExists = "DeploymentAlreadyExists"
	errReasonDeploymentNotAvaialble  = "DeploymentNotAvailable"
	errReasonUnknown                 = "Unknown"
)

var (
	// ErrorTypeToErrorReason maps error type to error reason
	ErrorTypeToErrorReason = map[reflect.Type]string{
		reflect.TypeOf(&EventdriverErrorDeploymentNotFound{}):      errReasonDeploymentNotFound,
		reflect.TypeOf(&EventdriverErrorDeploymentNotAvaialble{}):  errReasonDeploymentNotAvaialble,
		reflect.TypeOf(&EventdriverErrorDeploymentAlreadyExists{}): errReasonDeploymentAlreadyExists,
		reflect.TypeOf(&EventdriverErrorUnknown{}):                 errReasonUnknown,
	}
)

// EventdriverErrorDeploymentNotFound defines deployment not found error for event driver
type EventdriverErrorDeploymentNotFound struct {
	Err error `json:"err"`
}

func (err *EventdriverErrorDeploymentNotFound) Error() string {
	return err.Err.Error()
}

// EventdriverErrorDeploymentNotAvaialble defines image pull error for event driver
type EventdriverErrorDeploymentNotAvaialble struct {
	Err error `json:"err"`
}

func (err *EventdriverErrorDeploymentNotAvaialble) Error() string {
	return err.Err.Error()
}

// EventdriverErrorDeploymentAlreadyExists defines deployment already exists error for event driver
type EventdriverErrorDeploymentAlreadyExists struct {
	Err error `json:"err"`
}

func (err *EventdriverErrorDeploymentAlreadyExists) Error() string {
	return err.Err.Error()
}

// EventdriverErrorUnknown defines unknonwn error for event driver
type EventdriverErrorUnknown struct {
	Err error `json:"err"`
}

func (err *EventdriverErrorUnknown) Error() string {
	return err.Err.Error()
}
