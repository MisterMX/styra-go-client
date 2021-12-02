// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1DatasourcesPutRequest v1 datasources put request
//
// swagger:model v1.DatasourcesPutRequest
type V1DatasourcesPutRequest struct {

	// Registry ID
	RegistryID string `json:"RegistryId,omitempty"`

	// S3 Bucket
	Bucket string `json:"bucket,omitempty"`

	// Custom CA certificate
	CaCertificate string `json:"ca_certificate,omitempty"`

	// Must be `policy-library`
	// Required: true
	Category *string `json:"category"`

	// content type
	ContentType string `json:"content_type,omitempty"`

	// Secret ID with credentials
	Credentials string `json:"credentials,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Endpoint to S3 storage
	Endpoint string `json:"endpoint,omitempty"`

	// headers
	Headers []*DatasourcesHTTPHeader `json:"headers"`

	// strip out resource properties given their relative path
	Masks map[string][]string `json:"masks,omitempty"`

	// allows to include/exclude namespaces
	Namespaces map[string]bool `json:"namespaces,omitempty"`

	// Must be `true`
	// Required: true
	OnPremises *bool `json:"on_premises"`

	// S3 Path within a Bucket
	Path string `json:"path,omitempty"`

	// path regexp
	PathRegexp *string `json:"path_regexp,omitempty"`

	// policy filter
	PolicyFilter string `json:"policy_filter,omitempty"`

	// policy query
	PolicyQuery string `json:"policy_query,omitempty"`

	// polling interval
	PollingInterval *string `json:"polling_interval,omitempty"`

	// requests per second
	RateLimit *float64 `json:"rate_limit,omitempty"`

	// reference
	Reference *string `json:"reference,omitempty"`

	// AWS region
	Region string `json:"region,omitempty"`

	// field selectors per resource [group]
	Selectors map[string]string `json:"selectors,omitempty"`

	// Skip TLS verification
	SkipTLSVerification *bool `json:"skip_tls_verification,omitempty"`

	// ssh credentials
	SSHCredentials *V1DatasourcesPutRequestSSHCredentials `json:"ssh_credentials,omitempty"`

	// timeout
	Timeout *string `json:"timeout,omitempty"`

	// Must be `pull`
	// Required: true
	// Enum: [pull push]
	Type *string `json:"type"`

	// Git URL
	URL string `json:"url,omitempty"`
}

// Validate validates this v1 datasources put request
func (m *V1DatasourcesPutRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHeaders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOnPremises(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSSHCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1DatasourcesPutRequest) validateCategory(formats strfmt.Registry) error {

	if err := validate.Required("category", "body", m.Category); err != nil {
		return err
	}

	return nil
}

func (m *V1DatasourcesPutRequest) validateHeaders(formats strfmt.Registry) error {
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

func (m *V1DatasourcesPutRequest) validateOnPremises(formats strfmt.Registry) error {

	if err := validate.Required("on_premises", "body", m.OnPremises); err != nil {
		return err
	}

	return nil
}

func (m *V1DatasourcesPutRequest) validateSSHCredentials(formats strfmt.Registry) error {
	if swag.IsZero(m.SSHCredentials) { // not required
		return nil
	}

	if m.SSHCredentials != nil {
		if err := m.SSHCredentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ssh_credentials")
			}
			return err
		}
	}

	return nil
}

var v1DatasourcesPutRequestTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["pull","push"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1DatasourcesPutRequestTypeTypePropEnum = append(v1DatasourcesPutRequestTypeTypePropEnum, v)
	}
}

const (

	// V1DatasourcesPutRequestTypePull captures enum value "pull"
	V1DatasourcesPutRequestTypePull string = "pull"

	// V1DatasourcesPutRequestTypePush captures enum value "push"
	V1DatasourcesPutRequestTypePush string = "push"
)

// prop value enum
func (m *V1DatasourcesPutRequest) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, v1DatasourcesPutRequestTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V1DatasourcesPutRequest) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 datasources put request based on the context it is used
func (m *V1DatasourcesPutRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHeaders(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSSHCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1DatasourcesPutRequest) contextValidateHeaders(ctx context.Context, formats strfmt.Registry) error {

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

func (m *V1DatasourcesPutRequest) contextValidateSSHCredentials(ctx context.Context, formats strfmt.Registry) error {

	if m.SSHCredentials != nil {
		if err := m.SSHCredentials.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ssh_credentials")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1DatasourcesPutRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1DatasourcesPutRequest) UnmarshalBinary(b []byte) error {
	var res V1DatasourcesPutRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// V1DatasourcesPutRequestSSHCredentials v1 datasources put request SSH credentials
//
// swagger:model V1DatasourcesPutRequestSSHCredentials
type V1DatasourcesPutRequestSSHCredentials struct {

	// Secret ID with passphrase
	Passphrase string `json:"passphrase,omitempty"`

	// Secret ID with private key
	// Required: true
	PrivateKey *string `json:"private_key"`
}

// Validate validates this v1 datasources put request SSH credentials
func (m *V1DatasourcesPutRequestSSHCredentials) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrivateKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1DatasourcesPutRequestSSHCredentials) validatePrivateKey(formats strfmt.Registry) error {

	if err := validate.Required("ssh_credentials"+"."+"private_key", "body", m.PrivateKey); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 datasources put request SSH credentials based on context it is used
func (m *V1DatasourcesPutRequestSSHCredentials) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1DatasourcesPutRequestSSHCredentials) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1DatasourcesPutRequestSSHCredentials) UnmarshalBinary(b []byte) error {
	var res V1DatasourcesPutRequestSSHCredentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
