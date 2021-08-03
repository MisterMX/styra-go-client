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

// V2ResourceFilter v2 resource filter
//
// swagger:model v2.ResourceFilter
type V2ResourceFilter struct {

	// systems
	// Required: true
	Systems *V2Filter `json:"systems"`

	// workspaces
	// Required: true
	Workspaces *V2Filter `json:"workspaces"`
}

// Validate validates this v2 resource filter
func (m *V2ResourceFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSystems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkspaces(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V2ResourceFilter) validateSystems(formats strfmt.Registry) error {

	if err := validate.Required("systems", "body", m.Systems); err != nil {
		return err
	}

	if m.Systems != nil {
		if err := m.Systems.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("systems")
			}
			return err
		}
	}

	return nil
}

func (m *V2ResourceFilter) validateWorkspaces(formats strfmt.Registry) error {

	if err := validate.Required("workspaces", "body", m.Workspaces); err != nil {
		return err
	}

	if m.Workspaces != nil {
		if err := m.Workspaces.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workspaces")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v2 resource filter based on the context it is used
func (m *V2ResourceFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSystems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWorkspaces(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V2ResourceFilter) contextValidateSystems(ctx context.Context, formats strfmt.Registry) error {

	if m.Systems != nil {
		if err := m.Systems.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("systems")
			}
			return err
		}
	}

	return nil
}

func (m *V2ResourceFilter) contextValidateWorkspaces(ctx context.Context, formats strfmt.Registry) error {

	if m.Workspaces != nil {
		if err := m.Workspaces.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workspaces")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V2ResourceFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2ResourceFilter) UnmarshalBinary(b []byte) error {
	var res V2ResourceFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}