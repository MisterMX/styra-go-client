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

// PoliciesV1PoliciesPutRequest policies v1 policies put request
//
// swagger:model policies.v1.PoliciesPutRequest
type PoliciesV1PoliciesPutRequest struct {

	// modules
	// Required: true
	Modules map[string]string `json:"modules"`

	// signature
	// Required: true
	Signature *CryptoSignature `json:"signature"`
}

// Validate validates this policies v1 policies put request
func (m *PoliciesV1PoliciesPutRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateModules(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignature(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PoliciesV1PoliciesPutRequest) validateModules(formats strfmt.Registry) error {

	if err := validate.Required("modules", "body", m.Modules); err != nil {
		return err
	}

	return nil
}

func (m *PoliciesV1PoliciesPutRequest) validateSignature(formats strfmt.Registry) error {

	if err := validate.Required("signature", "body", m.Signature); err != nil {
		return err
	}

	if m.Signature != nil {
		if err := m.Signature.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("signature")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("signature")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this policies v1 policies put request based on the context it is used
func (m *PoliciesV1PoliciesPutRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSignature(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PoliciesV1PoliciesPutRequest) contextValidateSignature(ctx context.Context, formats strfmt.Registry) error {

	if m.Signature != nil {
		if err := m.Signature.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("signature")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("signature")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PoliciesV1PoliciesPutRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PoliciesV1PoliciesPutRequest) UnmarshalBinary(b []byte) error {
	var res PoliciesV1PoliciesPutRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}