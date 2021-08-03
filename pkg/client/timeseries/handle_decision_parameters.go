// Code generated by go-swagger; DO NOT EDIT.

package timeseries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewHandleDecisionParams creates a new HandleDecisionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHandleDecisionParams() *HandleDecisionParams {
	return &HandleDecisionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHandleDecisionParamsWithTimeout creates a new HandleDecisionParams object
// with the ability to set a timeout on a request.
func NewHandleDecisionParamsWithTimeout(timeout time.Duration) *HandleDecisionParams {
	return &HandleDecisionParams{
		timeout: timeout,
	}
}

// NewHandleDecisionParamsWithContext creates a new HandleDecisionParams object
// with the ability to set a context for a request.
func NewHandleDecisionParamsWithContext(ctx context.Context) *HandleDecisionParams {
	return &HandleDecisionParams{
		Context: ctx,
	}
}

// NewHandleDecisionParamsWithHTTPClient creates a new HandleDecisionParams object
// with the ability to set a custom HTTPClient for a request.
func NewHandleDecisionParamsWithHTTPClient(client *http.Client) *HandleDecisionParams {
	return &HandleDecisionParams{
		HTTPClient: client,
	}
}

/* HandleDecisionParams contains all the parameters to send to the API endpoint
   for the handle decision operation.

   Typically these are written to a http.Request.
*/
type HandleDecisionParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the handle decision params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HandleDecisionParams) WithDefaults() *HandleDecisionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the handle decision params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HandleDecisionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the handle decision params
func (o *HandleDecisionParams) WithTimeout(timeout time.Duration) *HandleDecisionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the handle decision params
func (o *HandleDecisionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the handle decision params
func (o *HandleDecisionParams) WithContext(ctx context.Context) *HandleDecisionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the handle decision params
func (o *HandleDecisionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the handle decision params
func (o *HandleDecisionParams) WithHTTPClient(client *http.Client) *HandleDecisionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the handle decision params
func (o *HandleDecisionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *HandleDecisionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
