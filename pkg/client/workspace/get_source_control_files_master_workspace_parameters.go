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
)

// NewGetSourceControlFilesMasterWorkspaceParams creates a new GetSourceControlFilesMasterWorkspaceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSourceControlFilesMasterWorkspaceParams() *GetSourceControlFilesMasterWorkspaceParams {
	return &GetSourceControlFilesMasterWorkspaceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSourceControlFilesMasterWorkspaceParamsWithTimeout creates a new GetSourceControlFilesMasterWorkspaceParams object
// with the ability to set a timeout on a request.
func NewGetSourceControlFilesMasterWorkspaceParamsWithTimeout(timeout time.Duration) *GetSourceControlFilesMasterWorkspaceParams {
	return &GetSourceControlFilesMasterWorkspaceParams{
		timeout: timeout,
	}
}

// NewGetSourceControlFilesMasterWorkspaceParamsWithContext creates a new GetSourceControlFilesMasterWorkspaceParams object
// with the ability to set a context for a request.
func NewGetSourceControlFilesMasterWorkspaceParamsWithContext(ctx context.Context) *GetSourceControlFilesMasterWorkspaceParams {
	return &GetSourceControlFilesMasterWorkspaceParams{
		Context: ctx,
	}
}

// NewGetSourceControlFilesMasterWorkspaceParamsWithHTTPClient creates a new GetSourceControlFilesMasterWorkspaceParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSourceControlFilesMasterWorkspaceParamsWithHTTPClient(client *http.Client) *GetSourceControlFilesMasterWorkspaceParams {
	return &GetSourceControlFilesMasterWorkspaceParams{
		HTTPClient: client,
	}
}

/* GetSourceControlFilesMasterWorkspaceParams contains all the parameters to send to the API endpoint
   for the get source control files master workspace operation.

   Typically these are written to a http.Request.
*/
type GetSourceControlFilesMasterWorkspaceParams struct {

	/* ID.

	   workspace id
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get source control files master workspace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSourceControlFilesMasterWorkspaceParams) WithDefaults() *GetSourceControlFilesMasterWorkspaceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get source control files master workspace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSourceControlFilesMasterWorkspaceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get source control files master workspace params
func (o *GetSourceControlFilesMasterWorkspaceParams) WithTimeout(timeout time.Duration) *GetSourceControlFilesMasterWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get source control files master workspace params
func (o *GetSourceControlFilesMasterWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get source control files master workspace params
func (o *GetSourceControlFilesMasterWorkspaceParams) WithContext(ctx context.Context) *GetSourceControlFilesMasterWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get source control files master workspace params
func (o *GetSourceControlFilesMasterWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get source control files master workspace params
func (o *GetSourceControlFilesMasterWorkspaceParams) WithHTTPClient(client *http.Client) *GetSourceControlFilesMasterWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get source control files master workspace params
func (o *GetSourceControlFilesMasterWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get source control files master workspace params
func (o *GetSourceControlFilesMasterWorkspaceParams) WithID(id string) *GetSourceControlFilesMasterWorkspaceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get source control files master workspace params
func (o *GetSourceControlFilesMasterWorkspaceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetSourceControlFilesMasterWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
