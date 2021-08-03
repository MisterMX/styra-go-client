// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V2RoleBindingsPutResponse v2 role bindings put response
//
// swagger:model v2.RoleBindingsPutResponse
type V2RoleBindingsPutResponse struct {

	// request id
	RequestID string `json:"request_id,omitempty"`
}

// Validate validates this v2 role bindings put response
func (m *V2RoleBindingsPutResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v2 role bindings put response based on context it is used
func (m *V2RoleBindingsPutResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V2RoleBindingsPutResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2RoleBindingsPutResponse) UnmarshalBinary(b []byte) error {
	var res V2RoleBindingsPutResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}