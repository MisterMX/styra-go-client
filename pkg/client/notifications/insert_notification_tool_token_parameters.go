// Code generated by go-swagger; DO NOT EDIT.

package notifications

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

// NewInsertNotificationToolTokenParams creates a new InsertNotificationToolTokenParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInsertNotificationToolTokenParams() *InsertNotificationToolTokenParams {
	return &InsertNotificationToolTokenParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInsertNotificationToolTokenParamsWithTimeout creates a new InsertNotificationToolTokenParams object
// with the ability to set a timeout on a request.
func NewInsertNotificationToolTokenParamsWithTimeout(timeout time.Duration) *InsertNotificationToolTokenParams {
	return &InsertNotificationToolTokenParams{
		timeout: timeout,
	}
}

// NewInsertNotificationToolTokenParamsWithContext creates a new InsertNotificationToolTokenParams object
// with the ability to set a context for a request.
func NewInsertNotificationToolTokenParamsWithContext(ctx context.Context) *InsertNotificationToolTokenParams {
	return &InsertNotificationToolTokenParams{
		Context: ctx,
	}
}

// NewInsertNotificationToolTokenParamsWithHTTPClient creates a new InsertNotificationToolTokenParams object
// with the ability to set a custom HTTPClient for a request.
func NewInsertNotificationToolTokenParamsWithHTTPClient(client *http.Client) *InsertNotificationToolTokenParams {
	return &InsertNotificationToolTokenParams{
		HTTPClient: client,
	}
}

/* InsertNotificationToolTokenParams contains all the parameters to send to the API endpoint
   for the insert notification tool token operation.

   Typically these are written to a http.Request.
*/
type InsertNotificationToolTokenParams struct {

	// Body.
	Body *models.V1NotificationToolTokenRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the insert notification tool token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InsertNotificationToolTokenParams) WithDefaults() *InsertNotificationToolTokenParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the insert notification tool token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InsertNotificationToolTokenParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the insert notification tool token params
func (o *InsertNotificationToolTokenParams) WithTimeout(timeout time.Duration) *InsertNotificationToolTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the insert notification tool token params
func (o *InsertNotificationToolTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the insert notification tool token params
func (o *InsertNotificationToolTokenParams) WithContext(ctx context.Context) *InsertNotificationToolTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the insert notification tool token params
func (o *InsertNotificationToolTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the insert notification tool token params
func (o *InsertNotificationToolTokenParams) WithHTTPClient(client *http.Client) *InsertNotificationToolTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the insert notification tool token params
func (o *InsertNotificationToolTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the insert notification tool token params
func (o *InsertNotificationToolTokenParams) WithBody(body *models.V1NotificationToolTokenRequest) *InsertNotificationToolTokenParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the insert notification tool token params
func (o *InsertNotificationToolTokenParams) SetBody(body *models.V1NotificationToolTokenRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *InsertNotificationToolTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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