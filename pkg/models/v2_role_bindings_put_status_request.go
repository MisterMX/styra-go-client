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

// V2RoleBindingsPutStatusRequest v2 role bindings put status request
//
// swagger:model v2.RoleBindingsPutStatusRequest
type V2RoleBindingsPutStatusRequest struct {

	// migrated
	// Required: true
	Migrated *bool `json:"migrated"`
}

// Validate validates this v2 role bindings put status request
func (m *V2RoleBindingsPutStatusRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMigrated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V2RoleBindingsPutStatusRequest) validateMigrated(formats strfmt.Registry) error {

	if err := validate.Required("migrated", "body", m.Migrated); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v2 role bindings put status request based on context it is used
func (m *V2RoleBindingsPutStatusRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V2RoleBindingsPutStatusRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2RoleBindingsPutStatusRequest) UnmarshalBinary(b []byte) error {
	var res V2RoleBindingsPutStatusRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
