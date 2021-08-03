// Code generated by go-swagger; DO NOT EDIT.

package identity_providers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mistermx/styra-go-client/pkg/models"
)

// UpdateProviderReader is a Reader for the UpdateProvider structure.
type UpdateProviderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProviderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateProviderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUpdateProviderNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateProviderOK creates a UpdateProviderOK with default headers values
func NewUpdateProviderOK() *UpdateProviderOK {
	return &UpdateProviderOK{}
}

/* UpdateProviderOK describes a response with status code 200, with default header values.

OK
*/
type UpdateProviderOK struct {
	Payload *models.V1ProvidersPutResponse
}

func (o *UpdateProviderOK) Error() string {
	return fmt.Sprintf("[PUT /v1/identity-providers/{providerId}][%d] updateProviderOK  %+v", 200, o.Payload)
}
func (o *UpdateProviderOK) GetPayload() *models.V1ProvidersPutResponse {
	return o.Payload
}

func (o *UpdateProviderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ProvidersPutResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProviderNotFound creates a UpdateProviderNotFound with default headers values
func NewUpdateProviderNotFound() *UpdateProviderNotFound {
	return &UpdateProviderNotFound{}
}

/* UpdateProviderNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateProviderNotFound struct {
	Payload *models.V1ErrorResponse
}

func (o *UpdateProviderNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/identity-providers/{providerId}][%d] updateProviderNotFound  %+v", 404, o.Payload)
}
func (o *UpdateProviderNotFound) GetPayload() *models.V1ErrorResponse {
	return o.Payload
}

func (o *UpdateProviderNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
