// Code generated by go-swagger; DO NOT EDIT.

package decisions

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

// NewSearchDecisionLogsPostParams creates a new SearchDecisionLogsPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchDecisionLogsPostParams() *SearchDecisionLogsPostParams {
	return &SearchDecisionLogsPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchDecisionLogsPostParamsWithTimeout creates a new SearchDecisionLogsPostParams object
// with the ability to set a timeout on a request.
func NewSearchDecisionLogsPostParamsWithTimeout(timeout time.Duration) *SearchDecisionLogsPostParams {
	return &SearchDecisionLogsPostParams{
		timeout: timeout,
	}
}

// NewSearchDecisionLogsPostParamsWithContext creates a new SearchDecisionLogsPostParams object
// with the ability to set a context for a request.
func NewSearchDecisionLogsPostParamsWithContext(ctx context.Context) *SearchDecisionLogsPostParams {
	return &SearchDecisionLogsPostParams{
		Context: ctx,
	}
}

// NewSearchDecisionLogsPostParamsWithHTTPClient creates a new SearchDecisionLogsPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchDecisionLogsPostParamsWithHTTPClient(client *http.Client) *SearchDecisionLogsPostParams {
	return &SearchDecisionLogsPostParams{
		HTTPClient: client,
	}
}

/* SearchDecisionLogsPostParams contains all the parameters to send to the API endpoint
   for the search decision logs post operation.

   Typically these are written to a http.Request.
*/
type SearchDecisionLogsPostParams struct {

	// Body.
	Body *models.V1DecisionsGetRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search decision logs post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchDecisionLogsPostParams) WithDefaults() *SearchDecisionLogsPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search decision logs post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchDecisionLogsPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search decision logs post params
func (o *SearchDecisionLogsPostParams) WithTimeout(timeout time.Duration) *SearchDecisionLogsPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search decision logs post params
func (o *SearchDecisionLogsPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search decision logs post params
func (o *SearchDecisionLogsPostParams) WithContext(ctx context.Context) *SearchDecisionLogsPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search decision logs post params
func (o *SearchDecisionLogsPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search decision logs post params
func (o *SearchDecisionLogsPostParams) WithHTTPClient(client *http.Client) *SearchDecisionLogsPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search decision logs post params
func (o *SearchDecisionLogsPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the search decision logs post params
func (o *SearchDecisionLogsPostParams) WithBody(body *models.V1DecisionsGetRequest) *SearchDecisionLogsPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the search decision logs post params
func (o *SearchDecisionLogsPostParams) SetBody(body *models.V1DecisionsGetRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SearchDecisionLogsPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
