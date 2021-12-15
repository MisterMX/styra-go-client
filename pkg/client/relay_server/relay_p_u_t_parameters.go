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

	"github.com/mistermx/styra-go-client/pkg/models"
)

// NewRelayPUTParams creates a new RelayPUTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRelayPUTParams() *RelayPUTParams {
	return &RelayPUTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRelayPUTParamsWithTimeout creates a new RelayPUTParams object
// with the ability to set a timeout on a request.
func NewRelayPUTParamsWithTimeout(timeout time.Duration) *RelayPUTParams {
	return &RelayPUTParams{
		timeout: timeout,
	}
}

// NewRelayPUTParamsWithContext creates a new RelayPUTParams object
// with the ability to set a context for a request.
func NewRelayPUTParamsWithContext(ctx context.Context) *RelayPUTParams {
	return &RelayPUTParams{
		Context: ctx,
	}
}

// NewRelayPUTParamsWithHTTPClient creates a new RelayPUTParams object
// with the ability to set a custom HTTPClient for a request.
func NewRelayPUTParamsWithHTTPClient(client *http.Client) *RelayPUTParams {
	return &RelayPUTParams{
		HTTPClient: client,
	}
}

/* RelayPUTParams contains all the parameters to send to the API endpoint
   for the relay p u t operation.

   Typically these are written to a http.Request.
*/
type RelayPUTParams struct {

	// Body.
	Body models.MetaV1RequestObject

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

// WithDefaults hydrates default values in the relay p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RelayPUTParams) WithDefaults() *RelayPUTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the relay p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RelayPUTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the relay p u t params
func (o *RelayPUTParams) WithTimeout(timeout time.Duration) *RelayPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the relay p u t params
func (o *RelayPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the relay p u t params
func (o *RelayPUTParams) WithContext(ctx context.Context) *RelayPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the relay p u t params
func (o *RelayPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the relay p u t params
func (o *RelayPUTParams) WithHTTPClient(client *http.Client) *RelayPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the relay p u t params
func (o *RelayPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the relay p u t params
func (o *RelayPUTParams) WithBody(body models.MetaV1RequestObject) *RelayPUTParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the relay p u t params
func (o *RelayPUTParams) SetBody(body models.MetaV1RequestObject) {
	o.Body = body
}

// WithKey adds the key to the relay p u t params
func (o *RelayPUTParams) WithKey(key string) *RelayPUTParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the relay p u t params
func (o *RelayPUTParams) SetKey(key string) {
	o.Key = key
}

// WithPath adds the path to the relay p u t params
func (o *RelayPUTParams) WithPath(path string) *RelayPUTParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the relay p u t params
func (o *RelayPUTParams) SetPath(path string) {
	o.Path = path
}

// WriteToRequest writes these params to a swagger request
func (o *RelayPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
