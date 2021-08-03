// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1StacksDeleteResponse v1 stacks delete response
//
// swagger:model v1.StacksDeleteResponse
type V1StacksDeleteResponse struct {

	// request id
	RequestID string `json:"request_id,omitempty"`
}

// Validate validates this v1 stacks delete response
func (m *V1StacksDeleteResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 stacks delete response based on context it is used
func (m *V1StacksDeleteResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1StacksDeleteResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1StacksDeleteResponse) UnmarshalBinary(b []byte) error {
	var res V1StacksDeleteResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
