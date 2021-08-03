// Code generated by go-swagger; DO NOT EDIT.

package tenants_config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mistermx/styra-go-client/pkg/models"
)

// CreateTenantConfigReader is a Reader for the CreateTenantConfig structure.
type CreateTenantConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTenantConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateTenantConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateTenantConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateTenantConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateTenantConfigOK creates a CreateTenantConfigOK with default headers values
func NewCreateTenantConfigOK() *CreateTenantConfigOK {
	return &CreateTenantConfigOK{}
}

/* CreateTenantConfigOK describes a response with status code 200, with default header values.

OK
*/
type CreateTenantConfigOK struct {
	Payload *models.V1TenantConfigPostResponse
}

func (o *CreateTenantConfigOK) Error() string {
	return fmt.Sprintf("[POST /v1/tenants-config][%d] createTenantConfigOK  %+v", 200, o.Payload)
}
func (o *CreateTenantConfigOK) GetPayload() *models.V1TenantConfigPostResponse {
	return o.Payload
}

func (o *CreateTenantConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1TenantConfigPostResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTenantConfigBadRequest creates a CreateTenantConfigBadRequest with default headers values
func NewCreateTenantConfigBadRequest() *CreateTenantConfigBadRequest {
	return &CreateTenantConfigBadRequest{}
}

/* CreateTenantConfigBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateTenantConfigBadRequest struct {
	Payload *models.V1ErrorResponse
}

func (o *CreateTenantConfigBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/tenants-config][%d] createTenantConfigBadRequest  %+v", 400, o.Payload)
}
func (o *CreateTenantConfigBadRequest) GetPayload() *models.V1ErrorResponse {
	return o.Payload
}

func (o *CreateTenantConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTenantConfigNotFound creates a CreateTenantConfigNotFound with default headers values
func NewCreateTenantConfigNotFound() *CreateTenantConfigNotFound {
	return &CreateTenantConfigNotFound{}
}

/* CreateTenantConfigNotFound describes a response with status code 404, with default header values.

Not found
*/
type CreateTenantConfigNotFound struct {
	Payload *models.V1ErrorResponse
}

func (o *CreateTenantConfigNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/tenants-config][%d] createTenantConfigNotFound  %+v", 404, o.Payload)
}
func (o *CreateTenantConfigNotFound) GetPayload() *models.V1ErrorResponse {
	return o.Payload
}

func (o *CreateTenantConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}