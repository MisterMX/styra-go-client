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

// NewGetRoleBindingV1Params creates a new GetRoleBindingV1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRoleBindingV1Params() *GetRoleBindingV1Params {
	return &GetRoleBindingV1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRoleBindingV1ParamsWithTimeout creates a new GetRoleBindingV1Params object
// with the ability to set a timeout on a request.
func NewGetRoleBindingV1ParamsWithTimeout(timeout time.Duration) *GetRoleBindingV1Params {
	return &GetRoleBindingV1Params{
		timeout: timeout,
	}
}

// NewGetRoleBindingV1ParamsWithContext creates a new GetRoleBindingV1Params object
// with the ability to set a context for a request.
func NewGetRoleBindingV1ParamsWithContext(ctx context.Context) *GetRoleBindingV1Params {
	return &GetRoleBindingV1Params{
		Context: ctx,
	}
}

// NewGetRoleBindingV1ParamsWithHTTPClient creates a new GetRoleBindingV1Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetRoleBindingV1ParamsWithHTTPClient(client *http.Client) *GetRoleBindingV1Params {
	return &GetRoleBindingV1Params{
		HTTPClient: client,
	}
}

/* GetRoleBindingV1Params contains all the parameters to send to the API endpoint
   for the get role binding v1 operation.

   Typically these are written to a http.Request.
*/
type GetRoleBindingV1Params struct {

	/* Resource.

	   resource id
	*/
	Resource string

	/* Resourcetype.

	   resource type
	*/
	Resourcetype string

	/* Rolebinding.

	   role binding id
	*/
	Rolebinding string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get role binding v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRoleBindingV1Params) WithDefaults() *GetRoleBindingV1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get role binding v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRoleBindingV1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get role binding v1 params
func (o *GetRoleBindingV1Params) WithTimeout(timeout time.Duration) *GetRoleBindingV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get role binding v1 params
func (o *GetRoleBindingV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get role binding v1 params
func (o *GetRoleBindingV1Params) WithContext(ctx context.Context) *GetRoleBindingV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get role binding v1 params
func (o *GetRoleBindingV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get role binding v1 params
func (o *GetRoleBindingV1Params) WithHTTPClient(client *http.Client) *GetRoleBindingV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get role binding v1 params
func (o *GetRoleBindingV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResource adds the resource to the get role binding v1 params
func (o *GetRoleBindingV1Params) WithResource(resource string) *GetRoleBindingV1Params {
	o.SetResource(resource)
	return o
}

// SetResource adds the resource to the get role binding v1 params
func (o *GetRoleBindingV1Params) SetResource(resource string) {
	o.Resource = resource
}

// WithResourcetype adds the resourcetype to the get role binding v1 params
func (o *GetRoleBindingV1Params) WithResourcetype(resourcetype string) *GetRoleBindingV1Params {
	o.SetResourcetype(resourcetype)
	return o
}

// SetResourcetype adds the resourcetype to the get role binding v1 params
func (o *GetRoleBindingV1Params) SetResourcetype(resourcetype string) {
	o.Resourcetype = resourcetype
}

// WithRolebinding adds the rolebinding to the get role binding v1 params
func (o *GetRoleBindingV1Params) WithRolebinding(rolebinding string) *GetRoleBindingV1Params {
	o.SetRolebinding(rolebinding)
	return o
}

// SetRolebinding adds the rolebinding to the get role binding v1 params
func (o *GetRoleBindingV1Params) SetRolebinding(rolebinding string) {
	o.Rolebinding = rolebinding
}

// WriteToRequest writes these params to a swagger request
func (o *GetRoleBindingV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param resource
	if err := r.SetPathParam("resource", o.Resource); err != nil {
		return err
	}

	// path param resourcetype
	if err := r.SetPathParam("resourcetype", o.Resourcetype); err != nil {
		return err
	}

	// path param rolebinding
	if err := r.SetPathParam("rolebinding", o.Rolebinding); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}