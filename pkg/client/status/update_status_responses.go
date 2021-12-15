// Code generated by go-swagger; DO NOT EDIT.

package status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mistermx/styra-go-client/pkg/models"
)

// UpdateStatusReader is a Reader for the UpdateStatus structure.
type UpdateStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateStatusOK creates a UpdateStatusOK with default headers values
func NewUpdateStatusOK() *UpdateStatusOK {
	return &UpdateStatusOK{}
}

/* UpdateStatusOK describes a response with status code 200, with default header values.

OK
*/
type UpdateStatusOK struct {
	Payload *models.StatusV1StatusPostResponse
}

func (o *UpdateStatusOK) Error() string {
	return fmt.Sprintf("[POST /v1/status][%d] updateStatusOK  %+v", 200, o.Payload)
}
func (o *UpdateStatusOK) GetPayload() *models.StatusV1StatusPostResponse {
	return o.Payload
}

func (o *UpdateStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StatusV1StatusPostResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
