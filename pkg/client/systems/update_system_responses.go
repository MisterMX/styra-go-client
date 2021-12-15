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

// UpdateSystemReader is a Reader for the UpdateSystem structure.
type UpdateSystemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSystemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSystemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateSystemOK creates a UpdateSystemOK with default headers values
func NewUpdateSystemOK() *UpdateSystemOK {
	return &UpdateSystemOK{}
}

/* UpdateSystemOK describes a response with status code 200, with default header values.

OK
*/
type UpdateSystemOK struct {
	Payload *models.SystemsV1SystemsPutResponse
}

func (o *UpdateSystemOK) Error() string {
	return fmt.Sprintf("[PUT /v1/systems/{system}][%d] updateSystemOK  %+v", 200, o.Payload)
}
func (o *UpdateSystemOK) GetPayload() *models.SystemsV1SystemsPutResponse {
	return o.Payload
}

func (o *UpdateSystemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SystemsV1SystemsPutResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
