// Code generated by go-swagger; DO NOT EDIT.

package stacks

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

// NewSourceControlVerifyConfigStackParams creates a new SourceControlVerifyConfigStackParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSourceControlVerifyConfigStackParams() *SourceControlVerifyConfigStackParams {
	return &SourceControlVerifyConfigStackParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSourceControlVerifyConfigStackParamsWithTimeout creates a new SourceControlVerifyConfigStackParams object
// with the ability to set a timeout on a request.
func NewSourceControlVerifyConfigStackParamsWithTimeout(timeout time.Duration) *SourceControlVerifyConfigStackParams {
	return &SourceControlVerifyConfigStackParams{
		timeout: timeout,
	}
}

// NewSourceControlVerifyConfigStackParamsWithContext creates a new SourceControlVerifyConfigStackParams object
// with the ability to set a context for a request.
func NewSourceControlVerifyConfigStackParamsWithContext(ctx context.Context) *SourceControlVerifyConfigStackParams {
	return &SourceControlVerifyConfigStackParams{
		Context: ctx,
	}
}

// NewSourceControlVerifyConfigStackParamsWithHTTPClient creates a new SourceControlVerifyConfigStackParams object
// with the ability to set a custom HTTPClient for a request.
func NewSourceControlVerifyConfigStackParamsWithHTTPClient(client *http.Client) *SourceControlVerifyConfigStackParams {
	return &SourceControlVerifyConfigStackParams{
		HTTPClient: client,
	}
}

/* SourceControlVerifyConfigStackParams contains all the parameters to send to the API endpoint
   for the source control verify config stack operation.

   Typically these are written to a http.Request.
*/
type SourceControlVerifyConfigStackParams struct {

	// Body.
	Body *models.GitV1VerifyConfigRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the source control verify config stack params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SourceControlVerifyConfigStackParams) WithDefaults() *SourceControlVerifyConfigStackParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the source control verify config stack params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SourceControlVerifyConfigStackParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the source control verify config stack params
func (o *SourceControlVerifyConfigStackParams) WithTimeout(timeout time.Duration) *SourceControlVerifyConfigStackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the source control verify config stack params
func (o *SourceControlVerifyConfigStackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the source control verify config stack params
func (o *SourceControlVerifyConfigStackParams) WithContext(ctx context.Context) *SourceControlVerifyConfigStackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the source control verify config stack params
func (o *SourceControlVerifyConfigStackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the source control verify config stack params
func (o *SourceControlVerifyConfigStackParams) WithHTTPClient(client *http.Client) *SourceControlVerifyConfigStackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the source control verify config stack params
func (o *SourceControlVerifyConfigStackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the source control verify config stack params
func (o *SourceControlVerifyConfigStackParams) WithBody(body *models.GitV1VerifyConfigRequest) *SourceControlVerifyConfigStackParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the source control verify config stack params
func (o *SourceControlVerifyConfigStackParams) SetBody(body *models.GitV1VerifyConfigRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SourceControlVerifyConfigStackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
