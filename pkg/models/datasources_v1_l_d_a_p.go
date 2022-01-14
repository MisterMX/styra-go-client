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

// DatasourcesV1LDAP ldap
// Example: {"category":"ldap","credentials":"ldapcreds","on_premises":false,"polling_interval":"60s","search":{"base_DN":"dc=test,dc=styra,dc=com","deref":"never","filter":"(objectclass=*)","scope":"whole-subtree"},"type":"pull","url":["ldaps://example.com:33636","ldap://example.com:33389"]}
//
// swagger:model datasources.v1.LDAP
type DatasourcesV1LDAP struct {
	DatasourcesV1Common

	DatasourcesV1Poller

	DatasourcesV1RegoFiltering

	DatasourcesV1TLSSettings

	DatasourcesV1RateLimiter

	// Secret ID with credentials
	Credentials string `json:"credentials,omitempty"`

	// search
	Search *DatasourcesV1LDAPAO5Search `json:"search,omitempty"`

	// List of URLs: main + replicas
	// Required: true
	Urls []string `json:"urls"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *DatasourcesV1LDAP) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 DatasourcesV1Common
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.DatasourcesV1Common = aO0

	// AO1
	var aO1 DatasourcesV1Poller
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.DatasourcesV1Poller = aO1

	// AO2
	var aO2 DatasourcesV1RegoFiltering
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.DatasourcesV1RegoFiltering = aO2

	// AO3
	var aO3 DatasourcesV1TLSSettings
	if err := swag.ReadJSON(raw, &aO3); err != nil {
		return err
	}
	m.DatasourcesV1TLSSettings = aO3

	// AO4
	var aO4 DatasourcesV1RateLimiter
	if err := swag.ReadJSON(raw, &aO4); err != nil {
		return err
	}
	m.DatasourcesV1RateLimiter = aO4

	// AO5
	var dataAO5 struct {
		Credentials string `json:"credentials,omitempty"`

		Search *DatasourcesV1LDAPAO5Search `json:"search,omitempty"`

		Urls []string `json:"urls"`
	}
	if err := swag.ReadJSON(raw, &dataAO5); err != nil {
		return err
	}

	m.Credentials = dataAO5.Credentials

	m.Search = dataAO5.Search

	m.Urls = dataAO5.Urls

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m DatasourcesV1LDAP) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 6)

	aO0, err := swag.WriteJSON(m.DatasourcesV1Common)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.DatasourcesV1Poller)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	aO2, err := swag.WriteJSON(m.DatasourcesV1RegoFiltering)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)

	aO3, err := swag.WriteJSON(m.DatasourcesV1TLSSettings)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO3)

	aO4, err := swag.WriteJSON(m.DatasourcesV1RateLimiter)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO4)
	var dataAO5 struct {
		Credentials string `json:"credentials,omitempty"`

		Search *DatasourcesV1LDAPAO5Search `json:"search,omitempty"`

		Urls []string `json:"urls"`
	}

	dataAO5.Credentials = m.Credentials

	dataAO5.Search = m.Search

	dataAO5.Urls = m.Urls

	jsonDataAO5, errAO5 := swag.WriteJSON(dataAO5)
	if errAO5 != nil {
		return nil, errAO5
	}
	_parts = append(_parts, jsonDataAO5)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this datasources v1 l d a p
func (m *DatasourcesV1LDAP) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with DatasourcesV1Common
	if err := m.DatasourcesV1Common.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with DatasourcesV1Poller
	if err := m.DatasourcesV1Poller.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with DatasourcesV1RegoFiltering
	if err := m.DatasourcesV1RegoFiltering.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with DatasourcesV1TLSSettings
	if err := m.DatasourcesV1TLSSettings.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with DatasourcesV1RateLimiter
	if err := m.DatasourcesV1RateLimiter.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSearch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUrls(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatasourcesV1LDAP) validateSearch(formats strfmt.Registry) error {

	if swag.IsZero(m.Search) { // not required
		return nil
	}

	if m.Search != nil {
		if err := m.Search.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("search")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("search")
			}
			return err
		}
	}

	return nil
}

func (m *DatasourcesV1LDAP) validateUrls(formats strfmt.Registry) error {

	if err := validate.Required("urls", "body", m.Urls); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this datasources v1 l d a p based on the context it is used
func (m *DatasourcesV1LDAP) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with DatasourcesV1Common
	if err := m.DatasourcesV1Common.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with DatasourcesV1Poller
	if err := m.DatasourcesV1Poller.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with DatasourcesV1RegoFiltering
	if err := m.DatasourcesV1RegoFiltering.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with DatasourcesV1TLSSettings
	if err := m.DatasourcesV1TLSSettings.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with DatasourcesV1RateLimiter
	if err := m.DatasourcesV1RateLimiter.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSearch(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatasourcesV1LDAP) contextValidateSearch(ctx context.Context, formats strfmt.Registry) error {

	if m.Search != nil {
		if err := m.Search.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("search")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("search")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DatasourcesV1LDAP) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatasourcesV1LDAP) UnmarshalBinary(b []byte) error {
	var res DatasourcesV1LDAP
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DatasourcesV1LDAPAO5Search Search Request.
// Documentation: https://ldapwiki.com/wiki/SearchRequest
//
//
// swagger:model DatasourcesV1LDAPAO5Search
type DatasourcesV1LDAPAO5Search struct {

	// Search attribute selection.
	// Documentation: https://ldapwiki.com/wiki/AttributeSelection
	//
	Attributes []string `json:"attributes"`

	// Search Base DN.
	// Documentation: https://ldapwiki.com/wiki/BaseDN
	//
	// Required: true
	BaseDN *string `json:"base_DN"`

	// Search dereference policy.
	// Documentation: https://ldapwiki.com/wiki/Dereference%20Policy
	//
	// Enum: [never searching finding always]
	Deref *string `json:"deref,omitempty"`

	// Search filter.
	// Documentation: https://ldapwiki.com/wiki/LDAP%20SearchFilters
	// Examples: https://ldapwiki.com/wiki/LDAP%20Query%20Examples
	//
	// Required: true
	Filter *string `json:"filter"`

	// Search page size.
	// Documentation: https://ldapwiki.com/wiki/MaxPageSize
	//
	// Maximum: 4.294967295e+09
	// Minimum: 0
	PageSize *int64 `json:"page_size,omitempty"`

	// Search scope.
	// Documentation: https://ldapwiki.com/wiki/LDAP%20Search%20Scopes
	//
	// Enum: [base-object single-level whole-subtree]
	Scope *string `json:"scope,omitempty"`

	// Search page limit.
	// Documentation: https://ldapwiki.com/wiki/SizeLimit
	//
	// Minimum: 0
	SizeLimit *int64 `json:"size_limit,omitempty"`
}

// Validate validates this datasources v1 l d a p a o5 search
func (m *DatasourcesV1LDAPAO5Search) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBaseDN(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeref(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePageSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSizeLimit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatasourcesV1LDAPAO5Search) validateBaseDN(formats strfmt.Registry) error {

	if err := validate.Required("search"+"."+"base_DN", "body", m.BaseDN); err != nil {
		return err
	}

	return nil
}

var datasourcesV1LDAPAO5SearchTypeDerefPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["never","searching","finding","always"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		datasourcesV1LDAPAO5SearchTypeDerefPropEnum = append(datasourcesV1LDAPAO5SearchTypeDerefPropEnum, v)
	}
}

const (

	// DatasourcesV1LDAPAO5SearchDerefNever captures enum value "never"
	DatasourcesV1LDAPAO5SearchDerefNever string = "never"

	// DatasourcesV1LDAPAO5SearchDerefSearching captures enum value "searching"
	DatasourcesV1LDAPAO5SearchDerefSearching string = "searching"

	// DatasourcesV1LDAPAO5SearchDerefFinding captures enum value "finding"
	DatasourcesV1LDAPAO5SearchDerefFinding string = "finding"

	// DatasourcesV1LDAPAO5SearchDerefAlways captures enum value "always"
	DatasourcesV1LDAPAO5SearchDerefAlways string = "always"
)

// prop value enum
func (m *DatasourcesV1LDAPAO5Search) validateDerefEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, datasourcesV1LDAPAO5SearchTypeDerefPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DatasourcesV1LDAPAO5Search) validateDeref(formats strfmt.Registry) error {
	if swag.IsZero(m.Deref) { // not required
		return nil
	}

	// value enum
	if err := m.validateDerefEnum("search"+"."+"deref", "body", *m.Deref); err != nil {
		return err
	}

	return nil
}

func (m *DatasourcesV1LDAPAO5Search) validateFilter(formats strfmt.Registry) error {

	if err := validate.Required("search"+"."+"filter", "body", m.Filter); err != nil {
		return err
	}

	return nil
}

func (m *DatasourcesV1LDAPAO5Search) validatePageSize(formats strfmt.Registry) error {
	if swag.IsZero(m.PageSize) { // not required
		return nil
	}

	if err := validate.MinimumInt("search"+"."+"page_size", "body", *m.PageSize, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("search"+"."+"page_size", "body", *m.PageSize, 4.294967295e+09, false); err != nil {
		return err
	}

	return nil
}

var datasourcesV1LDAPAO5SearchTypeScopePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["base-object","single-level","whole-subtree"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		datasourcesV1LDAPAO5SearchTypeScopePropEnum = append(datasourcesV1LDAPAO5SearchTypeScopePropEnum, v)
	}
}

const (

	// DatasourcesV1LDAPAO5SearchScopeBaseDashObject captures enum value "base-object"
	DatasourcesV1LDAPAO5SearchScopeBaseDashObject string = "base-object"

	// DatasourcesV1LDAPAO5SearchScopeSingleDashLevel captures enum value "single-level"
	DatasourcesV1LDAPAO5SearchScopeSingleDashLevel string = "single-level"

	// DatasourcesV1LDAPAO5SearchScopeWholeDashSubtree captures enum value "whole-subtree"
	DatasourcesV1LDAPAO5SearchScopeWholeDashSubtree string = "whole-subtree"
)

// prop value enum
func (m *DatasourcesV1LDAPAO5Search) validateScopeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, datasourcesV1LDAPAO5SearchTypeScopePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DatasourcesV1LDAPAO5Search) validateScope(formats strfmt.Registry) error {
	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	// value enum
	if err := m.validateScopeEnum("search"+"."+"scope", "body", *m.Scope); err != nil {
		return err
	}

	return nil
}

func (m *DatasourcesV1LDAPAO5Search) validateSizeLimit(formats strfmt.Registry) error {
	if swag.IsZero(m.SizeLimit) { // not required
		return nil
	}

	if err := validate.MinimumInt("search"+"."+"size_limit", "body", *m.SizeLimit, 0, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this datasources v1 l d a p a o5 search based on context it is used
func (m *DatasourcesV1LDAPAO5Search) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DatasourcesV1LDAPAO5Search) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatasourcesV1LDAPAO5Search) UnmarshalBinary(b []byte) error {
	var res DatasourcesV1LDAPAO5Search
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}