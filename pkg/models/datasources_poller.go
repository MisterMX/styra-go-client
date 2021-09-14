// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DatasourcesPoller datasources poller
//
// swagger:model datasources.Poller
type DatasourcesPoller struct {

	// polling interval
	PollingInterval *string `json:"polling_interval,omitempty"`
}

// Validate validates this datasources poller
func (m *DatasourcesPoller) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this datasources poller based on context it is used
func (m *DatasourcesPoller) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DatasourcesPoller) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatasourcesPoller) UnmarshalBinary(b []byte) error {
	var res DatasourcesPoller
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}