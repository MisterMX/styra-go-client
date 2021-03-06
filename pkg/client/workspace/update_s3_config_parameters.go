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

// NewUpdateS3ConfigParams creates a new UpdateS3ConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateS3ConfigParams() *UpdateS3ConfigParams {
	return &UpdateS3ConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateS3ConfigParamsWithTimeout creates a new UpdateS3ConfigParams object
// with the ability to set a timeout on a request.
func NewUpdateS3ConfigParamsWithTimeout(timeout time.Duration) *UpdateS3ConfigParams {
	return &UpdateS3ConfigParams{
		timeout: timeout,
	}
}

// NewUpdateS3ConfigParamsWithContext creates a new UpdateS3ConfigParams object
// with the ability to set a context for a request.
func NewUpdateS3ConfigParamsWithContext(ctx context.Context) *UpdateS3ConfigParams {
	return &UpdateS3ConfigParams{
		Context: ctx,
	}
}

// NewUpdateS3ConfigParamsWithHTTPClient creates a new UpdateS3ConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateS3ConfigParamsWithHTTPClient(client *http.Client) *UpdateS3ConfigParams {
	return &UpdateS3ConfigParams{
		HTTPClient: client,
	}
}

/* UpdateS3ConfigParams contains all the parameters to send to the API endpoint
   for the update s3 config operation.

   Typically these are written to a http.Request.
*/
type UpdateS3ConfigParams struct {

	// Body.
	Body *models.WorkspaceV1S3ConfigPutRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update s3 config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateS3ConfigParams) WithDefaults() *UpdateS3ConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update s3 config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateS3ConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update s3 config params
func (o *UpdateS3ConfigParams) WithTimeout(timeout time.Duration) *UpdateS3ConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update s3 config params
func (o *UpdateS3ConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update s3 config params
func (o *UpdateS3ConfigParams) WithContext(ctx context.Context) *UpdateS3ConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update s3 config params
func (o *UpdateS3ConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update s3 config params
func (o *UpdateS3ConfigParams) WithHTTPClient(client *http.Client) *UpdateS3ConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update s3 config params
func (o *UpdateS3ConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update s3 config params
func (o *UpdateS3ConfigParams) WithBody(body *models.WorkspaceV1S3ConfigPutRequest) *UpdateS3ConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update s3 config params
func (o *UpdateS3ConfigParams) SetBody(body *models.WorkspaceV1S3ConfigPutRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateS3ConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
