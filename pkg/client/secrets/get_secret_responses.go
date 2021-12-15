// Code generated by go-swagger; DO NOT EDIT.

package secrets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mistermx/styra-go-client/pkg/models"
)

// GetSecretReader is a Reader for the GetSecret structure.
type GetSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetSecretNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSecretOK creates a GetSecretOK with default headers values
func NewGetSecretOK() *GetSecretOK {
	return &GetSecretOK{}
}

/* GetSecretOK describes a response with status code 200, with default header values.

OK
*/
type GetSecretOK struct {
	Payload *models.SecretsV1SecretsGetResponse
}

func (o *GetSecretOK) Error() string {
	return fmt.Sprintf("[GET /v1/secrets/{secretId}][%d] getSecretOK  %+v", 200, o.Payload)
}
func (o *GetSecretOK) GetPayload() *models.SecretsV1SecretsGetResponse {
	return o.Payload
}

func (o *GetSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecretsV1SecretsGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSecretNotFound creates a GetSecretNotFound with default headers values
func NewGetSecretNotFound() *GetSecretNotFound {
	return &GetSecretNotFound{}
}

/* GetSecretNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetSecretNotFound struct {
	Payload *models.MetaV1ErrorResponse
}

func (o *GetSecretNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/secrets/{secretId}][%d] getSecretNotFound  %+v", 404, o.Payload)
}
func (o *GetSecretNotFound) GetPayload() *models.MetaV1ErrorResponse {
	return o.Payload
}

func (o *GetSecretNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetaV1ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
