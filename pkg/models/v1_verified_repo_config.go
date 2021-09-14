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

// V1VerifiedRepoConfig v1 verified repo config
//
// swagger:model v1.VerifiedRepoConfig
type V1VerifiedRepoConfig struct {

	// Credentials are looked under the key <name>/<creds>
	// Required: true
	Credentials *string `json:"credentials"`

	// Path to limit the import to
	// Required: true
	Path *string `json:"path"`

	// Remote reference, defaults to refs/heads/master
	// Required: true
	Reference *string `json:"reference"`

	// sha
	// Required: true
	Sha *string `json:"sha"`

	// SSHCredentials including ssh private key and passphrase for ssh-based auth
	SSHCredentials *V1SSHCredentials `json:"ssh_credentials,omitempty"`

	// Repository URL
	// Required: true
	URL *string `json:"url"`
}

// Validate validates this v1 verified repo config
func (m *V1VerifiedRepoConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSha(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSSHCredentials(formats); err != nil {
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

func (m *V1VerifiedRepoConfig) validateCredentials(formats strfmt.Registry) error {

	if err := validate.Required("credentials", "body", m.Credentials); err != nil {
		return err
	}

	return nil
}

func (m *V1VerifiedRepoConfig) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

func (m *V1VerifiedRepoConfig) validateReference(formats strfmt.Registry) error {

	if err := validate.Required("reference", "body", m.Reference); err != nil {
		return err
	}

	return nil
}

func (m *V1VerifiedRepoConfig) validateSha(formats strfmt.Registry) error {

	if err := validate.Required("sha", "body", m.Sha); err != nil {
		return err
	}

	return nil
}

func (m *V1VerifiedRepoConfig) validateSSHCredentials(formats strfmt.Registry) error {
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

func (m *V1VerifiedRepoConfig) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 verified repo config based on the context it is used
func (m *V1VerifiedRepoConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSSHCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VerifiedRepoConfig) contextValidateSSHCredentials(ctx context.Context, formats strfmt.Registry) error {

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
func (m *V1VerifiedRepoConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VerifiedRepoConfig) UnmarshalBinary(b []byte) error {
	var res V1VerifiedRepoConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}