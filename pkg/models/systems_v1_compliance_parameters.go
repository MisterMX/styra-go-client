// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SystemsV1ComplianceParameters systems v1 compliance parameters
//
// swagger:model systems.v1.ComplianceParameters
type SystemsV1ComplianceParameters struct {

	// run extended compliance validation that is specific for the system/stack type
	Extended bool `json:"extended,omitempty"`

	// policy type to narrow the monitor policy search (e.g. validating, mutating). Default (empty string or missing) is to run all monitoring policies
	PolicyType string `json:"policy_type,omitempty"`
}

// Validate validates this systems v1 compliance parameters
func (m *SystemsV1ComplianceParameters) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this systems v1 compliance parameters based on context it is used
func (m *SystemsV1ComplianceParameters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SystemsV1ComplianceParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemsV1ComplianceParameters) UnmarshalBinary(b []byte) error {
	var res SystemsV1ComplianceParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
