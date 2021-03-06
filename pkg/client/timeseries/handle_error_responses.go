// Code generated by go-swagger; DO NOT EDIT.

package timeseries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mistermx/styra-go-client/pkg/models"
)

// HandleErrorReader is a Reader for the HandleError structure.
type HandleErrorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HandleErrorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHandleErrorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewHandleErrorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewHandleErrorOK creates a HandleErrorOK with default headers values
func NewHandleErrorOK() *HandleErrorOK {
	return &HandleErrorOK{}
}

/* HandleErrorOK describes a response with status code 200, with default header values.

OK
*/
type HandleErrorOK struct {
	Payload *models.TimeseriesV1TimeSeriesPostResponse
}

func (o *HandleErrorOK) Error() string {
	return fmt.Sprintf("[POST /v1/timeseries/error][%d] handleErrorOK  %+v", 200, o.Payload)
}
func (o *HandleErrorOK) GetPayload() *models.TimeseriesV1TimeSeriesPostResponse {
	return o.Payload
}

func (o *HandleErrorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TimeseriesV1TimeSeriesPostResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHandleErrorBadRequest creates a HandleErrorBadRequest with default headers values
func NewHandleErrorBadRequest() *HandleErrorBadRequest {
	return &HandleErrorBadRequest{}
}

/* HandleErrorBadRequest describes a response with status code 400, with default header values.

Invalid Parameter
*/
type HandleErrorBadRequest struct {
	Payload *models.MetaV1ErrorResponse
}

func (o *HandleErrorBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/timeseries/error][%d] handleErrorBadRequest  %+v", 400, o.Payload)
}
func (o *HandleErrorBadRequest) GetPayload() *models.MetaV1ErrorResponse {
	return o.Payload
}

func (o *HandleErrorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetaV1ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
