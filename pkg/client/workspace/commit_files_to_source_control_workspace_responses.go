// Code generated by go-swagger; DO NOT EDIT.

package workspace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mistermx/styra-go-client/pkg/models"
)

// CommitFilesToSourceControlWorkspaceReader is a Reader for the CommitFilesToSourceControlWorkspace structure.
type CommitFilesToSourceControlWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CommitFilesToSourceControlWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCommitFilesToSourceControlWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewCommitFilesToSourceControlWorkspaceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCommitFilesToSourceControlWorkspaceOK creates a CommitFilesToSourceControlWorkspaceOK with default headers values
func NewCommitFilesToSourceControlWorkspaceOK() *CommitFilesToSourceControlWorkspaceOK {
	return &CommitFilesToSourceControlWorkspaceOK{}
}

/* CommitFilesToSourceControlWorkspaceOK describes a response with status code 200, with default header values.

OK
*/
type CommitFilesToSourceControlWorkspaceOK struct {
	Payload *models.GitV1PostCommitResponse
}

func (o *CommitFilesToSourceControlWorkspaceOK) Error() string {
	return fmt.Sprintf("[POST /v1/workspace/{id}/commits][%d] commitFilesToSourceControlWorkspaceOK  %+v", 200, o.Payload)
}
func (o *CommitFilesToSourceControlWorkspaceOK) GetPayload() *models.GitV1PostCommitResponse {
	return o.Payload
}

func (o *CommitFilesToSourceControlWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GitV1PostCommitResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCommitFilesToSourceControlWorkspaceNotFound creates a CommitFilesToSourceControlWorkspaceNotFound with default headers values
func NewCommitFilesToSourceControlWorkspaceNotFound() *CommitFilesToSourceControlWorkspaceNotFound {
	return &CommitFilesToSourceControlWorkspaceNotFound{}
}

/* CommitFilesToSourceControlWorkspaceNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CommitFilesToSourceControlWorkspaceNotFound struct {
	Payload *models.MetaV1ErrorResponse
}

func (o *CommitFilesToSourceControlWorkspaceNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/workspace/{id}/commits][%d] commitFilesToSourceControlWorkspaceNotFound  %+v", 404, o.Payload)
}
func (o *CommitFilesToSourceControlWorkspaceNotFound) GetPayload() *models.MetaV1ErrorResponse {
	return o.Payload
}

func (o *CommitFilesToSourceControlWorkspaceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetaV1ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
