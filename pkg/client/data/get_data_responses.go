// Code generated by go-swagger; DO NOT EDIT.

package data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mistermx/styra-go-client/pkg/models"
)

// GetDataReader is a Reader for the GetData structure.
type GetDataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetDataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDataOK creates a GetDataOK with default headers values
func NewGetDataOK() *GetDataOK {
	return &GetDataOK{}
}

/* GetDataOK describes a response with status code 200, with default header values.

OK
*/
type GetDataOK struct {
	Payload *models.V1DataResponse
}

func (o *GetDataOK) Error() string {
	return fmt.Sprintf("[GET /v1/data/{name}][%d] getDataOK  %+v", 200, o.Payload)
}
func (o *GetDataOK) GetPayload() *models.V1DataResponse {
	return o.Payload
}

func (o *GetDataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1DataResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataNotFound creates a GetDataNotFound with default headers values
func NewGetDataNotFound() *GetDataNotFound {
	return &GetDataNotFound{}
}

/* GetDataNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetDataNotFound struct {
	Payload *models.V1ErrorResponse
}

func (o *GetDataNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/data/{name}][%d] getDataNotFound  %+v", 404, o.Payload)
}
func (o *GetDataNotFound) GetPayload() *models.V1ErrorResponse {
	return o.Payload
}

func (o *GetDataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
