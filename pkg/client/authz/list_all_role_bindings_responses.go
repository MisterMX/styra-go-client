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

// ListAllRoleBindingsReader is a Reader for the ListAllRoleBindings structure.
type ListAllRoleBindingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAllRoleBindingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAllRoleBindingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListAllRoleBindingsOK creates a ListAllRoleBindingsOK with default headers values
func NewListAllRoleBindingsOK() *ListAllRoleBindingsOK {
	return &ListAllRoleBindingsOK{}
}

/* ListAllRoleBindingsOK describes a response with status code 200, with default header values.

OK
*/
type ListAllRoleBindingsOK struct {
	Payload *models.V1RoleBindingsListAllResponse
}

func (o *ListAllRoleBindingsOK) Error() string {
	return fmt.Sprintf("[GET /v1/authz/rolebindings][%d] listAllRoleBindingsOK  %+v", 200, o.Payload)
}
func (o *ListAllRoleBindingsOK) GetPayload() *models.V1RoleBindingsListAllResponse {
	return o.Payload
}

func (o *ListAllRoleBindingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1RoleBindingsListAllResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
