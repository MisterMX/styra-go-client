// Code generated by go-swagger; DO NOT EDIT.

package systems

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

// NewGetSystemAgentsParams creates a new GetSystemAgentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSystemAgentsParams() *GetSystemAgentsParams {
	return &GetSystemAgentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSystemAgentsParamsWithTimeout creates a new GetSystemAgentsParams object
// with the ability to set a timeout on a request.
func NewGetSystemAgentsParamsWithTimeout(timeout time.Duration) *GetSystemAgentsParams {
	return &GetSystemAgentsParams{
		timeout: timeout,
	}
}

// NewGetSystemAgentsParamsWithContext creates a new GetSystemAgentsParams object
// with the ability to set a context for a request.
func NewGetSystemAgentsParamsWithContext(ctx context.Context) *GetSystemAgentsParams {
	return &GetSystemAgentsParams{
		Context: ctx,
	}
}

// NewGetSystemAgentsParamsWithHTTPClient creates a new GetSystemAgentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSystemAgentsParamsWithHTTPClient(client *http.Client) *GetSystemAgentsParams {
	return &GetSystemAgentsParams{
		HTTPClient: client,
	}
}

/* GetSystemAgentsParams contains all the parameters to send to the API endpoint
   for the get system agents operation.

   Typically these are written to a http.Request.
*/
type GetSystemAgentsParams struct {

	/* System.

	   system ID
	*/
	System string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get system agents params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSystemAgentsParams) WithDefaults() *GetSystemAgentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get system agents params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSystemAgentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get system agents params
func (o *GetSystemAgentsParams) WithTimeout(timeout time.Duration) *GetSystemAgentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get system agents params
func (o *GetSystemAgentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get system agents params
func (o *GetSystemAgentsParams) WithContext(ctx context.Context) *GetSystemAgentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get system agents params
func (o *GetSystemAgentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get system agents params
func (o *GetSystemAgentsParams) WithHTTPClient(client *http.Client) *GetSystemAgentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get system agents params
func (o *GetSystemAgentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSystem adds the system to the get system agents params
func (o *GetSystemAgentsParams) WithSystem(system string) *GetSystemAgentsParams {
	o.SetSystem(system)
	return o
}

// SetSystem adds the system to the get system agents params
func (o *GetSystemAgentsParams) SetSystem(system string) {
	o.System = system
}

// WriteToRequest writes these params to a swagger request
func (o *GetSystemAgentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param system
	if err := r.SetPathParam("system", o.System); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
