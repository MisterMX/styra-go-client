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

// DatasourcesV1AWSCommon datasources v1 a w s common
//
// swagger:model datasources.v1.AWS.Common
type DatasourcesV1AWSCommon struct {

	// Secret ID with AWS credentials
	// Required: true
	Credentials *string `json:"credentials"`

	// AWS region
	// Required: true
	Region *string `json:"region"`
}

// Validate validates this datasources v1 a w s common
func (m *DatasourcesV1AWSCommon) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatasourcesV1AWSCommon) validateCredentials(formats strfmt.Registry) error {

	if err := validate.Required("credentials", "body", m.Credentials); err != nil {
		return err
	}

	return nil
}

func (m *DatasourcesV1AWSCommon) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("region", "body", m.Region); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this datasources v1 a w s common based on context it is used
func (m *DatasourcesV1AWSCommon) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DatasourcesV1AWSCommon) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatasourcesV1AWSCommon) UnmarshalBinary(b []byte) error {
	var res DatasourcesV1AWSCommon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
