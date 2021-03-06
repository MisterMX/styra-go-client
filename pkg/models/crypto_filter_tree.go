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

// CryptoFilterTree crypto filter tree
//
// swagger:model crypto.FilterTree
type CryptoFilterTree struct {

	// digest
	Digest string `json:"digest,omitempty"`

	// nodes
	Nodes map[string]CryptoFilterTree `json:"nodes,omitempty"`
}

// Validate validates this crypto filter tree
func (m *CryptoFilterTree) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNodes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CryptoFilterTree) validateNodes(formats strfmt.Registry) error {
	if swag.IsZero(m.Nodes) { // not required
		return nil
	}

	for k := range m.Nodes {

		if err := validate.Required("nodes"+"."+k, "body", m.Nodes[k]); err != nil {
			return err
		}
		if val, ok := m.Nodes[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this crypto filter tree based on the context it is used
func (m *CryptoFilterTree) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNodes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CryptoFilterTree) contextValidateNodes(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Nodes {

		if val, ok := m.Nodes[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CryptoFilterTree) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CryptoFilterTree) UnmarshalBinary(b []byte) error {
	var res CryptoFilterTree
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
