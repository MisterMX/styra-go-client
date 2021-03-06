// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SystemsV1SystemsPutBundleDeployResponse systems v1 systems put bundle deploy response
//
// swagger:model systems.v1.SystemsPutBundleDeployResponse
type SystemsV1SystemsPutBundleDeployResponse struct {

	// request id
	RequestID string `json:"request_id,omitempty"`
}

// Validate validates this systems v1 systems put bundle deploy response
func (m *SystemsV1SystemsPutBundleDeployResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this systems v1 systems put bundle deploy response based on context it is used
func (m *SystemsV1SystemsPutBundleDeployResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SystemsV1SystemsPutBundleDeployResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemsV1SystemsPutBundleDeployResponse) UnmarshalBinary(b []byte) error {
	var res SystemsV1SystemsPutBundleDeployResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
