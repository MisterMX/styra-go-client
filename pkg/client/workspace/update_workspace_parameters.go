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

// NewUpdateWorkspaceParams creates a new UpdateWorkspaceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateWorkspaceParams() *UpdateWorkspaceParams {
	return &UpdateWorkspaceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateWorkspaceParamsWithTimeout creates a new UpdateWorkspaceParams object
// with the ability to set a timeout on a request.
func NewUpdateWorkspaceParamsWithTimeout(timeout time.Duration) *UpdateWorkspaceParams {
	return &UpdateWorkspaceParams{
		timeout: timeout,
	}
}

// NewUpdateWorkspaceParamsWithContext creates a new UpdateWorkspaceParams object
// with the ability to set a context for a request.
func NewUpdateWorkspaceParamsWithContext(ctx context.Context) *UpdateWorkspaceParams {
	return &UpdateWorkspaceParams{
		Context: ctx,
	}
}

// NewUpdateWorkspaceParamsWithHTTPClient creates a new UpdateWorkspaceParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateWorkspaceParamsWithHTTPClient(client *http.Client) *UpdateWorkspaceParams {
	return &UpdateWorkspaceParams{
		HTTPClient: client,
	}
}

/* UpdateWorkspaceParams contains all the parameters to send to the API endpoint
   for the update workspace operation.

   Typically these are written to a http.Request.
*/
type UpdateWorkspaceParams struct {

	// Body.
	Body *models.V1WorkspacePutRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update workspace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateWorkspaceParams) WithDefaults() *UpdateWorkspaceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update workspace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateWorkspaceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update workspace params
func (o *UpdateWorkspaceParams) WithTimeout(timeout time.Duration) *UpdateWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update workspace params
func (o *UpdateWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update workspace params
func (o *UpdateWorkspaceParams) WithContext(ctx context.Context) *UpdateWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update workspace params
func (o *UpdateWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update workspace params
func (o *UpdateWorkspaceParams) WithHTTPClient(client *http.Client) *UpdateWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update workspace params
func (o *UpdateWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update workspace params
func (o *UpdateWorkspaceParams) WithBody(body *models.V1WorkspacePutRequest) *UpdateWorkspaceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update workspace params
func (o *UpdateWorkspaceParams) SetBody(body *models.V1WorkspacePutRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
