// Code generated by go-swagger; DO NOT EDIT.

package activity

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

// NewHandleActivityParams creates a new HandleActivityParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHandleActivityParams() *HandleActivityParams {
	return &HandleActivityParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHandleActivityParamsWithTimeout creates a new HandleActivityParams object
// with the ability to set a timeout on a request.
func NewHandleActivityParamsWithTimeout(timeout time.Duration) *HandleActivityParams {
	return &HandleActivityParams{
		timeout: timeout,
	}
}

// NewHandleActivityParamsWithContext creates a new HandleActivityParams object
// with the ability to set a context for a request.
func NewHandleActivityParamsWithContext(ctx context.Context) *HandleActivityParams {
	return &HandleActivityParams{
		Context: ctx,
	}
}

// NewHandleActivityParamsWithHTTPClient creates a new HandleActivityParams object
// with the ability to set a custom HTTPClient for a request.
func NewHandleActivityParamsWithHTTPClient(client *http.Client) *HandleActivityParams {
	return &HandleActivityParams{
		HTTPClient: client,
	}
}

/* HandleActivityParams contains all the parameters to send to the API endpoint
   for the handle activity operation.

   Typically these are written to a http.Request.
*/
type HandleActivityParams struct {

	// Body.
	Body *models.ActivityV1ActivityPostRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the handle activity params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HandleActivityParams) WithDefaults() *HandleActivityParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the handle activity params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HandleActivityParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the handle activity params
func (o *HandleActivityParams) WithTimeout(timeout time.Duration) *HandleActivityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the handle activity params
func (o *HandleActivityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the handle activity params
func (o *HandleActivityParams) WithContext(ctx context.Context) *HandleActivityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the handle activity params
func (o *HandleActivityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the handle activity params
func (o *HandleActivityParams) WithHTTPClient(client *http.Client) *HandleActivityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the handle activity params
func (o *HandleActivityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the handle activity params
func (o *HandleActivityParams) WithBody(body *models.ActivityV1ActivityPostRequest) *HandleActivityParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the handle activity params
func (o *HandleActivityParams) SetBody(body *models.ActivityV1ActivityPostRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *HandleActivityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
