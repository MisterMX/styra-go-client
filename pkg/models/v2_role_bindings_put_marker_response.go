// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V2RoleBindingsPutMarkerResponse v2 role bindings put marker response
//
// swagger:model v2.RoleBindingsPutMarkerResponse
type V2RoleBindingsPutMarkerResponse struct {

	// request id
	RequestID string `json:"request_id,omitempty"`
}

// Validate validates this v2 role bindings put marker response
func (m *V2RoleBindingsPutMarkerResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v2 role bindings put marker response based on context it is used
func (m *V2RoleBindingsPutMarkerResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V2RoleBindingsPutMarkerResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2RoleBindingsPutMarkerResponse) UnmarshalBinary(b []byte) error {
	var res V2RoleBindingsPutMarkerResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}