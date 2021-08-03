// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1NotificationToolResult v1 notification tool result
//
// swagger:model v1.NotificationToolResult
type V1NotificationToolResult struct {

	// installed
	Installed bool `json:"installed,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this v1 notification tool result
func (m *V1NotificationToolResult) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 notification tool result based on context it is used
func (m *V1NotificationToolResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1NotificationToolResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1NotificationToolResult) UnmarshalBinary(b []byte) error {
	var res V1NotificationToolResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}