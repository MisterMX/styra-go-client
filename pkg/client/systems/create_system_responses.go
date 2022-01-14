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

// CreateSystemReader is a Reader for the CreateSystem structure.
type CreateSystemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSystemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateSystemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateSystemOK creates a CreateSystemOK with default headers values
func NewCreateSystemOK() *CreateSystemOK {
	return &CreateSystemOK{}
}

/* CreateSystemOK describes a response with status code 200, with default header values.

OK
*/
type CreateSystemOK struct {
	Payload *models.SystemsV1SystemsPostResponse
}

func (o *CreateSystemOK) Error() string {
	return fmt.Sprintf("[POST /v1/systems][%d] createSystemOK  %+v", 200, o.Payload)
}
func (o *CreateSystemOK) GetPayload() *models.SystemsV1SystemsPostResponse {
	return o.Payload
}

func (o *CreateSystemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SystemsV1SystemsPostResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
