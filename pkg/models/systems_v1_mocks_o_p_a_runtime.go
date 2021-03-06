// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SystemsV1MocksOPARuntime systems v1 mocks o p a runtime
//
// swagger:model systems.v1.MocksOPARuntime
type SystemsV1MocksOPARuntime struct {

	// mock json result
	// Required: true
	Result interface{} `json:"result"`
}

// Validate validates this systems v1 mocks o p a runtime
func (m *SystemsV1MocksOPARuntime) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SystemsV1MocksOPARuntime) validateResult(formats strfmt.Registry) error {

	if m.Result == nil {
		return errors.Required("result", "body", nil)
	}

	return nil
}

// ContextValidate validates this systems v1 mocks o p a runtime based on context it is used
func (m *SystemsV1MocksOPARuntime) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SystemsV1MocksOPARuntime) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemsV1MocksOPARuntime) UnmarshalBinary(b []byte) error {
	var res SystemsV1MocksOPARuntime
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
