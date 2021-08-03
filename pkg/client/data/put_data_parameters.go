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

// NewPutDataParams creates a new PutDataParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutDataParams() *PutDataParams {
	return &PutDataParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutDataParamsWithTimeout creates a new PutDataParams object
// with the ability to set a timeout on a request.
func NewPutDataParamsWithTimeout(timeout time.Duration) *PutDataParams {
	return &PutDataParams{
		timeout: timeout,
	}
}

// NewPutDataParamsWithContext creates a new PutDataParams object
// with the ability to set a context for a request.
func NewPutDataParamsWithContext(ctx context.Context) *PutDataParams {
	return &PutDataParams{
		Context: ctx,
	}
}

// NewPutDataParamsWithHTTPClient creates a new PutDataParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutDataParamsWithHTTPClient(client *http.Client) *PutDataParams {
	return &PutDataParams{
		HTTPClient: client,
	}
}

/* PutDataParams contains all the parameters to send to the API endpoint
   for the put data operation.

   Typically these are written to a http.Request.
*/
type PutDataParams struct {

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

// WithDefaults hydrates default values in the put data params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutDataParams) WithDefaults() *PutDataParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put data params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutDataParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put data params
func (o *PutDataParams) WithTimeout(timeout time.Duration) *PutDataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put data params
func (o *PutDataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put data params
func (o *PutDataParams) WithContext(ctx context.Context) *PutDataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put data params
func (o *PutDataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put data params
func (o *PutDataParams) WithHTTPClient(client *http.Client) *PutDataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put data params
func (o *PutDataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfMatch adds the ifMatch to the put data params
func (o *PutDataParams) WithIfMatch(ifMatch *string) *PutDataParams {
	o.SetIfMatch(ifMatch)
	return o
}

// SetIfMatch adds the ifMatch to the put data params
func (o *PutDataParams) SetIfMatch(ifMatch *string) {
	o.IfMatch = ifMatch
}

// WithName adds the name to the put data params
func (o *PutDataParams) WithName(name string) *PutDataParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the put data params
func (o *PutDataParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *PutDataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
