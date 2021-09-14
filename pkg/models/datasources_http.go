// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DatasourcesHTTP http
// Example: {"ca_certificate":"-----BEGIN CERTIFICATE REQUEST----- MIIB9TCCAWACAQAwgbgxGTAXBgNVBAoMEFF1b1ZhZGlzIExpbWl0ZWQxHDAaBgNV BAsME0RvY3VtZW50IERlcGFydG1lbnQxOTA3BgNVBAMMMFdoeSBhcmUgeW91IGRl Y29kaW5nIG1lPyAgVGhpcyBpcyBvbmx5IGEgdGVzdCEhITERMA8GA1UEBwwISGFt aWx0b24xETAPBgNVBAgMCFBlbWJyb2tlMQswCQYDVQQGEwJCTTEPMA0GCSqGSIb3 DQEJARYAMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCJ9WRanG/fUvcfKiGl EL4aRLjGt537mZ28UU9/3eiJeJznNSOuNLnF+hmabAu7H0LT4K7EdqfF+XUZW/2j RKRYcvOUDGF9A7OjW7UfKk1In3+6QDCi7X34RE161jqoaJjrm/T18TOKcgkkhRzE apQnIDm0Ea/HVzX/PiSOGuertwIDAQABMAsGCSqGSIb3DQEBBQOBgQBzMJdAV4QP Awel8LzGx5uMOshezF/KfP67wJ93UW+N7zXY6AwPgoLj4Kjw+WtU684JL8Dtr9FX ozakE+8p06BpxegR4BR3FMHf6p+0jQxUEAkAyb/mVgm66TyghDGC6/YkiKoZptXQ 98TwDIK/39WEB/V607As+KoYazQG8drorw== -----END CERTIFICATE REQUEST-----\n","category":"http","headers":{"name":"TOKEN","secret_id":"http_token"},"on_premises":false,"polling_interval":"60s","type":"pull"}
//
// swagger:model datasources.HTTP
type DatasourcesHTTP struct {
	DatasourcesCommon

	DatasourcesPoller

	// Custom CA certificate
	CaCertificate string `json:"ca_certificate,omitempty"`

	// Must be `http`
	Category interface{} `json:"category,omitempty"`

	// headers
	Headers []*DatasourcesHTTPHeader `json:"headers"`

	// policy filter
	PolicyFilter string `json:"policy_filter,omitempty"`

	// policy query
	PolicyQuery string `json:"policy_query,omitempty"`

	// Skip TLS verification
	SkipTLSVerification *bool `json:"skip_tls_verification,omitempty"`

	// Must be `pull`
	Type interface{} `json:"type,omitempty"`

	// url
	// Required: true
	URL *string `json:"url"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *DatasourcesHTTP) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 DatasourcesCommon
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.DatasourcesCommon = aO0

	// AO1
	var aO1 DatasourcesPoller
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.DatasourcesPoller = aO1

	// AO2
	var dataAO2 struct {
		CaCertificate string `json:"ca_certificate,omitempty"`

		Category interface{} `json:"category,omitempty"`

		Headers []*DatasourcesHTTPHeader `json:"headers"`

		PolicyFilter string `json:"policy_filter,omitempty"`

		PolicyQuery string `json:"policy_query,omitempty"`

		SkipTLSVerification *bool `json:"skip_tls_verification,omitempty"`

		Type interface{} `json:"type,omitempty"`

		URL *string `json:"url"`
	}
	if err := swag.ReadJSON(raw, &dataAO2); err != nil {
		return err
	}

	m.CaCertificate = dataAO2.CaCertificate

	m.Category = dataAO2.Category

	m.Headers = dataAO2.Headers

	m.PolicyFilter = dataAO2.PolicyFilter

	m.PolicyQuery = dataAO2.PolicyQuery

	m.SkipTLSVerification = dataAO2.SkipTLSVerification

	m.Type = dataAO2.Type

	m.URL = dataAO2.URL

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m DatasourcesHTTP) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	aO0, err := swag.WriteJSON(m.DatasourcesCommon)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.DatasourcesPoller)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	var dataAO2 struct {
		CaCertificate string `json:"ca_certificate,omitempty"`

		Category interface{} `json:"category,omitempty"`

		Headers []*DatasourcesHTTPHeader `json:"headers"`

		PolicyFilter string `json:"policy_filter,omitempty"`

		PolicyQuery string `json:"policy_query,omitempty"`

		SkipTLSVerification *bool `json:"skip_tls_verification,omitempty"`

		Type interface{} `json:"type,omitempty"`

		URL *string `json:"url"`
	}

	dataAO2.CaCertificate = m.CaCertificate

	dataAO2.Category = m.Category

	dataAO2.Headers = m.Headers

	dataAO2.PolicyFilter = m.PolicyFilter

	dataAO2.PolicyQuery = m.PolicyQuery

	dataAO2.SkipTLSVerification = m.SkipTLSVerification

	dataAO2.Type = m.Type

	dataAO2.URL = m.URL

	jsonDataAO2, errAO2 := swag.WriteJSON(dataAO2)
	if errAO2 != nil {
		return nil, errAO2
	}
	_parts = append(_parts, jsonDataAO2)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this datasources HTTP
func (m *DatasourcesHTTP) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with DatasourcesCommon
	if err := m.DatasourcesCommon.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with DatasourcesPoller
	if err := m.DatasourcesPoller.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHeaders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatasourcesHTTP) validateHeaders(formats strfmt.Registry) error {

	if swag.IsZero(m.Headers) { // not required
		return nil
	}

	for i := 0; i < len(m.Headers); i++ {
		if swag.IsZero(m.Headers[i]) { // not required
			continue
		}

		if m.Headers[i] != nil {
			if err := m.Headers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("headers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DatasourcesHTTP) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this datasources HTTP based on the context it is used
func (m *DatasourcesHTTP) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with DatasourcesCommon
	if err := m.DatasourcesCommon.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with DatasourcesPoller
	if err := m.DatasourcesPoller.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHeaders(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatasourcesHTTP) contextValidateHeaders(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Headers); i++ {

		if m.Headers[i] != nil {
			if err := m.Headers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("headers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DatasourcesHTTP) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatasourcesHTTP) UnmarshalBinary(b []byte) error {
	var res DatasourcesHTTP
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}