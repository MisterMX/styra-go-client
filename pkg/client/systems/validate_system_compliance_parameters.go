// Code generated by go-swagger; DO NOT EDIT.

package systems

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

// NewValidateSystemComplianceParams creates a new ValidateSystemComplianceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewValidateSystemComplianceParams() *ValidateSystemComplianceParams {
	return &ValidateSystemComplianceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewValidateSystemComplianceParamsWithTimeout creates a new ValidateSystemComplianceParams object
// with the ability to set a timeout on a request.
func NewValidateSystemComplianceParamsWithTimeout(timeout time.Duration) *ValidateSystemComplianceParams {
	return &ValidateSystemComplianceParams{
		timeout: timeout,
	}
}

// NewValidateSystemComplianceParamsWithContext creates a new ValidateSystemComplianceParams object
// with the ability to set a context for a request.
func NewValidateSystemComplianceParamsWithContext(ctx context.Context) *ValidateSystemComplianceParams {
	return &ValidateSystemComplianceParams{
		Context: ctx,
	}
}

// NewValidateSystemComplianceParamsWithHTTPClient creates a new ValidateSystemComplianceParams object
// with the ability to set a custom HTTPClient for a request.
func NewValidateSystemComplianceParamsWithHTTPClient(client *http.Client) *ValidateSystemComplianceParams {
	return &ValidateSystemComplianceParams{
		HTTPClient: client,
	}
}

/* ValidateSystemComplianceParams contains all the parameters to send to the API endpoint
   for the validate system compliance operation.

   Typically these are written to a http.Request.
*/
type ValidateSystemComplianceParams struct {

	/* Asyncdelay.

	   set delay of asynchronous response HTTP(202); range [1s - compliance-api-timeout].
	*/
	Asyncdelay *string

	/* Asyncresponse.

	   get asynchronous response; see HTTP(202) Location parameter.
	*/
	Asyncresponse *string

	// Body.
	Body *models.SystemsV1SystemsComplianceRequest

	/* Interval.

	   if set to 'latest', get most recent cached results for specified system.
	*/
	Interval *string

	/* System.

	   system ID
	*/
	System string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the validate system compliance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateSystemComplianceParams) WithDefaults() *ValidateSystemComplianceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the validate system compliance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateSystemComplianceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the validate system compliance params
func (o *ValidateSystemComplianceParams) WithTimeout(timeout time.Duration) *ValidateSystemComplianceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate system compliance params
func (o *ValidateSystemComplianceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate system compliance params
func (o *ValidateSystemComplianceParams) WithContext(ctx context.Context) *ValidateSystemComplianceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate system compliance params
func (o *ValidateSystemComplianceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate system compliance params
func (o *ValidateSystemComplianceParams) WithHTTPClient(client *http.Client) *ValidateSystemComplianceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate system compliance params
func (o *ValidateSystemComplianceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAsyncdelay adds the asyncdelay to the validate system compliance params
func (o *ValidateSystemComplianceParams) WithAsyncdelay(asyncdelay *string) *ValidateSystemComplianceParams {
	o.SetAsyncdelay(asyncdelay)
	return o
}

// SetAsyncdelay adds the asyncdelay to the validate system compliance params
func (o *ValidateSystemComplianceParams) SetAsyncdelay(asyncdelay *string) {
	o.Asyncdelay = asyncdelay
}

// WithAsyncresponse adds the asyncresponse to the validate system compliance params
func (o *ValidateSystemComplianceParams) WithAsyncresponse(asyncresponse *string) *ValidateSystemComplianceParams {
	o.SetAsyncresponse(asyncresponse)
	return o
}

// SetAsyncresponse adds the asyncresponse to the validate system compliance params
func (o *ValidateSystemComplianceParams) SetAsyncresponse(asyncresponse *string) {
	o.Asyncresponse = asyncresponse
}

// WithBody adds the body to the validate system compliance params
func (o *ValidateSystemComplianceParams) WithBody(body *models.SystemsV1SystemsComplianceRequest) *ValidateSystemComplianceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the validate system compliance params
func (o *ValidateSystemComplianceParams) SetBody(body *models.SystemsV1SystemsComplianceRequest) {
	o.Body = body
}

// WithInterval adds the interval to the validate system compliance params
func (o *ValidateSystemComplianceParams) WithInterval(interval *string) *ValidateSystemComplianceParams {
	o.SetInterval(interval)
	return o
}

// SetInterval adds the interval to the validate system compliance params
func (o *ValidateSystemComplianceParams) SetInterval(interval *string) {
	o.Interval = interval
}

// WithSystem adds the system to the validate system compliance params
func (o *ValidateSystemComplianceParams) WithSystem(system string) *ValidateSystemComplianceParams {
	o.SetSystem(system)
	return o
}

// SetSystem adds the system to the validate system compliance params
func (o *ValidateSystemComplianceParams) SetSystem(system string) {
	o.System = system
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateSystemComplianceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param system
	if err := r.SetPathParam("system", o.System); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
