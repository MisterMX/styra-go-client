// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AuthzV2RoleBindingsPutStatusResponse authz v2 role bindings put status response
//
// swagger:model authz.v2.RoleBindingsPutStatusResponse
type AuthzV2RoleBindingsPutStatusResponse struct {

	// request id
	RequestID string `json:"request_id,omitempty"`
}

// Validate validates this authz v2 role bindings put status response
func (m *AuthzV2RoleBindingsPutStatusResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this authz v2 role bindings put status response based on context it is used
func (m *AuthzV2RoleBindingsPutStatusResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AuthzV2RoleBindingsPutStatusResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthzV2RoleBindingsPutStatusResponse) UnmarshalBinary(b []byte) error {
	var res AuthzV2RoleBindingsPutStatusResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
