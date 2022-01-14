// Code generated by go-swagger; DO NOT EDIT.

package stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mistermx/styra-go-client/pkg/models"
)

// ListStacksReader is a Reader for the ListStacks structure.
type ListStacksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListStacksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListStacksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListStacksOK creates a ListStacksOK with default headers values
func NewListStacksOK() *ListStacksOK {
	return &ListStacksOK{}
}

/* ListStacksOK describes a response with status code 200, with default header values.

OK
*/
type ListStacksOK struct {
	Payload *models.StacksV1StacksListResponse
}

func (o *ListStacksOK) Error() string {
	return fmt.Sprintf("[GET /v1/stacks][%d] listStacksOK  %+v", 200, o.Payload)
}
func (o *ListStacksOK) GetPayload() *models.StacksV1StacksListResponse {
	return o.Payload
}

func (o *ListStacksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StacksV1StacksListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
