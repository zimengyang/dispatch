///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package drivers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// GetDriverOKCode is the HTTP code returned for type GetDriverOK
const GetDriverOKCode int = 200

/*GetDriverOK Successful operation

swagger:response getDriverOK
*/
type GetDriverOK struct {

	/*
	  In: Body
	*/
	Payload *v1.EventDriver `json:"body,omitempty"`
}

// NewGetDriverOK creates GetDriverOK with default headers values
func NewGetDriverOK() *GetDriverOK {

	return &GetDriverOK{}
}

// WithPayload adds the payload to the get driver o k response
func (o *GetDriverOK) WithPayload(payload *v1.EventDriver) *GetDriverOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get driver o k response
func (o *GetDriverOK) SetPayload(payload *v1.EventDriver) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDriverOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetDriverBadRequestCode is the HTTP code returned for type GetDriverBadRequest
const GetDriverBadRequestCode int = 400

/*GetDriverBadRequest Invalid Name supplied

swagger:response getDriverBadRequest
*/
type GetDriverBadRequest struct {

	/*
	  In: Body
	*/
	Payload *v1.ErrorOAIGen `json:"body,omitempty"`
}

// NewGetDriverBadRequest creates GetDriverBadRequest with default headers values
func NewGetDriverBadRequest() *GetDriverBadRequest {

	return &GetDriverBadRequest{}
}

// WithPayload adds the payload to the get driver bad request response
func (o *GetDriverBadRequest) WithPayload(payload *v1.ErrorOAIGen) *GetDriverBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get driver bad request response
func (o *GetDriverBadRequest) SetPayload(payload *v1.ErrorOAIGen) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDriverBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetDriverNotFoundCode is the HTTP code returned for type GetDriverNotFound
const GetDriverNotFoundCode int = 404

/*GetDriverNotFound Driver not found

swagger:response getDriverNotFound
*/
type GetDriverNotFound struct {

	/*
	  In: Body
	*/
	Payload *v1.ErrorOAIGen `json:"body,omitempty"`
}

// NewGetDriverNotFound creates GetDriverNotFound with default headers values
func NewGetDriverNotFound() *GetDriverNotFound {

	return &GetDriverNotFound{}
}

// WithPayload adds the payload to the get driver not found response
func (o *GetDriverNotFound) WithPayload(payload *v1.ErrorOAIGen) *GetDriverNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get driver not found response
func (o *GetDriverNotFound) SetPayload(payload *v1.ErrorOAIGen) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDriverNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetDriverInternalServerErrorCode is the HTTP code returned for type GetDriverInternalServerError
const GetDriverInternalServerErrorCode int = 500

/*GetDriverInternalServerError Internal server error

swagger:response getDriverInternalServerError
*/
type GetDriverInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *v1.ErrorOAIGen `json:"body,omitempty"`
}

// NewGetDriverInternalServerError creates GetDriverInternalServerError with default headers values
func NewGetDriverInternalServerError() *GetDriverInternalServerError {

	return &GetDriverInternalServerError{}
}

// WithPayload adds the payload to the get driver internal server error response
func (o *GetDriverInternalServerError) WithPayload(payload *v1.ErrorOAIGen) *GetDriverInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get driver internal server error response
func (o *GetDriverInternalServerError) SetPayload(payload *v1.ErrorOAIGen) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDriverInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetDriverDefault Unknown error

swagger:response getDriverDefault
*/
type GetDriverDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *v1.ErrorOAIGen `json:"body,omitempty"`
}

// NewGetDriverDefault creates GetDriverDefault with default headers values
func NewGetDriverDefault(code int) *GetDriverDefault {
	if code <= 0 {
		code = 500
	}

	return &GetDriverDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get driver default response
func (o *GetDriverDefault) WithStatusCode(code int) *GetDriverDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get driver default response
func (o *GetDriverDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get driver default response
func (o *GetDriverDefault) WithPayload(payload *v1.ErrorOAIGen) *GetDriverDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get driver default response
func (o *GetDriverDefault) SetPayload(payload *v1.ErrorOAIGen) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDriverDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
