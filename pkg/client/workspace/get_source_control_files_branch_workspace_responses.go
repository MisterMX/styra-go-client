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

// GetSourceControlFilesBranchWorkspaceReader is a Reader for the GetSourceControlFilesBranchWorkspace structure.
type GetSourceControlFilesBranchWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSourceControlFilesBranchWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSourceControlFilesBranchWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetSourceControlFilesBranchWorkspaceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSourceControlFilesBranchWorkspaceOK creates a GetSourceControlFilesBranchWorkspaceOK with default headers values
func NewGetSourceControlFilesBranchWorkspaceOK() *GetSourceControlFilesBranchWorkspaceOK {
	return &GetSourceControlFilesBranchWorkspaceOK{}
}

/* GetSourceControlFilesBranchWorkspaceOK describes a response with status code 200, with default header values.

OK
*/
type GetSourceControlFilesBranchWorkspaceOK struct {
	Payload *models.GitV1GetFilesResponse
}

func (o *GetSourceControlFilesBranchWorkspaceOK) Error() string {
	return fmt.Sprintf("[GET /v1/workspace/{id}/branch][%d] getSourceControlFilesBranchWorkspaceOK  %+v", 200, o.Payload)
}
func (o *GetSourceControlFilesBranchWorkspaceOK) GetPayload() *models.GitV1GetFilesResponse {
	return o.Payload
}

func (o *GetSourceControlFilesBranchWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GitV1GetFilesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSourceControlFilesBranchWorkspaceNotFound creates a GetSourceControlFilesBranchWorkspaceNotFound with default headers values
func NewGetSourceControlFilesBranchWorkspaceNotFound() *GetSourceControlFilesBranchWorkspaceNotFound {
	return &GetSourceControlFilesBranchWorkspaceNotFound{}
}

/* GetSourceControlFilesBranchWorkspaceNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetSourceControlFilesBranchWorkspaceNotFound struct {
	Payload *models.MetaV1ErrorResponse
}

func (o *GetSourceControlFilesBranchWorkspaceNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/workspace/{id}/branch][%d] getSourceControlFilesBranchWorkspaceNotFound  %+v", 404, o.Payload)
}
func (o *GetSourceControlFilesBranchWorkspaceNotFound) GetPayload() *models.MetaV1ErrorResponse {
	return o.Payload
}

func (o *GetSourceControlFilesBranchWorkspaceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetaV1ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
