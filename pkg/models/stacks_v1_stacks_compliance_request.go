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
)

// StacksV1StacksComplianceRequest stacks v1 stacks compliance request
//
// swagger:model stacks.v1.StacksComplianceRequest
type StacksV1StacksComplianceRequest struct {

	// draft policies to be used for 'new' violations computation (path => rego)
	Drafts map[string]string `json:"drafts,omitempty"`

	// run extended compliance validation that is specific for the system/stack type
	Extended bool `json:"extended,omitempty"`

	// filter violations with this selector (dot.path => value)
	Filter interface{} `json:"filter,omitempty"`

	// group results by dot.path values (list of group levels with list of fields at each level)
	GroupBy []StacksV1StacksComplianceRequestGroupBy `json:"group_by"`

	// mock http.send and opa.runtime builtins
	Mocks *SystemsV1BuiltinMocks `json:"mocks,omitempty"`

	// validation mode. One of (delta, all, delta-count, all-count)
	Mode *string `json:"mode,omitempty"`

	// policy type to narrow the monitor policy search (e.g. validating, mutating). Default (empty string or missing) is to run all monitoring policies
	PolicyType string `json:"policy_type,omitempty"`

	// list of fields to sort by
	Sort []*SystemsV1SortField `json:"sort"`
}

// Validate validates this stacks v1 stacks compliance request
func (m *StacksV1StacksComplianceRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroupBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMocks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StacksV1StacksComplianceRequest) validateGroupBy(formats strfmt.Registry) error {
	if swag.IsZero(m.GroupBy) { // not required
		return nil
	}

	for i := 0; i < len(m.GroupBy); i++ {

		if err := m.GroupBy[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("group_by" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("group_by" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *StacksV1StacksComplianceRequest) validateMocks(formats strfmt.Registry) error {
	if swag.IsZero(m.Mocks) { // not required
		return nil
	}

	if m.Mocks != nil {
		if err := m.Mocks.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mocks")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mocks")
			}
			return err
		}
	}

	return nil
}

func (m *StacksV1StacksComplianceRequest) validateSort(formats strfmt.Registry) error {
	if swag.IsZero(m.Sort) { // not required
		return nil
	}

	for i := 0; i < len(m.Sort); i++ {
		if swag.IsZero(m.Sort[i]) { // not required
			continue
		}

		if m.Sort[i] != nil {
			if err := m.Sort[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sort" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sort" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this stacks v1 stacks compliance request based on the context it is used
func (m *StacksV1StacksComplianceRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGroupBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMocks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSort(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StacksV1StacksComplianceRequest) contextValidateGroupBy(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.GroupBy); i++ {

		if err := m.GroupBy[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("group_by" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("group_by" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *StacksV1StacksComplianceRequest) contextValidateMocks(ctx context.Context, formats strfmt.Registry) error {

	if m.Mocks != nil {
		if err := m.Mocks.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mocks")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mocks")
			}
			return err
		}
	}

	return nil
}

func (m *StacksV1StacksComplianceRequest) contextValidateSort(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Sort); i++ {

		if m.Sort[i] != nil {
			if err := m.Sort[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sort" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sort" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StacksV1StacksComplianceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StacksV1StacksComplianceRequest) UnmarshalBinary(b []byte) error {
	var res StacksV1StacksComplianceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
