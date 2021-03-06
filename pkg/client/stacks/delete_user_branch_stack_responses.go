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

// DeleteUserBranchStackReader is a Reader for the DeleteUserBranchStack structure.
type DeleteUserBranchStackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserBranchStackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteUserBranchStackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteUserBranchStackNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteUserBranchStackOK creates a DeleteUserBranchStackOK with default headers values
func NewDeleteUserBranchStackOK() *DeleteUserBranchStackOK {
	return &DeleteUserBranchStackOK{}
}

/* DeleteUserBranchStackOK describes a response with status code 200, with default header values.

OK
*/
type DeleteUserBranchStackOK struct {
	Payload *models.GitV1DeleteBranchResponse
}

func (o *DeleteUserBranchStackOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/stacks/{id}/branch][%d] deleteUserBranchStackOK  %+v", 200, o.Payload)
}
func (o *DeleteUserBranchStackOK) GetPayload() *models.GitV1DeleteBranchResponse {
	return o.Payload
}

func (o *DeleteUserBranchStackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GitV1DeleteBranchResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserBranchStackNotFound creates a DeleteUserBranchStackNotFound with default headers values
func NewDeleteUserBranchStackNotFound() *DeleteUserBranchStackNotFound {
	return &DeleteUserBranchStackNotFound{}
}

/* DeleteUserBranchStackNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteUserBranchStackNotFound struct {
	Payload *models.MetaV1ErrorResponse
}

func (o *DeleteUserBranchStackNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/stacks/{id}/branch][%d] deleteUserBranchStackNotFound  %+v", 404, o.Payload)
}
func (o *DeleteUserBranchStackNotFound) GetPayload() *models.MetaV1ErrorResponse {
	return o.Payload
}

func (o *DeleteUserBranchStackNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetaV1ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
