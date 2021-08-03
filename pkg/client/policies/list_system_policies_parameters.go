// Code generated by go-swagger; DO NOT EDIT.

package policies

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

// NewListSystemPoliciesParams creates a new ListSystemPoliciesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListSystemPoliciesParams() *ListSystemPoliciesParams {
	return &ListSystemPoliciesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListSystemPoliciesParamsWithTimeout creates a new ListSystemPoliciesParams object
// with the ability to set a timeout on a request.
func NewListSystemPoliciesParamsWithTimeout(timeout time.Duration) *ListSystemPoliciesParams {
	return &ListSystemPoliciesParams{
		timeout: timeout,
	}
}

// NewListSystemPoliciesParamsWithContext creates a new ListSystemPoliciesParams object
// with the ability to set a context for a request.
func NewListSystemPoliciesParamsWithContext(ctx context.Context) *ListSystemPoliciesParams {
	return &ListSystemPoliciesParams{
		Context: ctx,
	}
}

// NewListSystemPoliciesParamsWithHTTPClient creates a new ListSystemPoliciesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListSystemPoliciesParamsWithHTTPClient(client *http.Client) *ListSystemPoliciesParams {
	return &ListSystemPoliciesParams{
		HTTPClient: client,
	}
}

/* ListSystemPoliciesParams contains all the parameters to send to the API endpoint
   for the list system policies operation.

   Typically these are written to a http.Request.
*/
type ListSystemPoliciesParams struct {

	/* Drafts.

	   return rego metadata for draft policies (when metadata flag is used)
	*/
	Drafts bool

	/* Metadata.

	   return rego metadata of specified type or all if no type provided
	*/
	Metadata string

	/* System.

	   system id
	*/
	System string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list system policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListSystemPoliciesParams) WithDefaults() *ListSystemPoliciesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list system policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListSystemPoliciesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list system policies params
func (o *ListSystemPoliciesParams) WithTimeout(timeout time.Duration) *ListSystemPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list system policies params
func (o *ListSystemPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list system policies params
func (o *ListSystemPoliciesParams) WithContext(ctx context.Context) *ListSystemPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list system policies params
func (o *ListSystemPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list system policies params
func (o *ListSystemPoliciesParams) WithHTTPClient(client *http.Client) *ListSystemPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list system policies params
func (o *ListSystemPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDrafts adds the drafts to the list system policies params
func (o *ListSystemPoliciesParams) WithDrafts(drafts bool) *ListSystemPoliciesParams {
	o.SetDrafts(drafts)
	return o
}

// SetDrafts adds the drafts to the list system policies params
func (o *ListSystemPoliciesParams) SetDrafts(drafts bool) {
	o.Drafts = drafts
}

// WithMetadata adds the metadata to the list system policies params
func (o *ListSystemPoliciesParams) WithMetadata(metadata string) *ListSystemPoliciesParams {
	o.SetMetadata(metadata)
	return o
}

// SetMetadata adds the metadata to the list system policies params
func (o *ListSystemPoliciesParams) SetMetadata(metadata string) {
	o.Metadata = metadata
}

// WithSystem adds the system to the list system policies params
func (o *ListSystemPoliciesParams) WithSystem(system string) *ListSystemPoliciesParams {
	o.SetSystem(system)
	return o
}

// SetSystem adds the system to the list system policies params
func (o *ListSystemPoliciesParams) SetSystem(system string) {
	o.System = system
}

// WriteToRequest writes these params to a swagger request
func (o *ListSystemPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param drafts
	if err := r.SetPathParam("drafts", swag.FormatBool(o.Drafts)); err != nil {
		return err
	}

	// path param metadata
	if err := r.SetPathParam("metadata", o.Metadata); err != nil {
		return err
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
