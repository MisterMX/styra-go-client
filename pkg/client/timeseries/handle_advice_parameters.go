// Code generated by go-swagger; DO NOT EDIT.

package timeseries

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

// NewHandleAdviceParams creates a new HandleAdviceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHandleAdviceParams() *HandleAdviceParams {
	return &HandleAdviceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHandleAdviceParamsWithTimeout creates a new HandleAdviceParams object
// with the ability to set a timeout on a request.
func NewHandleAdviceParamsWithTimeout(timeout time.Duration) *HandleAdviceParams {
	return &HandleAdviceParams{
		timeout: timeout,
	}
}

// NewHandleAdviceParamsWithContext creates a new HandleAdviceParams object
// with the ability to set a context for a request.
func NewHandleAdviceParamsWithContext(ctx context.Context) *HandleAdviceParams {
	return &HandleAdviceParams{
		Context: ctx,
	}
}

// NewHandleAdviceParamsWithHTTPClient creates a new HandleAdviceParams object
// with the ability to set a custom HTTPClient for a request.
func NewHandleAdviceParamsWithHTTPClient(client *http.Client) *HandleAdviceParams {
	return &HandleAdviceParams{
		HTTPClient: client,
	}
}

/* HandleAdviceParams contains all the parameters to send to the API endpoint
   for the handle advice operation.

   Typically these are written to a http.Request.
*/
type HandleAdviceParams struct {

	// Body.
	Body *models.TimeseriesV1TimeSeriesPostRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the handle advice params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HandleAdviceParams) WithDefaults() *HandleAdviceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the handle advice params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HandleAdviceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the handle advice params
func (o *HandleAdviceParams) WithTimeout(timeout time.Duration) *HandleAdviceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the handle advice params
func (o *HandleAdviceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the handle advice params
func (o *HandleAdviceParams) WithContext(ctx context.Context) *HandleAdviceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the handle advice params
func (o *HandleAdviceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the handle advice params
func (o *HandleAdviceParams) WithHTTPClient(client *http.Client) *HandleAdviceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the handle advice params
func (o *HandleAdviceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the handle advice params
func (o *HandleAdviceParams) WithBody(body *models.TimeseriesV1TimeSeriesPostRequest) *HandleAdviceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the handle advice params
func (o *HandleAdviceParams) SetBody(body *models.TimeseriesV1TimeSeriesPostRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *HandleAdviceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
