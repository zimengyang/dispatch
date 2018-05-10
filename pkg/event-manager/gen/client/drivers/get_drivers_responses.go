///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package drivers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// GetDriversReader is a Reader for the GetDrivers structure.
type GetDriversReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDriversReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDriversOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetDriversInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetDriversDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDriversOK creates a GetDriversOK with default headers values
func NewGetDriversOK() *GetDriversOK {
	return &GetDriversOK{}
}

/*GetDriversOK handles this case with default header values.

Successful operation
*/
type GetDriversOK struct {
	Payload []*v1.EventDriver
}

func (o *GetDriversOK) Error() string {
	return fmt.Sprintf("[GET /drivers][%d] getDriversOK  %+v", 200, o.Payload)
}

func (o *GetDriversOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDriversInternalServerError creates a GetDriversInternalServerError with default headers values
func NewGetDriversInternalServerError() *GetDriversInternalServerError {
	return &GetDriversInternalServerError{}
}

/*GetDriversInternalServerError handles this case with default header values.

Internal server error
*/
type GetDriversInternalServerError struct {
	Payload *v1.ErrorOAIGen
}

func (o *GetDriversInternalServerError) Error() string {
	return fmt.Sprintf("[GET /drivers][%d] getDriversInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDriversInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.ErrorOAIGen)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDriversDefault creates a GetDriversDefault with default headers values
func NewGetDriversDefault(code int) *GetDriversDefault {
	return &GetDriversDefault{
		_statusCode: code,
	}
}

/*GetDriversDefault handles this case with default header values.

Unknown error
*/
type GetDriversDefault struct {
	_statusCode int

	Payload *v1.ErrorOAIGen
}

// Code gets the status code for the get drivers default response
func (o *GetDriversDefault) Code() int {
	return o._statusCode
}

func (o *GetDriversDefault) Error() string {
	return fmt.Sprintf("[GET /drivers][%d] getDrivers default  %+v", o._statusCode, o.Payload)
}

func (o *GetDriversDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.ErrorOAIGen)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
