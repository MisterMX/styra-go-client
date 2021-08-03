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

// NewGetSourceControlFilesBranchWorkspaceParams creates a new GetSourceControlFilesBranchWorkspaceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSourceControlFilesBranchWorkspaceParams() *GetSourceControlFilesBranchWorkspaceParams {
	return &GetSourceControlFilesBranchWorkspaceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSourceControlFilesBranchWorkspaceParamsWithTimeout creates a new GetSourceControlFilesBranchWorkspaceParams object
// with the ability to set a timeout on a request.
func NewGetSourceControlFilesBranchWorkspaceParamsWithTimeout(timeout time.Duration) *GetSourceControlFilesBranchWorkspaceParams {
	return &GetSourceControlFilesBranchWorkspaceParams{
		timeout: timeout,
	}
}

// NewGetSourceControlFilesBranchWorkspaceParamsWithContext creates a new GetSourceControlFilesBranchWorkspaceParams object
// with the ability to set a context for a request.
func NewGetSourceControlFilesBranchWorkspaceParamsWithContext(ctx context.Context) *GetSourceControlFilesBranchWorkspaceParams {
	return &GetSourceControlFilesBranchWorkspaceParams{
		Context: ctx,
	}
}

// NewGetSourceControlFilesBranchWorkspaceParamsWithHTTPClient creates a new GetSourceControlFilesBranchWorkspaceParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSourceControlFilesBranchWorkspaceParamsWithHTTPClient(client *http.Client) *GetSourceControlFilesBranchWorkspaceParams {
	return &GetSourceControlFilesBranchWorkspaceParams{
		HTTPClient: client,
	}
}

/* GetSourceControlFilesBranchWorkspaceParams contains all the parameters to send to the API endpoint
   for the get source control files branch workspace operation.

   Typically these are written to a http.Request.
*/
type GetSourceControlFilesBranchWorkspaceParams struct {

	/* ID.

	   workspace id
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get source control files branch workspace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSourceControlFilesBranchWorkspaceParams) WithDefaults() *GetSourceControlFilesBranchWorkspaceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get source control files branch workspace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSourceControlFilesBranchWorkspaceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get source control files branch workspace params
func (o *GetSourceControlFilesBranchWorkspaceParams) WithTimeout(timeout time.Duration) *GetSourceControlFilesBranchWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get source control files branch workspace params
func (o *GetSourceControlFilesBranchWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get source control files branch workspace params
func (o *GetSourceControlFilesBranchWorkspaceParams) WithContext(ctx context.Context) *GetSourceControlFilesBranchWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get source control files branch workspace params
func (o *GetSourceControlFilesBranchWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get source control files branch workspace params
func (o *GetSourceControlFilesBranchWorkspaceParams) WithHTTPClient(client *http.Client) *GetSourceControlFilesBranchWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get source control files branch workspace params
func (o *GetSourceControlFilesBranchWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get source control files branch workspace params
func (o *GetSourceControlFilesBranchWorkspaceParams) WithID(id string) *GetSourceControlFilesBranchWorkspaceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get source control files branch workspace params
func (o *GetSourceControlFilesBranchWorkspaceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetSourceControlFilesBranchWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
