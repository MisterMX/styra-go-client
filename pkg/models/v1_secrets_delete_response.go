// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1SecretsDeleteResponse v1 secrets delete response
//
// swagger:model v1.SecretsDeleteResponse
type V1SecretsDeleteResponse struct {

	// request id
	RequestID string `json:"request_id,omitempty"`
}

// Validate validates this v1 secrets delete response
func (m *V1SecretsDeleteResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 secrets delete response based on context it is used
func (m *V1SecretsDeleteResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1SecretsDeleteResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SecretsDeleteResponse) UnmarshalBinary(b []byte) error {
	var res V1SecretsDeleteResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
