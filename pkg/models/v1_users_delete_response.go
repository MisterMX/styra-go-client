// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1UsersDeleteResponse v1 users delete response
//
// swagger:model v1.UsersDeleteResponse
type V1UsersDeleteResponse struct {

	// request id
	RequestID string `json:"request_id,omitempty"`
}

// Validate validates this v1 users delete response
func (m *V1UsersDeleteResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 users delete response based on context it is used
func (m *V1UsersDeleteResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1UsersDeleteResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1UsersDeleteResponse) UnmarshalBinary(b []byte) error {
	var res V1UsersDeleteResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
