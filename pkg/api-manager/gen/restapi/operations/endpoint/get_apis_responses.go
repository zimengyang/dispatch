///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// GetApisOKCode is the HTTP code returned for type GetApisOK
const GetApisOKCode int = 200

/*GetApisOK Successful operation

swagger:response getApisOK
*/
type GetApisOK struct {

	/*
	  In: Body
	*/
	Payload []*v1.API `json:"body,omitempty"`
}

// NewGetApisOK creates GetApisOK with default headers values
func NewGetApisOK() *GetApisOK {

	return &GetApisOK{}
}

// WithPayload adds the payload to the get apis o k response
func (o *GetApisOK) WithPayload(payload []*v1.API) *GetApisOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get apis o k response
func (o *GetApisOK) SetPayload(payload []*v1.API) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetApisOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*v1.API, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetApisInternalServerErrorCode is the HTTP code returned for type GetApisInternalServerError
const GetApisInternalServerErrorCode int = 500

/*GetApisInternalServerError Internal Error

swagger:response getApisInternalServerError
*/
type GetApisInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewGetApisInternalServerError creates GetApisInternalServerError with default headers values
func NewGetApisInternalServerError() *GetApisInternalServerError {

	return &GetApisInternalServerError{}
}

// WithPayload adds the payload to the get apis internal server error response
func (o *GetApisInternalServerError) WithPayload(payload *v1.Error) *GetApisInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get apis internal server error response
func (o *GetApisInternalServerError) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetApisInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetApisDefault Unexpected Error

swagger:response getApisDefault
*/
type GetApisDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewGetApisDefault creates GetApisDefault with default headers values
func NewGetApisDefault(code int) *GetApisDefault {
	if code <= 0 {
		code = 500
	}

	return &GetApisDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get apis default response
func (o *GetApisDefault) WithStatusCode(code int) *GetApisDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get apis default response
func (o *GetApisDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get apis default response
func (o *GetApisDefault) WithPayload(payload *v1.Error) *GetApisDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get apis default response
func (o *GetApisDefault) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetApisDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
