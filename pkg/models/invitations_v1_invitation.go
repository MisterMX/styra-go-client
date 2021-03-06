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

// InvitationsV1Invitation invitations v1 invitation
//
// swagger:model invitations.v1.Invitation
type InvitationsV1Invitation struct {

	// invitation expiration date/time
	// Required: true
	// Format: date-time
	Expiration *strfmt.DateTime `json:"expiration"`

	// invitation ID (user ID)
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this invitations v1 invitation
func (m *InvitationsV1Invitation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpiration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InvitationsV1Invitation) validateExpiration(formats strfmt.Registry) error {

	if err := validate.Required("expiration", "body", m.Expiration); err != nil {
		return err
	}

	if err := validate.FormatOf("expiration", "body", "date-time", m.Expiration.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InvitationsV1Invitation) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this invitations v1 invitation based on context it is used
func (m *InvitationsV1Invitation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InvitationsV1Invitation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InvitationsV1Invitation) UnmarshalBinary(b []byte) error {
	var res InvitationsV1Invitation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
