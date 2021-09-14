// Code generated by go-swagger; DO NOT EDIT.

package workspace

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

// NewFunc3Params creates a new Func3Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFunc3Params() *Func3Params {
	return &Func3Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewFunc3ParamsWithTimeout creates a new Func3Params object
// with the ability to set a timeout on a request.
func NewFunc3ParamsWithTimeout(timeout time.Duration) *Func3Params {
	return &Func3Params{
		timeout: timeout,
	}
}

// NewFunc3ParamsWithContext creates a new Func3Params object
// with the ability to set a context for a request.
func NewFunc3ParamsWithContext(ctx context.Context) *Func3Params {
	return &Func3Params{
		Context: ctx,
	}
}

// NewFunc3ParamsWithHTTPClient creates a new Func3Params object
// with the ability to set a custom HTTPClient for a request.
func NewFunc3ParamsWithHTTPClient(client *http.Client) *Func3Params {
	return &Func3Params{
		HTTPClient: client,
	}
}

/* Func3Params contains all the parameters to send to the API endpoint
   for the func3 operation.

   Typically these are written to a http.Request.
*/
type Func3Params struct {

	// Body.
	Body *models.V1VerifyConfigRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the func3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Func3Params) WithDefaults() *Func3Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the func3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Func3Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the func3 params
func (o *Func3Params) WithTimeout(timeout time.Duration) *Func3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the func3 params
func (o *Func3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the func3 params
func (o *Func3Params) WithContext(ctx context.Context) *Func3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the func3 params
func (o *Func3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the func3 params
func (o *Func3Params) WithHTTPClient(client *http.Client) *Func3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the func3 params
func (o *Func3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the func3 params
func (o *Func3Params) WithBody(body *models.V1VerifyConfigRequest) *Func3Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the func3 params
func (o *Func3Params) SetBody(body *models.V1VerifyConfigRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *Func3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}