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

// NewGetOPADiscoveryConfigParams creates a new GetOPADiscoveryConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOPADiscoveryConfigParams() *GetOPADiscoveryConfigParams {
	return &GetOPADiscoveryConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOPADiscoveryConfigParamsWithTimeout creates a new GetOPADiscoveryConfigParams object
// with the ability to set a timeout on a request.
func NewGetOPADiscoveryConfigParamsWithTimeout(timeout time.Duration) *GetOPADiscoveryConfigParams {
	return &GetOPADiscoveryConfigParams{
		timeout: timeout,
	}
}

// NewGetOPADiscoveryConfigParamsWithContext creates a new GetOPADiscoveryConfigParams object
// with the ability to set a context for a request.
func NewGetOPADiscoveryConfigParamsWithContext(ctx context.Context) *GetOPADiscoveryConfigParams {
	return &GetOPADiscoveryConfigParams{
		Context: ctx,
	}
}

// NewGetOPADiscoveryConfigParamsWithHTTPClient creates a new GetOPADiscoveryConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOPADiscoveryConfigParamsWithHTTPClient(client *http.Client) *GetOPADiscoveryConfigParams {
	return &GetOPADiscoveryConfigParams{
		HTTPClient: client,
	}
}

/* GetOPADiscoveryConfigParams contains all the parameters to send to the API endpoint
   for the get o p a discovery config operation.

   Typically these are written to a http.Request.
*/
type GetOPADiscoveryConfigParams struct {

	/* IfNoneMatch.

	   etag
	*/
	IfNoneMatch *string

	/* System.

	   system ID
	*/
	System string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get o p a discovery config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOPADiscoveryConfigParams) WithDefaults() *GetOPADiscoveryConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get o p a discovery config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOPADiscoveryConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get o p a discovery config params
func (o *GetOPADiscoveryConfigParams) WithTimeout(timeout time.Duration) *GetOPADiscoveryConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get o p a discovery config params
func (o *GetOPADiscoveryConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get o p a discovery config params
func (o *GetOPADiscoveryConfigParams) WithContext(ctx context.Context) *GetOPADiscoveryConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get o p a discovery config params
func (o *GetOPADiscoveryConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get o p a discovery config params
func (o *GetOPADiscoveryConfigParams) WithHTTPClient(client *http.Client) *GetOPADiscoveryConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get o p a discovery config params
func (o *GetOPADiscoveryConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get o p a discovery config params
func (o *GetOPADiscoveryConfigParams) WithIfNoneMatch(ifNoneMatch *string) *GetOPADiscoveryConfigParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get o p a discovery config params
func (o *GetOPADiscoveryConfigParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithSystem adds the system to the get o p a discovery config params
func (o *GetOPADiscoveryConfigParams) WithSystem(system string) *GetOPADiscoveryConfigParams {
	o.SetSystem(system)
	return o
}

// SetSystem adds the system to the get o p a discovery config params
func (o *GetOPADiscoveryConfigParams) SetSystem(system string) {
	o.System = system
}

// WriteToRequest writes these params to a swagger request
func (o *GetOPADiscoveryConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IfNoneMatch != nil {

		// header param If-None-Match
		if err := r.SetHeaderParam("If-None-Match", *o.IfNoneMatch); err != nil {
			return err
		}
	}

	// path param system
	if err := r.SetPathParam("system", o.System); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
