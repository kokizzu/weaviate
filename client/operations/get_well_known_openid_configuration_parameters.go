/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/semi-technologies/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@semi.technology
 */ // Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetWellKnownOpenidConfigurationParams creates a new GetWellKnownOpenidConfigurationParams object
// with the default values initialized.
func NewGetWellKnownOpenidConfigurationParams() *GetWellKnownOpenidConfigurationParams {

	return &GetWellKnownOpenidConfigurationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWellKnownOpenidConfigurationParamsWithTimeout creates a new GetWellKnownOpenidConfigurationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWellKnownOpenidConfigurationParamsWithTimeout(timeout time.Duration) *GetWellKnownOpenidConfigurationParams {

	return &GetWellKnownOpenidConfigurationParams{

		timeout: timeout,
	}
}

// NewGetWellKnownOpenidConfigurationParamsWithContext creates a new GetWellKnownOpenidConfigurationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWellKnownOpenidConfigurationParamsWithContext(ctx context.Context) *GetWellKnownOpenidConfigurationParams {

	return &GetWellKnownOpenidConfigurationParams{

		Context: ctx,
	}
}

// NewGetWellKnownOpenidConfigurationParamsWithHTTPClient creates a new GetWellKnownOpenidConfigurationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWellKnownOpenidConfigurationParamsWithHTTPClient(client *http.Client) *GetWellKnownOpenidConfigurationParams {

	return &GetWellKnownOpenidConfigurationParams{
		HTTPClient: client,
	}
}

/*GetWellKnownOpenidConfigurationParams contains all the parameters to send to the API endpoint
for the get well known openid configuration operation typically these are written to a http.Request
*/
type GetWellKnownOpenidConfigurationParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get well known openid configuration params
func (o *GetWellKnownOpenidConfigurationParams) WithTimeout(timeout time.Duration) *GetWellKnownOpenidConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get well known openid configuration params
func (o *GetWellKnownOpenidConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get well known openid configuration params
func (o *GetWellKnownOpenidConfigurationParams) WithContext(ctx context.Context) *GetWellKnownOpenidConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get well known openid configuration params
func (o *GetWellKnownOpenidConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get well known openid configuration params
func (o *GetWellKnownOpenidConfigurationParams) WithHTTPClient(client *http.Client) *GetWellKnownOpenidConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get well known openid configuration params
func (o *GetWellKnownOpenidConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetWellKnownOpenidConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}