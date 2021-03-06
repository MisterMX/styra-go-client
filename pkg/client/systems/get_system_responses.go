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

// GetSystemReader is a Reader for the GetSystem structure.
type GetSystemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSystemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSystemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetSystemNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSystemOK creates a GetSystemOK with default headers values
func NewGetSystemOK() *GetSystemOK {
	return &GetSystemOK{}
}

/* GetSystemOK describes a response with status code 200, with default header values.

OK
*/
type GetSystemOK struct {
	Payload *models.SystemsV1SystemsGetResponse
}

func (o *GetSystemOK) Error() string {
	return fmt.Sprintf("[GET /v1/systems/{system}][%d] getSystemOK  %+v", 200, o.Payload)
}
func (o *GetSystemOK) GetPayload() *models.SystemsV1SystemsGetResponse {
	return o.Payload
}

func (o *GetSystemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SystemsV1SystemsGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSystemNotFound creates a GetSystemNotFound with default headers values
func NewGetSystemNotFound() *GetSystemNotFound {
	return &GetSystemNotFound{}
}

/* GetSystemNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetSystemNotFound struct {
	Payload *models.MetaV1ErrorResponse
}

func (o *GetSystemNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/systems/{system}][%d] getSystemNotFound  %+v", 404, o.Payload)
}
func (o *GetSystemNotFound) GetPayload() *models.MetaV1ErrorResponse {
	return o.Payload
}

func (o *GetSystemNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetaV1ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
