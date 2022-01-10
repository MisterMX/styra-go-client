// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DatasourcesV1Common datasources v1 common
//
// swagger:model datasources.v1.Common
type DatasourcesV1Common struct {

	// category
	// Required: true
	Category *string `json:"category"`

	// description
	Description string `json:"description,omitempty"`

	// enabled
	Enabled *bool `json:"enabled,omitempty"`

	// on premises
	// Required: true
	OnPremises *bool `json:"on_premises"`

	// type
	// Required: true
	// Enum: [pull push]
	Type *string `json:"type"`
}

// Validate validates this datasources v1 common
func (m *DatasourcesV1Common) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOnPremises(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatasourcesV1Common) validateCategory(formats strfmt.Registry) error {

	if err := validate.Required("category", "body", m.Category); err != nil {
		return err
	}

	return nil
}

func (m *DatasourcesV1Common) validateOnPremises(formats strfmt.Registry) error {

	if err := validate.Required("on_premises", "body", m.OnPremises); err != nil {
		return err
	}

	return nil
}

var datasourcesV1CommonTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["pull","push"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		datasourcesV1CommonTypeTypePropEnum = append(datasourcesV1CommonTypeTypePropEnum, v)
	}
}

const (

	// DatasourcesV1CommonTypePull captures enum value "pull"
	DatasourcesV1CommonTypePull string = "pull"

	// DatasourcesV1CommonTypePush captures enum value "push"
	DatasourcesV1CommonTypePush string = "push"
)

// prop value enum
func (m *DatasourcesV1Common) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, datasourcesV1CommonTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DatasourcesV1Common) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this datasources v1 common based on context it is used
func (m *DatasourcesV1Common) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DatasourcesV1Common) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatasourcesV1Common) UnmarshalBinary(b []byte) error {
	var res DatasourcesV1Common
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
