// Code generated by go-swagger; DO NOT EDIT.

package tenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mistermx/styra-go-client/pkg/models"
)

// TenantActiveReader is a Reader for the TenantActive structure.
type TenantActiveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenantActiveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenantActiveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewTenantActiveNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTenantActiveOK creates a TenantActiveOK with default headers values
func NewTenantActiveOK() *TenantActiveOK {
	return &TenantActiveOK{}
}

/* TenantActiveOK describes a response with status code 200, with default header values.

OK
*/
type TenantActiveOK struct {
	Payload *models.V1ResponseHeader
}

func (o *TenantActiveOK) Error() string {
	return fmt.Sprintf("[GET /v1/tenants/active][%d] tenantActiveOK  %+v", 200, o.Payload)
}
func (o *TenantActiveOK) GetPayload() *models.V1ResponseHeader {
	return o.Payload
}

func (o *TenantActiveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ResponseHeader)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenantActiveNotFound creates a TenantActiveNotFound with default headers values
func NewTenantActiveNotFound() *TenantActiveNotFound {
	return &TenantActiveNotFound{}
}

/* TenantActiveNotFound describes a response with status code 404, with default header values.

Not found
*/
type TenantActiveNotFound struct {
	Payload *models.V1ErrorResponse
}

func (o *TenantActiveNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/tenants/active][%d] tenantActiveNotFound  %+v", 404, o.Payload)
}
func (o *TenantActiveNotFound) GetPayload() *models.V1ErrorResponse {
	return o.Payload
}

func (o *TenantActiveNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}