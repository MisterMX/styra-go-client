// Code generated by go-swagger; DO NOT EDIT.

package login

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

// NewOAuth2CallbackParams creates a new OAuth2CallbackParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOAuth2CallbackParams() *OAuth2CallbackParams {
	return &OAuth2CallbackParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOAuth2CallbackParamsWithTimeout creates a new OAuth2CallbackParams object
// with the ability to set a timeout on a request.
func NewOAuth2CallbackParamsWithTimeout(timeout time.Duration) *OAuth2CallbackParams {
	return &OAuth2CallbackParams{
		timeout: timeout,
	}
}

// NewOAuth2CallbackParamsWithContext creates a new OAuth2CallbackParams object
// with the ability to set a context for a request.
func NewOAuth2CallbackParamsWithContext(ctx context.Context) *OAuth2CallbackParams {
	return &OAuth2CallbackParams{
		Context: ctx,
	}
}

// NewOAuth2CallbackParamsWithHTTPClient creates a new OAuth2CallbackParams object
// with the ability to set a custom HTTPClient for a request.
func NewOAuth2CallbackParamsWithHTTPClient(client *http.Client) *OAuth2CallbackParams {
	return &OAuth2CallbackParams{
		HTTPClient: client,
	}
}

/* OAuth2CallbackParams contains all the parameters to send to the API endpoint
   for the o auth2 callback operation.

   Typically these are written to a http.Request.
*/
type OAuth2CallbackParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the o auth2 callback params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OAuth2CallbackParams) WithDefaults() *OAuth2CallbackParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the o auth2 callback params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OAuth2CallbackParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the o auth2 callback params
func (o *OAuth2CallbackParams) WithTimeout(timeout time.Duration) *OAuth2CallbackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the o auth2 callback params
func (o *OAuth2CallbackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the o auth2 callback params
func (o *OAuth2CallbackParams) WithContext(ctx context.Context) *OAuth2CallbackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the o auth2 callback params
func (o *OAuth2CallbackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the o auth2 callback params
func (o *OAuth2CallbackParams) WithHTTPClient(client *http.Client) *OAuth2CallbackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the o auth2 callback params
func (o *OAuth2CallbackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *OAuth2CallbackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
