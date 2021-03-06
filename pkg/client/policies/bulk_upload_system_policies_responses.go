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

// BulkUploadSystemPoliciesReader is a Reader for the BulkUploadSystemPolicies structure.
type BulkUploadSystemPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BulkUploadSystemPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBulkUploadSystemPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBulkUploadSystemPoliciesOK creates a BulkUploadSystemPoliciesOK with default headers values
func NewBulkUploadSystemPoliciesOK() *BulkUploadSystemPoliciesOK {
	return &BulkUploadSystemPoliciesOK{}
}

/* BulkUploadSystemPoliciesOK describes a response with status code 200, with default header values.

OK
*/
type BulkUploadSystemPoliciesOK struct {
	Payload *models.PoliciesV1PoliciesBulkUploadResponse
}

func (o *BulkUploadSystemPoliciesOK) Error() string {
	return fmt.Sprintf("[POST /v1/policies/systems/{system}][%d] bulkUploadSystemPoliciesOK  %+v", 200, o.Payload)
}
func (o *BulkUploadSystemPoliciesOK) GetPayload() *models.PoliciesV1PoliciesBulkUploadResponse {
	return o.Payload
}

func (o *BulkUploadSystemPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PoliciesV1PoliciesBulkUploadResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
