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

// AuthzV2RoleBindingsPutSubjectsRequest authz v2 role bindings put subjects request
//
// swagger:model authz.v2.RoleBindingsPutSubjectsRequest
type AuthzV2RoleBindingsPutSubjectsRequest struct {

	// subjects
	// Required: true
	Subjects []*AuthzV2Subject `json:"subjects"`
}

// Validate validates this authz v2 role bindings put subjects request
func (m *AuthzV2RoleBindingsPutSubjectsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSubjects(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthzV2RoleBindingsPutSubjectsRequest) validateSubjects(formats strfmt.Registry) error {

	if err := validate.Required("subjects", "body", m.Subjects); err != nil {
		return err
	}

	for i := 0; i < len(m.Subjects); i++ {
		if swag.IsZero(m.Subjects[i]) { // not required
			continue
		}

		if m.Subjects[i] != nil {
			if err := m.Subjects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subjects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("subjects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this authz v2 role bindings put subjects request based on the context it is used
func (m *AuthzV2RoleBindingsPutSubjectsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSubjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthzV2RoleBindingsPutSubjectsRequest) contextValidateSubjects(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Subjects); i++ {

		if m.Subjects[i] != nil {
			if err := m.Subjects[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subjects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("subjects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuthzV2RoleBindingsPutSubjectsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthzV2RoleBindingsPutSubjectsRequest) UnmarshalBinary(b []byte) error {
	var res AuthzV2RoleBindingsPutSubjectsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}