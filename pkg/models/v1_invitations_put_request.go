// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1InvitationsPutRequest v1 invitations put request
//
// swagger:model v1.InvitationsPutRequest
type V1InvitationsPutRequest struct {

	// new user password
	// Required: true
	Password *string `json:"password"`

	// terms of service were accepted
	// Required: true
	TosChecked *bool `json:"tos_checked"`

	// new user ID
	// Required: true
	UserID *string `json:"user_id"`
}

// Validate validates this v1 invitations put request
func (m *V1InvitationsPutRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTosChecked(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1InvitationsPutRequest) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

func (m *V1InvitationsPutRequest) validateTosChecked(formats strfmt.Registry) error {

	if err := validate.Required("tos_checked", "body", m.TosChecked); err != nil {
		return err
	}

	return nil
}

func (m *V1InvitationsPutRequest) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("user_id", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 invitations put request based on context it is used
func (m *V1InvitationsPutRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1InvitationsPutRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1InvitationsPutRequest) UnmarshalBinary(b []byte) error {
	var res V1InvitationsPutRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
