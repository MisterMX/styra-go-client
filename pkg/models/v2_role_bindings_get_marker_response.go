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

// V2RoleBindingsGetMarkerResponse v2 role bindings get marker response
//
// swagger:model v2.RoleBindingsGetMarkerResponse
type V2RoleBindingsGetMarkerResponse struct {

	// migrated
	// Required: true
	Migrated *bool `json:"migrated"`

	// request id
	RequestID string `json:"request_id,omitempty"`
}

// Validate validates this v2 role bindings get marker response
func (m *V2RoleBindingsGetMarkerResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMigrated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V2RoleBindingsGetMarkerResponse) validateMigrated(formats strfmt.Registry) error {

	if err := validate.Required("migrated", "body", m.Migrated); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v2 role bindings get marker response based on context it is used
func (m *V2RoleBindingsGetMarkerResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V2RoleBindingsGetMarkerResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2RoleBindingsGetMarkerResponse) UnmarshalBinary(b []byte) error {
	var res V2RoleBindingsGetMarkerResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}