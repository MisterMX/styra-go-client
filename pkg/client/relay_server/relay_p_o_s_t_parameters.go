// Code generated by go-swagger; DO NOT EDIT.

package relay_server

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

// NewRelayPOSTParams creates a new RelayPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRelayPOSTParams() *RelayPOSTParams {
	return &RelayPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRelayPOSTParamsWithTimeout creates a new RelayPOSTParams object
// with the ability to set a timeout on a request.
func NewRelayPOSTParamsWithTimeout(timeout time.Duration) *RelayPOSTParams {
	return &RelayPOSTParams{
		timeout: timeout,
	}
}

// NewRelayPOSTParamsWithContext creates a new RelayPOSTParams object
// with the ability to set a context for a request.
func NewRelayPOSTParamsWithContext(ctx context.Context) *RelayPOSTParams {
	return &RelayPOSTParams{
		Context: ctx,
	}
}

// NewRelayPOSTParamsWithHTTPClient creates a new RelayPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewRelayPOSTParamsWithHTTPClient(client *http.Client) *RelayPOSTParams {
	return &RelayPOSTParams{
		HTTPClient: client,
	}
}

/* RelayPOSTParams contains all the parameters to send to the API endpoint
   for the relay p o s t operation.

   Typically these are written to a http.Request.
*/
type RelayPOSTParams struct {

	/* Key.

	   identifies the relay client that should execute relayed request
	*/
	Key string

	/* Path.

	   path to be queried on the base-url that relay client is configured with
	*/
	Path string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the relay p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RelayPOSTParams) WithDefaults() *RelayPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the relay p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RelayPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the relay p o s t params
func (o *RelayPOSTParams) WithTimeout(timeout time.Duration) *RelayPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the relay p o s t params
func (o *RelayPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the relay p o s t params
func (o *RelayPOSTParams) WithContext(ctx context.Context) *RelayPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the relay p o s t params
func (o *RelayPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the relay p o s t params
func (o *RelayPOSTParams) WithHTTPClient(client *http.Client) *RelayPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the relay p o s t params
func (o *RelayPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKey adds the key to the relay p o s t params
func (o *RelayPOSTParams) WithKey(key string) *RelayPOSTParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the relay p o s t params
func (o *RelayPOSTParams) SetKey(key string) {
	o.Key = key
}

// WithPath adds the path to the relay p o s t params
func (o *RelayPOSTParams) WithPath(path string) *RelayPOSTParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the relay p o s t params
func (o *RelayPOSTParams) SetPath(path string) {
	o.Path = path
}

// WriteToRequest writes these params to a swagger request
func (o *RelayPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param key
	if err := r.SetPathParam("key", o.Key); err != nil {
		return err
	}

	// path param path
	if err := r.SetPathParam("path", o.Path); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
