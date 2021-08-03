// Code generated by go-swagger; DO NOT EDIT.

package identity_providers

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

// NewGetProviderParams creates a new GetProviderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProviderParams() *GetProviderParams {
	return &GetProviderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProviderParamsWithTimeout creates a new GetProviderParams object
// with the ability to set a timeout on a request.
func NewGetProviderParamsWithTimeout(timeout time.Duration) *GetProviderParams {
	return &GetProviderParams{
		timeout: timeout,
	}
}

// NewGetProviderParamsWithContext creates a new GetProviderParams object
// with the ability to set a context for a request.
func NewGetProviderParamsWithContext(ctx context.Context) *GetProviderParams {
	return &GetProviderParams{
		Context: ctx,
	}
}

// NewGetProviderParamsWithHTTPClient creates a new GetProviderParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProviderParamsWithHTTPClient(client *http.Client) *GetProviderParams {
	return &GetProviderParams{
		HTTPClient: client,
	}
}

/* GetProviderParams contains all the parameters to send to the API endpoint
   for the get provider operation.

   Typically these are written to a http.Request.
*/
type GetProviderParams struct {

	/* ProviderID.

	   provider ID
	*/
	ProviderID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProviderParams) WithDefaults() *GetProviderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProviderParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get provider params
func (o *GetProviderParams) WithTimeout(timeout time.Duration) *GetProviderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get provider params
func (o *GetProviderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get provider params
func (o *GetProviderParams) WithContext(ctx context.Context) *GetProviderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get provider params
func (o *GetProviderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get provider params
func (o *GetProviderParams) WithHTTPClient(client *http.Client) *GetProviderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get provider params
func (o *GetProviderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProviderID adds the providerID to the get provider params
func (o *GetProviderParams) WithProviderID(providerID string) *GetProviderParams {
	o.SetProviderID(providerID)
	return o
}

// SetProviderID adds the providerId to the get provider params
func (o *GetProviderParams) SetProviderID(providerID string) {
	o.ProviderID = providerID
}

// WriteToRequest writes these params to a swagger request
func (o *GetProviderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param providerId
	if err := r.SetPathParam("providerId", o.ProviderID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
