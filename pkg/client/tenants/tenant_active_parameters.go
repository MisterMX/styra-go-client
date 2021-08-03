// Code generated by go-swagger; DO NOT EDIT.

package tenants

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

// NewTenantActiveParams creates a new TenantActiveParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTenantActiveParams() *TenantActiveParams {
	return &TenantActiveParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTenantActiveParamsWithTimeout creates a new TenantActiveParams object
// with the ability to set a timeout on a request.
func NewTenantActiveParamsWithTimeout(timeout time.Duration) *TenantActiveParams {
	return &TenantActiveParams{
		timeout: timeout,
	}
}

// NewTenantActiveParamsWithContext creates a new TenantActiveParams object
// with the ability to set a context for a request.
func NewTenantActiveParamsWithContext(ctx context.Context) *TenantActiveParams {
	return &TenantActiveParams{
		Context: ctx,
	}
}

// NewTenantActiveParamsWithHTTPClient creates a new TenantActiveParams object
// with the ability to set a custom HTTPClient for a request.
func NewTenantActiveParamsWithHTTPClient(client *http.Client) *TenantActiveParams {
	return &TenantActiveParams{
		HTTPClient: client,
	}
}

/* TenantActiveParams contains all the parameters to send to the API endpoint
   for the tenant active operation.

   Typically these are written to a http.Request.
*/
type TenantActiveParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tenant active params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenantActiveParams) WithDefaults() *TenantActiveParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tenant active params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenantActiveParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tenant active params
func (o *TenantActiveParams) WithTimeout(timeout time.Duration) *TenantActiveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tenant active params
func (o *TenantActiveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tenant active params
func (o *TenantActiveParams) WithContext(ctx context.Context) *TenantActiveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tenant active params
func (o *TenantActiveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tenant active params
func (o *TenantActiveParams) WithHTTPClient(client *http.Client) *TenantActiveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tenant active params
func (o *TenantActiveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *TenantActiveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}