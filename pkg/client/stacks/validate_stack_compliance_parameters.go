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

// NewValidateStackComplianceParams creates a new ValidateStackComplianceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewValidateStackComplianceParams() *ValidateStackComplianceParams {
	return &ValidateStackComplianceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewValidateStackComplianceParamsWithTimeout creates a new ValidateStackComplianceParams object
// with the ability to set a timeout on a request.
func NewValidateStackComplianceParamsWithTimeout(timeout time.Duration) *ValidateStackComplianceParams {
	return &ValidateStackComplianceParams{
		timeout: timeout,
	}
}

// NewValidateStackComplianceParamsWithContext creates a new ValidateStackComplianceParams object
// with the ability to set a context for a request.
func NewValidateStackComplianceParamsWithContext(ctx context.Context) *ValidateStackComplianceParams {
	return &ValidateStackComplianceParams{
		Context: ctx,
	}
}

// NewValidateStackComplianceParamsWithHTTPClient creates a new ValidateStackComplianceParams object
// with the ability to set a custom HTTPClient for a request.
func NewValidateStackComplianceParamsWithHTTPClient(client *http.Client) *ValidateStackComplianceParams {
	return &ValidateStackComplianceParams{
		HTTPClient: client,
	}
}

/* ValidateStackComplianceParams contains all the parameters to send to the API endpoint
   for the validate stack compliance operation.

   Typically these are written to a http.Request.
*/
type ValidateStackComplianceParams struct {

	/* Asyncdelay.

	   set delay of asynchronous response HTTP(202); range [1s - compliance-api-timeout].
	*/
	Asyncdelay *string

	/* Asyncresponse.

	   get asynchronous response; see HTTP(202) Location parameter
	*/
	Asyncresponse *string

	// Body.
	Body *models.StacksV1StacksComplianceRequest

	/* Interval.

	   if set to 'latest', get most recent cached results for specified stack.
	*/
	Interval *string

	/* Stack.

	   stack id
	*/
	Stack string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the validate stack compliance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateStackComplianceParams) WithDefaults() *ValidateStackComplianceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the validate stack compliance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateStackComplianceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the validate stack compliance params
func (o *ValidateStackComplianceParams) WithTimeout(timeout time.Duration) *ValidateStackComplianceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate stack compliance params
func (o *ValidateStackComplianceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate stack compliance params
func (o *ValidateStackComplianceParams) WithContext(ctx context.Context) *ValidateStackComplianceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate stack compliance params
func (o *ValidateStackComplianceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate stack compliance params
func (o *ValidateStackComplianceParams) WithHTTPClient(client *http.Client) *ValidateStackComplianceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate stack compliance params
func (o *ValidateStackComplianceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAsyncdelay adds the asyncdelay to the validate stack compliance params
func (o *ValidateStackComplianceParams) WithAsyncdelay(asyncdelay *string) *ValidateStackComplianceParams {
	o.SetAsyncdelay(asyncdelay)
	return o
}

// SetAsyncdelay adds the asyncdelay to the validate stack compliance params
func (o *ValidateStackComplianceParams) SetAsyncdelay(asyncdelay *string) {
	o.Asyncdelay = asyncdelay
}

// WithAsyncresponse adds the asyncresponse to the validate stack compliance params
func (o *ValidateStackComplianceParams) WithAsyncresponse(asyncresponse *string) *ValidateStackComplianceParams {
	o.SetAsyncresponse(asyncresponse)
	return o
}

// SetAsyncresponse adds the asyncresponse to the validate stack compliance params
func (o *ValidateStackComplianceParams) SetAsyncresponse(asyncresponse *string) {
	o.Asyncresponse = asyncresponse
}

// WithBody adds the body to the validate stack compliance params
func (o *ValidateStackComplianceParams) WithBody(body *models.StacksV1StacksComplianceRequest) *ValidateStackComplianceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the validate stack compliance params
func (o *ValidateStackComplianceParams) SetBody(body *models.StacksV1StacksComplianceRequest) {
	o.Body = body
}

// WithInterval adds the interval to the validate stack compliance params
func (o *ValidateStackComplianceParams) WithInterval(interval *string) *ValidateStackComplianceParams {
	o.SetInterval(interval)
	return o
}

// SetInterval adds the interval to the validate stack compliance params
func (o *ValidateStackComplianceParams) SetInterval(interval *string) {
	o.Interval = interval
}

// WithStack adds the stack to the validate stack compliance params
func (o *ValidateStackComplianceParams) WithStack(stack string) *ValidateStackComplianceParams {
	o.SetStack(stack)
	return o
}

// SetStack adds the stack to the validate stack compliance params
func (o *ValidateStackComplianceParams) SetStack(stack string) {
	o.Stack = stack
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateStackComplianceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Asyncdelay != nil {

		// query param asyncdelay
		var qrAsyncdelay string

		if o.Asyncdelay != nil {
			qrAsyncdelay = *o.Asyncdelay
		}
		qAsyncdelay := qrAsyncdelay
		if qAsyncdelay != "" {

			if err := r.SetQueryParam("asyncdelay", qAsyncdelay); err != nil {
				return err
			}
		}
	}

	if o.Asyncresponse != nil {

		// query param asyncresponse
		var qrAsyncresponse string

		if o.Asyncresponse != nil {
			qrAsyncresponse = *o.Asyncresponse
		}
		qAsyncresponse := qrAsyncresponse
		if qAsyncresponse != "" {

			if err := r.SetQueryParam("asyncresponse", qAsyncresponse); err != nil {
				return err
			}
		}
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.Interval != nil {

		// query param interval
		var qrInterval string

		if o.Interval != nil {
			qrInterval = *o.Interval
		}
		qInterval := qrInterval
		if qInterval != "" {

			if err := r.SetQueryParam("interval", qInterval); err != nil {
				return err
			}
		}
	}

	// path param stack
	if err := r.SetPathParam("stack", o.Stack); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
