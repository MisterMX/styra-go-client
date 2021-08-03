// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1StacksPutResponse v1 stacks put response
//
// swagger:model v1.StacksPutResponse
type V1StacksPutResponse struct {

	// request id
	RequestID string `json:"request_id,omitempty"`
}

// Validate validates this v1 stacks put response
func (m *V1StacksPutResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 stacks put response based on context it is used
func (m *V1StacksPutResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1StacksPutResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1StacksPutResponse) UnmarshalBinary(b []byte) error {
	var res V1StacksPutResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
