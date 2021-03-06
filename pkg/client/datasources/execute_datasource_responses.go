// Code generated by go-swagger; DO NOT EDIT.

package datasources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mistermx/styra-go-client/pkg/models"
)

// ExecuteDatasourceReader is a Reader for the ExecuteDatasource structure.
type ExecuteDatasourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExecuteDatasourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExecuteDatasourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewExecuteDatasourceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExecuteDatasourceOK creates a ExecuteDatasourceOK with default headers values
func NewExecuteDatasourceOK() *ExecuteDatasourceOK {
	return &ExecuteDatasourceOK{}
}

/* ExecuteDatasourceOK describes a response with status code 200, with default header values.

OK
*/
type ExecuteDatasourceOK struct {
	Payload *models.DatasourcesV1DatasourcesPostResponse
}

func (o *ExecuteDatasourceOK) Error() string {
	return fmt.Sprintf("[POST /v1/datasources/{datasource}][%d] executeDatasourceOK  %+v", 200, o.Payload)
}
func (o *ExecuteDatasourceOK) GetPayload() *models.DatasourcesV1DatasourcesPostResponse {
	return o.Payload
}

func (o *ExecuteDatasourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DatasourcesV1DatasourcesPostResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExecuteDatasourceNotFound creates a ExecuteDatasourceNotFound with default headers values
func NewExecuteDatasourceNotFound() *ExecuteDatasourceNotFound {
	return &ExecuteDatasourceNotFound{}
}

/* ExecuteDatasourceNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ExecuteDatasourceNotFound struct {
	Payload *models.MetaV1ErrorResponse
}

func (o *ExecuteDatasourceNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/datasources/{datasource}][%d] executeDatasourceNotFound  %+v", 404, o.Payload)
}
func (o *ExecuteDatasourceNotFound) GetPayload() *models.MetaV1ErrorResponse {
	return o.Payload
}

func (o *ExecuteDatasourceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetaV1ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
