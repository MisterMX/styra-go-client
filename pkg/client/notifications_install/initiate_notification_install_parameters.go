// Code generated by go-swagger; DO NOT EDIT.

package notifications_install

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

// NewInitiateNotificationInstallParams creates a new InitiateNotificationInstallParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInitiateNotificationInstallParams() *InitiateNotificationInstallParams {
	return &InitiateNotificationInstallParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInitiateNotificationInstallParamsWithTimeout creates a new InitiateNotificationInstallParams object
// with the ability to set a timeout on a request.
func NewInitiateNotificationInstallParamsWithTimeout(timeout time.Duration) *InitiateNotificationInstallParams {
	return &InitiateNotificationInstallParams{
		timeout: timeout,
	}
}

// NewInitiateNotificationInstallParamsWithContext creates a new InitiateNotificationInstallParams object
// with the ability to set a context for a request.
func NewInitiateNotificationInstallParamsWithContext(ctx context.Context) *InitiateNotificationInstallParams {
	return &InitiateNotificationInstallParams{
		Context: ctx,
	}
}

// NewInitiateNotificationInstallParamsWithHTTPClient creates a new InitiateNotificationInstallParams object
// with the ability to set a custom HTTPClient for a request.
func NewInitiateNotificationInstallParamsWithHTTPClient(client *http.Client) *InitiateNotificationInstallParams {
	return &InitiateNotificationInstallParams{
		HTTPClient: client,
	}
}

/* InitiateNotificationInstallParams contains all the parameters to send to the API endpoint
   for the initiate notification install operation.

   Typically these are written to a http.Request.
*/
type InitiateNotificationInstallParams struct {

	/* RedirectURL.

	   the landing page when OAuth is successfully done.
	*/
	RedirectURL *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the initiate notification install params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InitiateNotificationInstallParams) WithDefaults() *InitiateNotificationInstallParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the initiate notification install params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InitiateNotificationInstallParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the initiate notification install params
func (o *InitiateNotificationInstallParams) WithTimeout(timeout time.Duration) *InitiateNotificationInstallParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the initiate notification install params
func (o *InitiateNotificationInstallParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the initiate notification install params
func (o *InitiateNotificationInstallParams) WithContext(ctx context.Context) *InitiateNotificationInstallParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the initiate notification install params
func (o *InitiateNotificationInstallParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the initiate notification install params
func (o *InitiateNotificationInstallParams) WithHTTPClient(client *http.Client) *InitiateNotificationInstallParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the initiate notification install params
func (o *InitiateNotificationInstallParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRedirectURL adds the redirectURL to the initiate notification install params
func (o *InitiateNotificationInstallParams) WithRedirectURL(redirectURL *string) *InitiateNotificationInstallParams {
	o.SetRedirectURL(redirectURL)
	return o
}

// SetRedirectURL adds the redirectUrl to the initiate notification install params
func (o *InitiateNotificationInstallParams) SetRedirectURL(redirectURL *string) {
	o.RedirectURL = redirectURL
}

// WriteToRequest writes these params to a swagger request
func (o *InitiateNotificationInstallParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.RedirectURL != nil {

		// query param redirect_url
		var qrRedirectURL string

		if o.RedirectURL != nil {
			qrRedirectURL = *o.RedirectURL
		}
		qRedirectURL := qrRedirectURL
		if qRedirectURL != "" {

			if err := r.SetQueryParam("redirect_url", qRedirectURL); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}