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

// GetSourceControlFilesMasterSystemReader is a Reader for the GetSourceControlFilesMasterSystem structure.
type GetSourceControlFilesMasterSystemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSourceControlFilesMasterSystemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSourceControlFilesMasterSystemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetSourceControlFilesMasterSystemNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSourceControlFilesMasterSystemOK creates a GetSourceControlFilesMasterSystemOK with default headers values
func NewGetSourceControlFilesMasterSystemOK() *GetSourceControlFilesMasterSystemOK {
	return &GetSourceControlFilesMasterSystemOK{}
}

/* GetSourceControlFilesMasterSystemOK describes a response with status code 200, with default header values.

OK
*/
type GetSourceControlFilesMasterSystemOK struct {
	Payload *models.GitV1GetFilesResponse
}

func (o *GetSourceControlFilesMasterSystemOK) Error() string {
	return fmt.Sprintf("[GET /v1/systems/{id}/master][%d] getSourceControlFilesMasterSystemOK  %+v", 200, o.Payload)
}
func (o *GetSourceControlFilesMasterSystemOK) GetPayload() *models.GitV1GetFilesResponse {
	return o.Payload
}

func (o *GetSourceControlFilesMasterSystemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GitV1GetFilesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSourceControlFilesMasterSystemNotFound creates a GetSourceControlFilesMasterSystemNotFound with default headers values
func NewGetSourceControlFilesMasterSystemNotFound() *GetSourceControlFilesMasterSystemNotFound {
	return &GetSourceControlFilesMasterSystemNotFound{}
}

/* GetSourceControlFilesMasterSystemNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetSourceControlFilesMasterSystemNotFound struct {
	Payload *models.MetaV1ErrorResponse
}

func (o *GetSourceControlFilesMasterSystemNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/systems/{id}/master][%d] getSourceControlFilesMasterSystemNotFound  %+v", 404, o.Payload)
}
func (o *GetSourceControlFilesMasterSystemNotFound) GetPayload() *models.MetaV1ErrorResponse {
	return o.Payload
}

func (o *GetSourceControlFilesMasterSystemNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetaV1ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
