// Code generated by go-swagger; DO NOT EDIT.

package data

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

// NewPatchDataParams creates a new PatchDataParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchDataParams() *PatchDataParams {
	return &PatchDataParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchDataParamsWithTimeout creates a new PatchDataParams object
// with the ability to set a timeout on a request.
func NewPatchDataParamsWithTimeout(timeout time.Duration) *PatchDataParams {
	return &PatchDataParams{
		timeout: timeout,
	}
}

// NewPatchDataParamsWithContext creates a new PatchDataParams object
// with the ability to set a context for a request.
func NewPatchDataParamsWithContext(ctx context.Context) *PatchDataParams {
	return &PatchDataParams{
		Context: ctx,
	}
}

// NewPatchDataParamsWithHTTPClient creates a new PatchDataParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchDataParamsWithHTTPClient(client *http.Client) *PatchDataParams {
	return &PatchDataParams{
		HTTPClient: client,
	}
}

/* PatchDataParams contains all the parameters to send to the API endpoint
   for the patch data operation.

   Typically these are written to a http.Request.
*/
type PatchDataParams struct {

	/* IfMatch.

	   etag
	*/
	IfMatch *string

	/* Name.

	   data name
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch data params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchDataParams) WithDefaults() *PatchDataParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch data params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchDataParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch data params
func (o *PatchDataParams) WithTimeout(timeout time.Duration) *PatchDataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch data params
func (o *PatchDataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch data params
func (o *PatchDataParams) WithContext(ctx context.Context) *PatchDataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch data params
func (o *PatchDataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch data params
func (o *PatchDataParams) WithHTTPClient(client *http.Client) *PatchDataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch data params
func (o *PatchDataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfMatch adds the ifMatch to the patch data params
func (o *PatchDataParams) WithIfMatch(ifMatch *string) *PatchDataParams {
	o.SetIfMatch(ifMatch)
	return o
}

// SetIfMatch adds the ifMatch to the patch data params
func (o *PatchDataParams) SetIfMatch(ifMatch *string) {
	o.IfMatch = ifMatch
}

// WithName adds the name to the patch data params
func (o *PatchDataParams) WithName(name string) *PatchDataParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the patch data params
func (o *PatchDataParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *PatchDataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IfMatch != nil {

		// header param If-Match
		if err := r.SetHeaderParam("If-Match", *o.IfMatch); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}