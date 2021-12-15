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

// S3VerifyConfigReader is a Reader for the S3VerifyConfig structure.
type S3VerifyConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *S3VerifyConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewS3VerifyConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewS3VerifyConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewS3VerifyConfigOK creates a S3VerifyConfigOK with default headers values
func NewS3VerifyConfigOK() *S3VerifyConfigOK {
	return &S3VerifyConfigOK{}
}

/* S3VerifyConfigOK describes a response with status code 200, with default header values.

OK
*/
type S3VerifyConfigOK struct {
	Payload *models.WorkspaceV1S3VerifyResponse
}

func (o *S3VerifyConfigOK) Error() string {
	return fmt.Sprintf("[POST /v1/workspace/s3/verify-config][%d] s3VerifyConfigOK  %+v", 200, o.Payload)
}
func (o *S3VerifyConfigOK) GetPayload() *models.WorkspaceV1S3VerifyResponse {
	return o.Payload
}

func (o *S3VerifyConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkspaceV1S3VerifyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3VerifyConfigBadRequest creates a S3VerifyConfigBadRequest with default headers values
func NewS3VerifyConfigBadRequest() *S3VerifyConfigBadRequest {
	return &S3VerifyConfigBadRequest{}
}

/* S3VerifyConfigBadRequest describes a response with status code 400, with default header values.

Invalid Parameter
*/
type S3VerifyConfigBadRequest struct {
	Payload *models.MetaV1ErrorResponse
}

func (o *S3VerifyConfigBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/workspace/s3/verify-config][%d] s3VerifyConfigBadRequest  %+v", 400, o.Payload)
}
func (o *S3VerifyConfigBadRequest) GetPayload() *models.MetaV1ErrorResponse {
	return o.Payload
}

func (o *S3VerifyConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetaV1ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}