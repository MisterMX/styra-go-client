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

// NewListPoliciesParams creates a new ListPoliciesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListPoliciesParams() *ListPoliciesParams {
	return &ListPoliciesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListPoliciesParamsWithTimeout creates a new ListPoliciesParams object
// with the ability to set a timeout on a request.
func NewListPoliciesParamsWithTimeout(timeout time.Duration) *ListPoliciesParams {
	return &ListPoliciesParams{
		timeout: timeout,
	}
}

// NewListPoliciesParamsWithContext creates a new ListPoliciesParams object
// with the ability to set a context for a request.
func NewListPoliciesParamsWithContext(ctx context.Context) *ListPoliciesParams {
	return &ListPoliciesParams{
		Context: ctx,
	}
}

// NewListPoliciesParamsWithHTTPClient creates a new ListPoliciesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListPoliciesParamsWithHTTPClient(client *http.Client) *ListPoliciesParams {
	return &ListPoliciesParams{
		HTTPClient: client,
	}
}

/* ListPoliciesParams contains all the parameters to send to the API endpoint
   for the list policies operation.

   Typically these are written to a http.Request.
*/
type ListPoliciesParams struct {

	/* Drafts.

	   return rego metadata for draft policies (when metadata flag is used)
	*/
	Drafts *bool

	/* Metadata.

	   return rego metadata of specified type or all if no type provided
	*/
	Metadata *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListPoliciesParams) WithDefaults() *ListPoliciesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListPoliciesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list policies params
func (o *ListPoliciesParams) WithTimeout(timeout time.Duration) *ListPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list policies params
func (o *ListPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list policies params
func (o *ListPoliciesParams) WithContext(ctx context.Context) *ListPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list policies params
func (o *ListPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list policies params
func (o *ListPoliciesParams) WithHTTPClient(client *http.Client) *ListPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list policies params
func (o *ListPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDrafts adds the drafts to the list policies params
func (o *ListPoliciesParams) WithDrafts(drafts *bool) *ListPoliciesParams {
	o.SetDrafts(drafts)
	return o
}

// SetDrafts adds the drafts to the list policies params
func (o *ListPoliciesParams) SetDrafts(drafts *bool) {
	o.Drafts = drafts
}

// WithMetadata adds the metadata to the list policies params
func (o *ListPoliciesParams) WithMetadata(metadata *string) *ListPoliciesParams {
	o.SetMetadata(metadata)
	return o
}

// SetMetadata adds the metadata to the list policies params
func (o *ListPoliciesParams) SetMetadata(metadata *string) {
	o.Metadata = metadata
}

// WriteToRequest writes these params to a swagger request
func (o *ListPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Drafts != nil {

		// query param drafts
		var qrDrafts bool

		if o.Drafts != nil {
			qrDrafts = *o.Drafts
		}
		qDrafts := swag.FormatBool(qrDrafts)
		if qDrafts != "" {

			if err := r.SetQueryParam("drafts", qDrafts); err != nil {
				return err
			}
		}
	}

	if o.Metadata != nil {

		// query param metadata
		var qrMetadata string

		if o.Metadata != nil {
			qrMetadata = *o.Metadata
		}
		qMetadata := qrMetadata
		if qMetadata != "" {

			if err := r.SetQueryParam("metadata", qMetadata); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
