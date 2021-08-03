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

// V2ComposedRoleBindings v2 composed role bindings
//
// swagger:model v2.ComposedRoleBindings
type V2ComposedRoleBindings struct {

	// kind
	// Required: true
	Kind *string `json:"kind"`

	// policy
	Policy V2ComposedRoleBindingsPolicy `json:"policy,omitempty"`

	// structured
	Structured []*V2RoleBinding `json:"structured"`
}

// Validate validates this v2 composed role bindings
func (m *V2ComposedRoleBindings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStructured(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V2ComposedRoleBindings) validateKind(formats strfmt.Registry) error {

	if err := validate.Required("kind", "body", m.Kind); err != nil {
		return err
	}

	return nil
}

func (m *V2ComposedRoleBindings) validateStructured(formats strfmt.Registry) error {
	if swag.IsZero(m.Structured) { // not required
		return nil
	}

	for i := 0; i < len(m.Structured); i++ {
		if swag.IsZero(m.Structured[i]) { // not required
			continue
		}

		if m.Structured[i] != nil {
			if err := m.Structured[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("structured" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v2 composed role bindings based on the context it is used
func (m *V2ComposedRoleBindings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStructured(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V2ComposedRoleBindings) contextValidateStructured(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Structured); i++ {

		if m.Structured[i] != nil {
			if err := m.Structured[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("structured" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V2ComposedRoleBindings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2ComposedRoleBindings) UnmarshalBinary(b []byte) error {
	var res V2ComposedRoleBindings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
