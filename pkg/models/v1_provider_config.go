// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1ProviderConfig v1 provider config
//
// swagger:model v1.ProviderConfig
type V1ProviderConfig struct {

	// allow idp initiated
	// Required: true
	AllowIdpInitiated *bool `json:"allow_idp_initiated"`

	// allowed domains
	// Required: true
	AllowedDomains []string `json:"allowed_domains"`

	// auth url
	// Required: true
	AuthURL *string `json:"auth_url"`

	// client id
	// Required: true
	ClientID *string `json:"client_id"`

	// client secret
	// Required: true
	ClientSecret *string `json:"client_secret"`

	// email attribute
	// Required: true
	EmailAttribute *string `json:"email_attribute"`

	// enabled
	// Required: true
	Enabled *bool `json:"enabled"`

	// id
	// Required: true
	ID *string `json:"id"`

	// issuer url
	// Required: true
	IssuerURL *string `json:"issuer_url"`

	// jit
	// Required: true
	Jit *bool `json:"jit"`

	// key certificate
	// Required: true
	KeyCertificate *string `json:"key_certificate"`

	// metadata
	// Required: true
	Metadata *string `json:"metadata"`

	// proxy url
	// Required: true
	ProxyURL *string `json:"proxy_url"`

	// scopes
	// Required: true
	Scopes []string `json:"scopes"`

	// token url
	// Required: true
	TokenURL *string `json:"token_url"`

	// type
	// Required: true
	Type *string `json:"type"`

	// user info url
	// Required: true
	UserInfoURL *string `json:"user_info_url"`
}

// Validate validates this v1 provider config
func (m *V1ProviderConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowIdpInitiated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAllowedDomains(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientSecret(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmailAttribute(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssuerURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeyCertificate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProxyURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScopes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTokenURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserInfoURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ProviderConfig) validateAllowIdpInitiated(formats strfmt.Registry) error {

	if err := validate.Required("allow_idp_initiated", "body", m.AllowIdpInitiated); err != nil {
		return err
	}

	return nil
}

func (m *V1ProviderConfig) validateAllowedDomains(formats strfmt.Registry) error {

	if err := validate.Required("allowed_domains", "body", m.AllowedDomains); err != nil {
		return err
	}

	return nil
}

func (m *V1ProviderConfig) validateAuthURL(formats strfmt.Registry) error {

	if err := validate.Required("auth_url", "body", m.AuthURL); err != nil {
		return err
	}

	return nil
}

func (m *V1ProviderConfig) validateClientID(formats strfmt.Registry) error {

	if err := validate.Required("client_id", "body", m.ClientID); err != nil {
		return err
	}

	return nil
}

func (m *V1ProviderConfig) validateClientSecret(formats strfmt.Registry) error {

	if err := validate.Required("client_secret", "body", m.ClientSecret); err != nil {
		return err
	}

	return nil
}

func (m *V1ProviderConfig) validateEmailAttribute(formats strfmt.Registry) error {

	if err := validate.Required("email_attribute", "body", m.EmailAttribute); err != nil {
		return err
	}

	return nil
}

func (m *V1ProviderConfig) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *V1ProviderConfig) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *V1ProviderConfig) validateIssuerURL(formats strfmt.Registry) error {

	if err := validate.Required("issuer_url", "body", m.IssuerURL); err != nil {
		return err
	}

	return nil
}

func (m *V1ProviderConfig) validateJit(formats strfmt.Registry) error {

	if err := validate.Required("jit", "body", m.Jit); err != nil {
		return err
	}

	return nil
}

func (m *V1ProviderConfig) validateKeyCertificate(formats strfmt.Registry) error {

	if err := validate.Required("key_certificate", "body", m.KeyCertificate); err != nil {
		return err
	}

	return nil
}

func (m *V1ProviderConfig) validateMetadata(formats strfmt.Registry) error {

	if err := validate.Required("metadata", "body", m.Metadata); err != nil {
		return err
	}

	return nil
}

func (m *V1ProviderConfig) validateProxyURL(formats strfmt.Registry) error {

	if err := validate.Required("proxy_url", "body", m.ProxyURL); err != nil {
		return err
	}

	return nil
}

func (m *V1ProviderConfig) validateScopes(formats strfmt.Registry) error {

	if err := validate.Required("scopes", "body", m.Scopes); err != nil {
		return err
	}

	return nil
}

func (m *V1ProviderConfig) validateTokenURL(formats strfmt.Registry) error {

	if err := validate.Required("token_url", "body", m.TokenURL); err != nil {
		return err
	}

	return nil
}

func (m *V1ProviderConfig) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *V1ProviderConfig) validateUserInfoURL(formats strfmt.Registry) error {

	if err := validate.Required("user_info_url", "body", m.UserInfoURL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 provider config based on context it is used
func (m *V1ProviderConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ProviderConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ProviderConfig) UnmarshalBinary(b []byte) error {
	var res V1ProviderConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
