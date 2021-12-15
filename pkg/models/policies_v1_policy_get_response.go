// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PoliciesV1PolicyGetResponse policies v1 policy get response
//
// swagger:model policies.v1.PolicyGetResponse
type PoliciesV1PolicyGetResponse struct {

	// request id
	RequestID string `json:"request_id,omitempty"`

	// result
	// Required: true
	Result PoliciesV1PolicyGetResponseResult `json:"result"`
}

// Validate validates this policies v1 policy get response
func (m *PoliciesV1PolicyGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PoliciesV1PolicyGetResponse) validateResult(formats strfmt.Registry) error {

	if m.Result == nil {
		return errors.Required("result", "body", nil)
	}

	return nil
}

// ContextValidate validates this policies v1 policy get response based on context it is used
func (m *PoliciesV1PolicyGetResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PoliciesV1PolicyGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PoliciesV1PolicyGetResponse) UnmarshalBinary(b []byte) error {
	var res PoliciesV1PolicyGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}