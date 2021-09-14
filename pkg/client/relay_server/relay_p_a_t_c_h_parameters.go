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

// NewRelayPATCHParams creates a new RelayPATCHParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRelayPATCHParams() *RelayPATCHParams {
	return &RelayPATCHParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRelayPATCHParamsWithTimeout creates a new RelayPATCHParams object
// with the ability to set a timeout on a request.
func NewRelayPATCHParamsWithTimeout(timeout time.Duration) *RelayPATCHParams {
	return &RelayPATCHParams{
		timeout: timeout,
	}
}

// NewRelayPATCHParamsWithContext creates a new RelayPATCHParams object
// with the ability to set a context for a request.
func NewRelayPATCHParamsWithContext(ctx context.Context) *RelayPATCHParams {
	return &RelayPATCHParams{
		Context: ctx,
	}
}

// NewRelayPATCHParamsWithHTTPClient creates a new RelayPATCHParams object
// with the ability to set a custom HTTPClient for a request.
func NewRelayPATCHParamsWithHTTPClient(client *http.Client) *RelayPATCHParams {
	return &RelayPATCHParams{
		HTTPClient: client,
	}
}

/* RelayPATCHParams contains all the parameters to send to the API endpoint
   for the relay p a t c h operation.

   Typically these are written to a http.Request.
*/
type RelayPATCHParams struct {

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

// WithDefaults hydrates default values in the relay p a t c h params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RelayPATCHParams) WithDefaults() *RelayPATCHParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the relay p a t c h params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RelayPATCHParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the relay p a t c h params
func (o *RelayPATCHParams) WithTimeout(timeout time.Duration) *RelayPATCHParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the relay p a t c h params
func (o *RelayPATCHParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the relay p a t c h params
func (o *RelayPATCHParams) WithContext(ctx context.Context) *RelayPATCHParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the relay p a t c h params
func (o *RelayPATCHParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the relay p a t c h params
func (o *RelayPATCHParams) WithHTTPClient(client *http.Client) *RelayPATCHParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the relay p a t c h params
func (o *RelayPATCHParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKey adds the key to the relay p a t c h params
func (o *RelayPATCHParams) WithKey(key string) *RelayPATCHParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the relay p a t c h params
func (o *RelayPATCHParams) SetKey(key string) {
	o.Key = key
}

// WithPath adds the path to the relay p a t c h params
func (o *RelayPATCHParams) WithPath(path string) *RelayPATCHParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the relay p a t c h params
func (o *RelayPATCHParams) SetPath(path string) {
	o.Path = path
}

// WriteToRequest writes these params to a swagger request
func (o *RelayPATCHParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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