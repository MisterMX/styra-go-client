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

// AuthzV1RoleBindingsListAllResponse authz v1 role bindings list all response
//
// swagger:model authz.v1.RoleBindingsListAllResponse
type AuthzV1RoleBindingsListAllResponse struct {

	// request id
	RequestID string `json:"request_id,omitempty"`

	// result
	// Required: true
	Result map[string]AuthzV1RoleBindingsListAllResponseResult `json:"result"`
}

// Validate validates this authz v1 role bindings list all response
func (m *AuthzV1RoleBindingsListAllResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthzV1RoleBindingsListAllResponse) validateResult(formats strfmt.Registry) error {

	if err := validate.Required("result", "body", m.Result); err != nil {
		return err
	}

	for k := range m.Result {

		if err := validate.Required("result"+"."+k, "body", m.Result[k]); err != nil {
			return err
		}

		if err := validate.Required("result"+"."+k, "body", AuthzV1RoleBindingsListAllResponseResult(m.Result[k])); err != nil {
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

// ContextValidate validate this authz v1 role bindings list all response based on the context it is used
func (m *AuthzV1RoleBindingsListAllResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthzV1RoleBindingsListAllResponse) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.Required("result", "body", m.Result); err != nil {
		return err
	}

	for k := range m.Result {

		if err := validate.Required("result"+"."+k, "body", AuthzV1RoleBindingsListAllResponseResult(m.Result[k])); err != nil {
			return err
		}

		if val, ok := m.Result[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuthzV1RoleBindingsListAllResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthzV1RoleBindingsListAllResponse) UnmarshalBinary(b []byte) error {
	var res AuthzV1RoleBindingsListAllResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}