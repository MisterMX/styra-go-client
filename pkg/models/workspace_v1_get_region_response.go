// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WorkspaceV1GetRegionResponse workspace v1 get region response
//
// swagger:model workspace.v1.GetRegionResponse
type WorkspaceV1GetRegionResponse struct {

	// result
	// Required: true
	Result []*WorkspaceV1RegionFF `json:"result"`
}

// Validate validates this workspace v1 get region response
func (m *WorkspaceV1GetRegionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkspaceV1GetRegionResponse) validateResult(formats strfmt.Registry) error {

	if err := validate.Required("result", "body", m.Result); err != nil {
		return err
	}

	for i := 0; i < len(m.Result); i++ {
		if swag.IsZero(m.Result[i]) { // not required
			continue
		}

		if m.Result[i] != nil {
			if err := m.Result[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("result" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("result" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this workspace v1 get region response based on the context it is used
func (m *WorkspaceV1GetRegionResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkspaceV1GetRegionResponse) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Result); i++ {

		if m.Result[i] != nil {
			if err := m.Result[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("result" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("result" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkspaceV1GetRegionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkspaceV1GetRegionResponse) UnmarshalBinary(b []byte) error {
	var res WorkspaceV1GetRegionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}