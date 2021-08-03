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

	"github.com/mistermx/styra-go-client/pkg/models"
)

// NewQueryRoleBindingsParams creates a new QueryRoleBindingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQueryRoleBindingsParams() *QueryRoleBindingsParams {
	return &QueryRoleBindingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQueryRoleBindingsParamsWithTimeout creates a new QueryRoleBindingsParams object
// with the ability to set a timeout on a request.
func NewQueryRoleBindingsParamsWithTimeout(timeout time.Duration) *QueryRoleBindingsParams {
	return &QueryRoleBindingsParams{
		timeout: timeout,
	}
}

// NewQueryRoleBindingsParamsWithContext creates a new QueryRoleBindingsParams object
// with the ability to set a context for a request.
func NewQueryRoleBindingsParamsWithContext(ctx context.Context) *QueryRoleBindingsParams {
	return &QueryRoleBindingsParams{
		Context: ctx,
	}
}

// NewQueryRoleBindingsParamsWithHTTPClient creates a new QueryRoleBindingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewQueryRoleBindingsParamsWithHTTPClient(client *http.Client) *QueryRoleBindingsParams {
	return &QueryRoleBindingsParams{
		HTTPClient: client,
	}
}

/* QueryRoleBindingsParams contains all the parameters to send to the API endpoint
   for the query role bindings operation.

   Typically these are written to a http.Request.
*/
type QueryRoleBindingsParams struct {

	// Body.
	Body *models.V2RoleBindingsPostRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the query role bindings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryRoleBindingsParams) WithDefaults() *QueryRoleBindingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the query role bindings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryRoleBindingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the query role bindings params
func (o *QueryRoleBindingsParams) WithTimeout(timeout time.Duration) *QueryRoleBindingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query role bindings params
func (o *QueryRoleBindingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query role bindings params
func (o *QueryRoleBindingsParams) WithContext(ctx context.Context) *QueryRoleBindingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query role bindings params
func (o *QueryRoleBindingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query role bindings params
func (o *QueryRoleBindingsParams) WithHTTPClient(client *http.Client) *QueryRoleBindingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query role bindings params
func (o *QueryRoleBindingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the query role bindings params
func (o *QueryRoleBindingsParams) WithBody(body *models.V2RoleBindingsPostRequest) *QueryRoleBindingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the query role bindings params
func (o *QueryRoleBindingsParams) SetBody(body *models.V2RoleBindingsPostRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *QueryRoleBindingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
