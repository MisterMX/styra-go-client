// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StatusV1StatusPostResponse status v1 status post response
//
// swagger:model status.v1.StatusPostResponse
type StatusV1StatusPostResponse struct {

	// request id
	RequestID string `json:"request_id,omitempty"`
}

// Validate validates this status v1 status post response
func (m *StatusV1StatusPostResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this status v1 status post response based on context it is used
func (m *StatusV1StatusPostResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StatusV1StatusPostResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatusV1StatusPostResponse) UnmarshalBinary(b []byte) error {
	var res StatusV1StatusPostResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
