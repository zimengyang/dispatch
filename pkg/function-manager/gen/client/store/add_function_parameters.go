///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// NewAddFunctionParams creates a new AddFunctionParams object
// with the default values initialized.
func NewAddFunctionParams() *AddFunctionParams {
	var ()
	return &AddFunctionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddFunctionParamsWithTimeout creates a new AddFunctionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddFunctionParamsWithTimeout(timeout time.Duration) *AddFunctionParams {
	var ()
	return &AddFunctionParams{

		timeout: timeout,
	}
}

// NewAddFunctionParamsWithContext creates a new AddFunctionParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddFunctionParamsWithContext(ctx context.Context) *AddFunctionParams {
	var ()
	return &AddFunctionParams{

		Context: ctx,
	}
}

// NewAddFunctionParamsWithHTTPClient creates a new AddFunctionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddFunctionParamsWithHTTPClient(client *http.Client) *AddFunctionParams {
	var ()
	return &AddFunctionParams{
		HTTPClient: client,
	}
}

/*AddFunctionParams contains all the parameters to send to the API endpoint
for the add function operation typically these are written to a http.Request
*/
type AddFunctionParams struct {

	/*Body
	  function object

	*/
	Body *v1.Function

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add function params
func (o *AddFunctionParams) WithTimeout(timeout time.Duration) *AddFunctionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add function params
func (o *AddFunctionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add function params
func (o *AddFunctionParams) WithContext(ctx context.Context) *AddFunctionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add function params
func (o *AddFunctionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add function params
func (o *AddFunctionParams) WithHTTPClient(client *http.Client) *AddFunctionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add function params
func (o *AddFunctionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add function params
func (o *AddFunctionParams) WithBody(body *v1.Function) *AddFunctionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add function params
func (o *AddFunctionParams) SetBody(body *v1.Function) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddFunctionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
