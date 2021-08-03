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

// V1SourceControlConfig v1 source control config
//
// swagger:model v1.SourceControlConfig
type V1SourceControlConfig struct {

	// origin
	// Required: true
	Origin *V1GitRepoConfig `json:"origin"`
}

// Validate validates this v1 source control config
func (m *V1SourceControlConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrigin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SourceControlConfig) validateOrigin(formats strfmt.Registry) error {

	if err := validate.Required("origin", "body", m.Origin); err != nil {
		return err
	}

	if m.Origin != nil {
		if err := m.Origin.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("origin")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 source control config based on the context it is used
func (m *V1SourceControlConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOrigin(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SourceControlConfig) contextValidateOrigin(ctx context.Context, formats strfmt.Registry) error {

	if m.Origin != nil {
		if err := m.Origin.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("origin")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SourceControlConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SourceControlConfig) UnmarshalBinary(b []byte) error {
	var res V1SourceControlConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
