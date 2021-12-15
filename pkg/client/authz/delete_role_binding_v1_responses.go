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

// DeleteRoleBindingV1Reader is a Reader for the DeleteRoleBindingV1 structure.
type DeleteRoleBindingV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRoleBindingV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRoleBindingV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteRoleBindingV1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRoleBindingV1OK creates a DeleteRoleBindingV1OK with default headers values
func NewDeleteRoleBindingV1OK() *DeleteRoleBindingV1OK {
	return &DeleteRoleBindingV1OK{}
}

/* DeleteRoleBindingV1OK describes a response with status code 200, with default header values.

OK
*/
type DeleteRoleBindingV1OK struct {
	Payload *models.AuthzV1RoleBindingsDeleteResponse
}

func (o *DeleteRoleBindingV1OK) Error() string {
	return fmt.Sprintf("[DELETE /v1/authz/rolebindings/{resourcetype}/{resource}/{rolebinding}][%d] deleteRoleBindingV1OK  %+v", 200, o.Payload)
}
func (o *DeleteRoleBindingV1OK) GetPayload() *models.AuthzV1RoleBindingsDeleteResponse {
	return o.Payload
}

func (o *DeleteRoleBindingV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthzV1RoleBindingsDeleteResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoleBindingV1NotFound creates a DeleteRoleBindingV1NotFound with default headers values
func NewDeleteRoleBindingV1NotFound() *DeleteRoleBindingV1NotFound {
	return &DeleteRoleBindingV1NotFound{}
}

/* DeleteRoleBindingV1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteRoleBindingV1NotFound struct {
	Payload *models.MetaV1ErrorResponse
}

func (o *DeleteRoleBindingV1NotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/authz/rolebindings/{resourcetype}/{resource}/{rolebinding}][%d] deleteRoleBindingV1NotFound  %+v", 404, o.Payload)
}
func (o *DeleteRoleBindingV1NotFound) GetPayload() *models.MetaV1ErrorResponse {
	return o.Payload
}

func (o *DeleteRoleBindingV1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetaV1ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}