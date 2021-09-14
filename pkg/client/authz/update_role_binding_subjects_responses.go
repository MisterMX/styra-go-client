// Code generated by go-swagger; DO NOT EDIT.

package authz

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mistermx/styra-go-client/pkg/models"
)

// UpdateRoleBindingSubjectsReader is a Reader for the UpdateRoleBindingSubjects structure.
type UpdateRoleBindingSubjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRoleBindingSubjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRoleBindingSubjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateRoleBindingSubjectsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateRoleBindingSubjectsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRoleBindingSubjectsOK creates a UpdateRoleBindingSubjectsOK with default headers values
func NewUpdateRoleBindingSubjectsOK() *UpdateRoleBindingSubjectsOK {
	return &UpdateRoleBindingSubjectsOK{}
}

/* UpdateRoleBindingSubjectsOK describes a response with status code 200, with default header values.

OK
*/
type UpdateRoleBindingSubjectsOK struct {
	Payload *models.V2RoleBindingsPostSubjectsResponse
}

func (o *UpdateRoleBindingSubjectsOK) Error() string {
	return fmt.Sprintf("[POST /v2/authz/rolebindings/{id}/subjects][%d] updateRoleBindingSubjectsOK  %+v", 200, o.Payload)
}
func (o *UpdateRoleBindingSubjectsOK) GetPayload() *models.V2RoleBindingsPostSubjectsResponse {
	return o.Payload
}

func (o *UpdateRoleBindingSubjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V2RoleBindingsPostSubjectsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRoleBindingSubjectsBadRequest creates a UpdateRoleBindingSubjectsBadRequest with default headers values
func NewUpdateRoleBindingSubjectsBadRequest() *UpdateRoleBindingSubjectsBadRequest {
	return &UpdateRoleBindingSubjectsBadRequest{}
}

/* UpdateRoleBindingSubjectsBadRequest describes a response with status code 400, with default header values.

Invalid Parameter
*/
type UpdateRoleBindingSubjectsBadRequest struct {
	Payload *models.V2RoleBindingsPostResponse
}

func (o *UpdateRoleBindingSubjectsBadRequest) Error() string {
	return fmt.Sprintf("[POST /v2/authz/rolebindings/{id}/subjects][%d] updateRoleBindingSubjectsBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateRoleBindingSubjectsBadRequest) GetPayload() *models.V2RoleBindingsPostResponse {
	return o.Payload
}

func (o *UpdateRoleBindingSubjectsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V2RoleBindingsPostResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRoleBindingSubjectsNotFound creates a UpdateRoleBindingSubjectsNotFound with default headers values
func NewUpdateRoleBindingSubjectsNotFound() *UpdateRoleBindingSubjectsNotFound {
	return &UpdateRoleBindingSubjectsNotFound{}
}

/* UpdateRoleBindingSubjectsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateRoleBindingSubjectsNotFound struct {
	Payload *models.V1ErrorResponse
}

func (o *UpdateRoleBindingSubjectsNotFound) Error() string {
	return fmt.Sprintf("[POST /v2/authz/rolebindings/{id}/subjects][%d] updateRoleBindingSubjectsNotFound  %+v", 404, o.Payload)
}
func (o *UpdateRoleBindingSubjectsNotFound) GetPayload() *models.V1ErrorResponse {
	return o.Payload
}

func (o *UpdateRoleBindingSubjectsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
