///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package base_image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// AddBaseImageCreatedCode is the HTTP code returned for type AddBaseImageCreated
const AddBaseImageCreatedCode int = 201

/*AddBaseImageCreated created

swagger:response addBaseImageCreated
*/
type AddBaseImageCreated struct {

	/*
	  In: Body
	*/
	Payload *v1.BaseImage `json:"body,omitempty"`
}

// NewAddBaseImageCreated creates AddBaseImageCreated with default headers values
func NewAddBaseImageCreated() *AddBaseImageCreated {

	return &AddBaseImageCreated{}
}

// WithPayload adds the payload to the add base image created response
func (o *AddBaseImageCreated) WithPayload(payload *v1.BaseImage) *AddBaseImageCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add base image created response
func (o *AddBaseImageCreated) SetPayload(payload *v1.BaseImage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddBaseImageCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddBaseImageBadRequestCode is the HTTP code returned for type AddBaseImageBadRequest
const AddBaseImageBadRequestCode int = 400

/*AddBaseImageBadRequest Invalid input

swagger:response addBaseImageBadRequest
*/
type AddBaseImageBadRequest struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewAddBaseImageBadRequest creates AddBaseImageBadRequest with default headers values
func NewAddBaseImageBadRequest() *AddBaseImageBadRequest {

	return &AddBaseImageBadRequest{}
}

// WithPayload adds the payload to the add base image bad request response
func (o *AddBaseImageBadRequest) WithPayload(payload *v1.Error) *AddBaseImageBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add base image bad request response
func (o *AddBaseImageBadRequest) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddBaseImageBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddBaseImageUnauthorizedCode is the HTTP code returned for type AddBaseImageUnauthorized
const AddBaseImageUnauthorizedCode int = 401

/*AddBaseImageUnauthorized Unauthorized Request

swagger:response addBaseImageUnauthorized
*/
type AddBaseImageUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewAddBaseImageUnauthorized creates AddBaseImageUnauthorized with default headers values
func NewAddBaseImageUnauthorized() *AddBaseImageUnauthorized {

	return &AddBaseImageUnauthorized{}
}

// WithPayload adds the payload to the add base image unauthorized response
func (o *AddBaseImageUnauthorized) WithPayload(payload *v1.Error) *AddBaseImageUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add base image unauthorized response
func (o *AddBaseImageUnauthorized) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddBaseImageUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddBaseImageForbiddenCode is the HTTP code returned for type AddBaseImageForbidden
const AddBaseImageForbiddenCode int = 403

/*AddBaseImageForbidden access to this resource is forbidden

swagger:response addBaseImageForbidden
*/
type AddBaseImageForbidden struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewAddBaseImageForbidden creates AddBaseImageForbidden with default headers values
func NewAddBaseImageForbidden() *AddBaseImageForbidden {

	return &AddBaseImageForbidden{}
}

// WithPayload adds the payload to the add base image forbidden response
func (o *AddBaseImageForbidden) WithPayload(payload *v1.Error) *AddBaseImageForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add base image forbidden response
func (o *AddBaseImageForbidden) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddBaseImageForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddBaseImageConflictCode is the HTTP code returned for type AddBaseImageConflict
const AddBaseImageConflictCode int = 409

/*AddBaseImageConflict Already Exists

swagger:response addBaseImageConflict
*/
type AddBaseImageConflict struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewAddBaseImageConflict creates AddBaseImageConflict with default headers values
func NewAddBaseImageConflict() *AddBaseImageConflict {

	return &AddBaseImageConflict{}
}

// WithPayload adds the payload to the add base image conflict response
func (o *AddBaseImageConflict) WithPayload(payload *v1.Error) *AddBaseImageConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add base image conflict response
func (o *AddBaseImageConflict) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddBaseImageConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*AddBaseImageDefault Generic error response

swagger:response addBaseImageDefault
*/
type AddBaseImageDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewAddBaseImageDefault creates AddBaseImageDefault with default headers values
func NewAddBaseImageDefault(code int) *AddBaseImageDefault {
	if code <= 0 {
		code = 500
	}

	return &AddBaseImageDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the add base image default response
func (o *AddBaseImageDefault) WithStatusCode(code int) *AddBaseImageDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the add base image default response
func (o *AddBaseImageDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the add base image default response
func (o *AddBaseImageDefault) WithPayload(payload *v1.Error) *AddBaseImageDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add base image default response
func (o *AddBaseImageDefault) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddBaseImageDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
