///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// GetAppsOKCode is the HTTP code returned for type GetAppsOK
const GetAppsOKCode int = 200

/*GetAppsOK Successful operation

swagger:response getAppsOK
*/
type GetAppsOK struct {

	/*
	  In: Body
	*/
	Payload []*v1.Application `json:"body,omitempty"`
}

// NewGetAppsOK creates GetAppsOK with default headers values
func NewGetAppsOK() *GetAppsOK {

	return &GetAppsOK{}
}

// WithPayload adds the payload to the get apps o k response
func (o *GetAppsOK) WithPayload(payload []*v1.Application) *GetAppsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get apps o k response
func (o *GetAppsOK) SetPayload(payload []*v1.Application) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAppsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*v1.Application, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetAppsInternalServerErrorCode is the HTTP code returned for type GetAppsInternalServerError
const GetAppsInternalServerErrorCode int = 500

/*GetAppsInternalServerError Internal Error

swagger:response getAppsInternalServerError
*/
type GetAppsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewGetAppsInternalServerError creates GetAppsInternalServerError with default headers values
func NewGetAppsInternalServerError() *GetAppsInternalServerError {

	return &GetAppsInternalServerError{}
}

// WithPayload adds the payload to the get apps internal server error response
func (o *GetAppsInternalServerError) WithPayload(payload *v1.Error) *GetAppsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get apps internal server error response
func (o *GetAppsInternalServerError) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAppsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetAppsDefault Unexpected Error

swagger:response getAppsDefault
*/
type GetAppsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewGetAppsDefault creates GetAppsDefault with default headers values
func NewGetAppsDefault(code int) *GetAppsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetAppsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get apps default response
func (o *GetAppsDefault) WithStatusCode(code int) *GetAppsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get apps default response
func (o *GetAppsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get apps default response
func (o *GetAppsDefault) WithPayload(payload *v1.Error) *GetAppsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get apps default response
func (o *GetAppsDefault) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAppsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
