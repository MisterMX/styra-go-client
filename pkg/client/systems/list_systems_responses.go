// Code generated by go-swagger; DO NOT EDIT.

package systems

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mistermx/styra-go-client/pkg/models"
)

// ListSystemsReader is a Reader for the ListSystems structure.
type ListSystemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSystemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListSystemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListSystemsOK creates a ListSystemsOK with default headers values
func NewListSystemsOK() *ListSystemsOK {
	return &ListSystemsOK{}
}

/* ListSystemsOK describes a response with status code 200, with default header values.

OK
*/
type ListSystemsOK struct {
	Payload *models.V1SystemsListResponse
}

func (o *ListSystemsOK) Error() string {
	return fmt.Sprintf("[GET /v1/systems][%d] listSystemsOK  %+v", 200, o.Payload)
}
func (o *ListSystemsOK) GetPayload() *models.V1SystemsListResponse {
	return o.Payload
}

func (o *ListSystemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SystemsListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
