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

// V1DeploymentInstruction v1 deployment instruction
//
// swagger:model v1.DeploymentInstruction
type V1DeploymentInstruction struct {

	// category/tool name
	// Required: true
	// Read Only: true
	Category string `json:"category"`

	// commands
	// Required: true
	// Read Only: true
	Commands []*V1DeploymentCommand `json:"commands"`
}

// Validate validates this v1 deployment instruction
func (m *V1DeploymentInstruction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommands(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1DeploymentInstruction) validateCategory(formats strfmt.Registry) error {

	if err := validate.RequiredString("category", "body", m.Category); err != nil {
		return err
	}

	return nil
}

func (m *V1DeploymentInstruction) validateCommands(formats strfmt.Registry) error {

	if err := validate.Required("commands", "body", m.Commands); err != nil {
		return err
	}

	for i := 0; i < len(m.Commands); i++ {
		if swag.IsZero(m.Commands[i]) { // not required
			continue
		}

		if m.Commands[i] != nil {
			if err := m.Commands[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("commands" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 deployment instruction based on the context it is used
func (m *V1DeploymentInstruction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCategory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCommands(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1DeploymentInstruction) contextValidateCategory(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "category", "body", string(m.Category)); err != nil {
		return err
	}

	return nil
}

func (m *V1DeploymentInstruction) contextValidateCommands(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "commands", "body", []*V1DeploymentCommand(m.Commands)); err != nil {
		return err
	}

	for i := 0; i < len(m.Commands); i++ {

		if m.Commands[i] != nil {
			if err := m.Commands[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("commands" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1DeploymentInstruction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1DeploymentInstruction) UnmarshalBinary(b []byte) error {
	var res V1DeploymentInstruction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}