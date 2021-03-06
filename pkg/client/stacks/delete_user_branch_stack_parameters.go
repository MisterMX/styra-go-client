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
)

// NewDeleteUserBranchStackParams creates a new DeleteUserBranchStackParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteUserBranchStackParams() *DeleteUserBranchStackParams {
	return &DeleteUserBranchStackParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUserBranchStackParamsWithTimeout creates a new DeleteUserBranchStackParams object
// with the ability to set a timeout on a request.
func NewDeleteUserBranchStackParamsWithTimeout(timeout time.Duration) *DeleteUserBranchStackParams {
	return &DeleteUserBranchStackParams{
		timeout: timeout,
	}
}

// NewDeleteUserBranchStackParamsWithContext creates a new DeleteUserBranchStackParams object
// with the ability to set a context for a request.
func NewDeleteUserBranchStackParamsWithContext(ctx context.Context) *DeleteUserBranchStackParams {
	return &DeleteUserBranchStackParams{
		Context: ctx,
	}
}

// NewDeleteUserBranchStackParamsWithHTTPClient creates a new DeleteUserBranchStackParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteUserBranchStackParamsWithHTTPClient(client *http.Client) *DeleteUserBranchStackParams {
	return &DeleteUserBranchStackParams{
		HTTPClient: client,
	}
}

/* DeleteUserBranchStackParams contains all the parameters to send to the API endpoint
   for the delete user branch stack operation.

   Typically these are written to a http.Request.
*/
type DeleteUserBranchStackParams struct {

	/* ID.

	   stack id
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete user branch stack params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUserBranchStackParams) WithDefaults() *DeleteUserBranchStackParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete user branch stack params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUserBranchStackParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete user branch stack params
func (o *DeleteUserBranchStackParams) WithTimeout(timeout time.Duration) *DeleteUserBranchStackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete user branch stack params
func (o *DeleteUserBranchStackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete user branch stack params
func (o *DeleteUserBranchStackParams) WithContext(ctx context.Context) *DeleteUserBranchStackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete user branch stack params
func (o *DeleteUserBranchStackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete user branch stack params
func (o *DeleteUserBranchStackParams) WithHTTPClient(client *http.Client) *DeleteUserBranchStackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete user branch stack params
func (o *DeleteUserBranchStackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete user branch stack params
func (o *DeleteUserBranchStackParams) WithID(id string) *DeleteUserBranchStackParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete user branch stack params
func (o *DeleteUserBranchStackParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUserBranchStackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
