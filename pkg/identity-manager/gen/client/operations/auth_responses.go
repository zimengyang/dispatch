///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// AuthReader is a Reader for the Auth structure.
type AuthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewAuthAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewAuthUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewAuthForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewAuthDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAuthAccepted creates a AuthAccepted with default headers values
func NewAuthAccepted() *AuthAccepted {
	return &AuthAccepted{}
}

/*AuthAccepted handles this case with default header values.

default response if authorized
*/
type AuthAccepted struct {
	Payload *v1.Message
}

func (o *AuthAccepted) Error() string {
	return fmt.Sprintf("[GET /v1/iam/auth][%d] authAccepted  %+v", 202, o.Payload)
}

func (o *AuthAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthUnauthorized creates a AuthUnauthorized with default headers values
func NewAuthUnauthorized() *AuthUnauthorized {
	return &AuthUnauthorized{}
}

/*AuthUnauthorized handles this case with default header values.

Unauthorized
*/
type AuthUnauthorized struct {
}

func (o *AuthUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/iam/auth][%d] authUnauthorized ", 401)
}

func (o *AuthUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAuthForbidden creates a AuthForbidden with default headers values
func NewAuthForbidden() *AuthForbidden {
	return &AuthForbidden{}
}

/*AuthForbidden handles this case with default header values.

Forbidden
*/
type AuthForbidden struct {
}

func (o *AuthForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/iam/auth][%d] authForbidden ", 403)
}

func (o *AuthForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAuthDefault creates a AuthDefault with default headers values
func NewAuthDefault(code int) *AuthDefault {
	return &AuthDefault{
		_statusCode: code,
	}
}

/*AuthDefault handles this case with default header values.

error
*/
type AuthDefault struct {
	_statusCode int

	Payload *v1.Error
}

// Code gets the status code for the auth default response
func (o *AuthDefault) Code() int {
	return o._statusCode
}

func (o *AuthDefault) Error() string {
	return fmt.Sprintf("[GET /v1/iam/auth][%d] auth default  %+v", o._statusCode, o.Payload)
}

func (o *AuthDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
