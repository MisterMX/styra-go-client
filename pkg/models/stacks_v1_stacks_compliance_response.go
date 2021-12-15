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

// StacksV1StacksComplianceResponse stacks v1 stacks compliance response
//
// swagger:model stacks.v1.StacksComplianceResponse
type StacksV1StacksComplianceResponse struct {

	// Mock results
	Mocks *SystemsV1BuiltinMocksResponse `json:"mocks,omitempty"`

	// request id
	RequestID string `json:"request_id,omitempty"`

	// result
	// Required: true
	Result map[string]SystemsV1ComplianceValidation `json:"result"`
}

// Validate validates this stacks v1 stacks compliance response
func (m *StacksV1StacksComplianceResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMocks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StacksV1StacksComplianceResponse) validateMocks(formats strfmt.Registry) error {
	if swag.IsZero(m.Mocks) { // not required
		return nil
	}

	if m.Mocks != nil {
		if err := m.Mocks.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mocks")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mocks")
			}
			return err
		}
	}

	return nil
}

func (m *StacksV1StacksComplianceResponse) validateResult(formats strfmt.Registry) error {

	if err := validate.Required("result", "body", m.Result); err != nil {
		return err
	}

	for k := range m.Result {

		if err := validate.Required("result"+"."+k, "body", m.Result[k]); err != nil {
			return err
		}
		if val, ok := m.Result[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this stacks v1 stacks compliance response based on the context it is used
func (m *StacksV1StacksComplianceResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMocks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StacksV1StacksComplianceResponse) contextValidateMocks(ctx context.Context, formats strfmt.Registry) error {

	if m.Mocks != nil {
		if err := m.Mocks.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mocks")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mocks")
			}
			return err
		}
	}

	return nil
}

func (m *StacksV1StacksComplianceResponse) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.Required("result", "body", m.Result); err != nil {
		return err
	}

	for k := range m.Result {

		if val, ok := m.Result[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StacksV1StacksComplianceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StacksV1StacksComplianceResponse) UnmarshalBinary(b []byte) error {
	var res StacksV1StacksComplianceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
