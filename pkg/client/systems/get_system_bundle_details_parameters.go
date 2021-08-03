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
	"github.com/go-openapi/swag"
)

// NewGetSystemBundleDetailsParams creates a new GetSystemBundleDetailsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSystemBundleDetailsParams() *GetSystemBundleDetailsParams {
	return &GetSystemBundleDetailsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSystemBundleDetailsParamsWithTimeout creates a new GetSystemBundleDetailsParams object
// with the ability to set a timeout on a request.
func NewGetSystemBundleDetailsParamsWithTimeout(timeout time.Duration) *GetSystemBundleDetailsParams {
	return &GetSystemBundleDetailsParams{
		timeout: timeout,
	}
}

// NewGetSystemBundleDetailsParamsWithContext creates a new GetSystemBundleDetailsParams object
// with the ability to set a context for a request.
func NewGetSystemBundleDetailsParamsWithContext(ctx context.Context) *GetSystemBundleDetailsParams {
	return &GetSystemBundleDetailsParams{
		Context: ctx,
	}
}

// NewGetSystemBundleDetailsParamsWithHTTPClient creates a new GetSystemBundleDetailsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSystemBundleDetailsParamsWithHTTPClient(client *http.Client) *GetSystemBundleDetailsParams {
	return &GetSystemBundleDetailsParams{
		HTTPClient: client,
	}
}

/* GetSystemBundleDetailsParams contains all the parameters to send to the API endpoint
   for the get system bundle details operation.

   Typically these are written to a http.Request.
*/
type GetSystemBundleDetailsParams struct {

	/* Bundle.

	   bundle ID
	*/
	Bundle string

	/* System.

	   system ID
	*/
	System string

	/* Version.

	   version #

	   Format: int32
	*/
	Version int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get system bundle details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSystemBundleDetailsParams) WithDefaults() *GetSystemBundleDetailsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get system bundle details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSystemBundleDetailsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get system bundle details params
func (o *GetSystemBundleDetailsParams) WithTimeout(timeout time.Duration) *GetSystemBundleDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get system bundle details params
func (o *GetSystemBundleDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get system bundle details params
func (o *GetSystemBundleDetailsParams) WithContext(ctx context.Context) *GetSystemBundleDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get system bundle details params
func (o *GetSystemBundleDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get system bundle details params
func (o *GetSystemBundleDetailsParams) WithHTTPClient(client *http.Client) *GetSystemBundleDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get system bundle details params
func (o *GetSystemBundleDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBundle adds the bundle to the get system bundle details params
func (o *GetSystemBundleDetailsParams) WithBundle(bundle string) *GetSystemBundleDetailsParams {
	o.SetBundle(bundle)
	return o
}

// SetBundle adds the bundle to the get system bundle details params
func (o *GetSystemBundleDetailsParams) SetBundle(bundle string) {
	o.Bundle = bundle
}

// WithSystem adds the system to the get system bundle details params
func (o *GetSystemBundleDetailsParams) WithSystem(system string) *GetSystemBundleDetailsParams {
	o.SetSystem(system)
	return o
}

// SetSystem adds the system to the get system bundle details params
func (o *GetSystemBundleDetailsParams) SetSystem(system string) {
	o.System = system
}

// WithVersion adds the version to the get system bundle details params
func (o *GetSystemBundleDetailsParams) WithVersion(version int32) *GetSystemBundleDetailsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get system bundle details params
func (o *GetSystemBundleDetailsParams) SetVersion(version int32) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetSystemBundleDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param bundle
	if err := r.SetPathParam("bundle", o.Bundle); err != nil {
		return err
	}

	// path param system
	if err := r.SetPathParam("system", o.System); err != nil {
		return err
	}

	// path param version
	if err := r.SetPathParam("version", swag.FormatInt32(o.Version)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
