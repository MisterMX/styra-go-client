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

// GitV1CommitInput git v1 commit input
//
// swagger:model git.v1.CommitInput
type GitV1CommitInput struct {

	// author
	// Required: true
	Author *string `json:"author"`

	// email
	// Required: true
	Email *string `json:"email"`

	// Map of filenames to file contents
	// Required: true
	Files map[string]string `json:"files"`

	// List of filenames to delete from the repo
	// Required: true
	FilesToDelete []string `json:"files_to_delete"`

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this git v1 commit input
func (m *GitV1CommitInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilesToDelete(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitV1CommitInput) validateAuthor(formats strfmt.Registry) error {

	if err := validate.Required("author", "body", m.Author); err != nil {
		return err
	}

	return nil
}

func (m *GitV1CommitInput) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

func (m *GitV1CommitInput) validateFiles(formats strfmt.Registry) error {

	if err := validate.Required("files", "body", m.Files); err != nil {
		return err
	}

	return nil
}

func (m *GitV1CommitInput) validateFilesToDelete(formats strfmt.Registry) error {

	if err := validate.Required("files_to_delete", "body", m.FilesToDelete); err != nil {
		return err
	}

	return nil
}

func (m *GitV1CommitInput) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this git v1 commit input based on context it is used
func (m *GitV1CommitInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GitV1CommitInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitV1CommitInput) UnmarshalBinary(b []byte) error {
	var res GitV1CommitInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
