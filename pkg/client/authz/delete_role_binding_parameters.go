// Code generated by go-swagger; DO NOT EDIT.

package authz

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

// NewDeleteRoleBindingParams creates a new DeleteRoleBindingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteRoleBindingParams() *DeleteRoleBindingParams {
	return &DeleteRoleBindingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRoleBindingParamsWithTimeout creates a new DeleteRoleBindingParams object
// with the ability to set a timeout on a request.
func NewDeleteRoleBindingParamsWithTimeout(timeout time.Duration) *DeleteRoleBindingParams {
	return &DeleteRoleBindingParams{
		timeout: timeout,
	}
}

// NewDeleteRoleBindingParamsWithContext creates a new DeleteRoleBindingParams object
// with the ability to set a context for a request.
func NewDeleteRoleBindingParamsWithContext(ctx context.Context) *DeleteRoleBindingParams {
	return &DeleteRoleBindingParams{
		Context: ctx,
	}
}

// NewDeleteRoleBindingParamsWithHTTPClient creates a new DeleteRoleBindingParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteRoleBindingParamsWithHTTPClient(client *http.Client) *DeleteRoleBindingParams {
	return &DeleteRoleBindingParams{
		HTTPClient: client,
	}
}

/* DeleteRoleBindingParams contains all the parameters to send to the API endpoint
   for the delete role binding operation.

   Typically these are written to a http.Request.
*/
type DeleteRoleBindingParams struct {

	/* IfMatch.

	   if set to '*', will return success if not found
	*/
	IfMatch *string

	/* ID.

	   rolebinding ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete role binding params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRoleBindingParams) WithDefaults() *DeleteRoleBindingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete role binding params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRoleBindingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete role binding params
func (o *DeleteRoleBindingParams) WithTimeout(timeout time.Duration) *DeleteRoleBindingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete role binding params
func (o *DeleteRoleBindingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete role binding params
func (o *DeleteRoleBindingParams) WithContext(ctx context.Context) *DeleteRoleBindingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete role binding params
func (o *DeleteRoleBindingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete role binding params
func (o *DeleteRoleBindingParams) WithHTTPClient(client *http.Client) *DeleteRoleBindingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete role binding params
func (o *DeleteRoleBindingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfMatch adds the ifMatch to the delete role binding params
func (o *DeleteRoleBindingParams) WithIfMatch(ifMatch *string) *DeleteRoleBindingParams {
	o.SetIfMatch(ifMatch)
	return o
}

// SetIfMatch adds the ifMatch to the delete role binding params
func (o *DeleteRoleBindingParams) SetIfMatch(ifMatch *string) {
	o.IfMatch = ifMatch
}

// WithID adds the id to the delete role binding params
func (o *DeleteRoleBindingParams) WithID(id string) *DeleteRoleBindingParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete role binding params
func (o *DeleteRoleBindingParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRoleBindingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IfMatch != nil {

		// header param If-Match
		if err := r.SetHeaderParam("If-Match", *o.IfMatch); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
