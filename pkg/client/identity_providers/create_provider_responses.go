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

// CreateProviderReader is a Reader for the CreateProvider structure.
type CreateProviderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateProviderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateProviderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewCreateProviderNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateProviderOK creates a CreateProviderOK with default headers values
func NewCreateProviderOK() *CreateProviderOK {
	return &CreateProviderOK{}
}

/* CreateProviderOK describes a response with status code 200, with default header values.

OK
*/
type CreateProviderOK struct {
	Payload *models.V1ProvidersPostResponse
}

func (o *CreateProviderOK) Error() string {
	return fmt.Sprintf("[POST /v1/identity-providers][%d] createProviderOK  %+v", 200, o.Payload)
}
func (o *CreateProviderOK) GetPayload() *models.V1ProvidersPostResponse {
	return o.Payload
}

func (o *CreateProviderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ProvidersPostResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProviderNotFound creates a CreateProviderNotFound with default headers values
func NewCreateProviderNotFound() *CreateProviderNotFound {
	return &CreateProviderNotFound{}
}

/* CreateProviderNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateProviderNotFound struct {
	Payload *models.V1ErrorResponse
}

func (o *CreateProviderNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/identity-providers][%d] createProviderNotFound  %+v", 404, o.Payload)
}
func (o *CreateProviderNotFound) GetPayload() *models.V1ErrorResponse {
	return o.Payload
}

func (o *CreateProviderNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
