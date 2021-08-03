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

// DeleteProviderReader is a Reader for the DeleteProvider structure.
type DeleteProviderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProviderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteProviderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteProviderNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteProviderOK creates a DeleteProviderOK with default headers values
func NewDeleteProviderOK() *DeleteProviderOK {
	return &DeleteProviderOK{}
}

/* DeleteProviderOK describes a response with status code 200, with default header values.

OK
*/
type DeleteProviderOK struct {
	Payload *models.V1ProvidersDeleteResponse
}

func (o *DeleteProviderOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/identity-providers/{providerId}][%d] deleteProviderOK  %+v", 200, o.Payload)
}
func (o *DeleteProviderOK) GetPayload() *models.V1ProvidersDeleteResponse {
	return o.Payload
}

func (o *DeleteProviderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ProvidersDeleteResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProviderNotFound creates a DeleteProviderNotFound with default headers values
func NewDeleteProviderNotFound() *DeleteProviderNotFound {
	return &DeleteProviderNotFound{}
}

/* DeleteProviderNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteProviderNotFound struct {
	Payload *models.V1ErrorResponse
}

func (o *DeleteProviderNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/identity-providers/{providerId}][%d] deleteProviderNotFound  %+v", 404, o.Payload)
}
func (o *DeleteProviderNotFound) GetPayload() *models.V1ErrorResponse {
	return o.Payload
}

func (o *DeleteProviderNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
