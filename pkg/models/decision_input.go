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

// DecisionInput decision input
//
// swagger:model decision.Input
type DecisionInput struct {

	// body
	Body interface{} `json:"body,omitempty"`

	// method
	// Required: true
	Method *string `json:"method"`

	// path
	// Required: true
	Path *string `json:"path"`

	// user
	// Required: true
	User *string `json:"user"`

	// user claims
	UserClaims interface{} `json:"user_claims,omitempty"`
}

// Validate validates this decision input
func (m *DecisionInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DecisionInput) validateMethod(formats strfmt.Registry) error {

	if err := validate.Required("method", "body", m.Method); err != nil {
		return err
	}

	return nil
}

func (m *DecisionInput) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

func (m *DecisionInput) validateUser(formats strfmt.Registry) error {

	if err := validate.Required("user", "body", m.User); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this decision input based on context it is used
func (m *DecisionInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DecisionInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DecisionInput) UnmarshalBinary(b []byte) error {
	var res DecisionInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
