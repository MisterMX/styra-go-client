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

// ListRolesReader is a Reader for the ListRoles structure.
type ListRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListRolesOK creates a ListRolesOK with default headers values
func NewListRolesOK() *ListRolesOK {
	return &ListRolesOK{}
}

/* ListRolesOK describes a response with status code 200, with default header values.

OK
*/
type ListRolesOK struct {
	Payload *models.AuthzV2RolesListResponse
}

func (o *ListRolesOK) Error() string {
	return fmt.Sprintf("[GET /v2/authz/roles][%d] listRolesOK  %+v", 200, o.Payload)
}
func (o *ListRolesOK) GetPayload() *models.AuthzV2RolesListResponse {
	return o.Payload
}

func (o *ListRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthzV2RolesListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
