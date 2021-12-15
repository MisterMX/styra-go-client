// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mistermx/styra-go-client/pkg/models"
)

// DeletePolicyReader is a Reader for the DeletePolicy structure.
type DeletePolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeletePolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeletePolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeletePolicyOK creates a DeletePolicyOK with default headers values
func NewDeletePolicyOK() *DeletePolicyOK {
	return &DeletePolicyOK{}
}

/* DeletePolicyOK describes a response with status code 200, with default header values.

OK
*/
type DeletePolicyOK struct {
	Payload *models.PoliciesV1PolicyDeleteResponse
}

func (o *DeletePolicyOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/policies/{policy}][%d] deletePolicyOK  %+v", 200, o.Payload)
}
func (o *DeletePolicyOK) GetPayload() *models.PoliciesV1PolicyDeleteResponse {
	return o.Payload
}

func (o *DeletePolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PoliciesV1PolicyDeleteResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePolicyNotFound creates a DeletePolicyNotFound with default headers values
func NewDeletePolicyNotFound() *DeletePolicyNotFound {
	return &DeletePolicyNotFound{}
}

/* DeletePolicyNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeletePolicyNotFound struct {
	Payload *models.MetaV1ErrorResponse
}

func (o *DeletePolicyNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/policies/{policy}][%d] deletePolicyNotFound  %+v", 404, o.Payload)
}
func (o *DeletePolicyNotFound) GetPayload() *models.MetaV1ErrorResponse {
	return o.Payload
}

func (o *DeletePolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetaV1ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
