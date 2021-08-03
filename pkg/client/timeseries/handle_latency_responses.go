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

// HandleLatencyReader is a Reader for the HandleLatency structure.
type HandleLatencyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HandleLatencyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHandleLatencyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewHandleLatencyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewHandleLatencyOK creates a HandleLatencyOK with default headers values
func NewHandleLatencyOK() *HandleLatencyOK {
	return &HandleLatencyOK{}
}

/* HandleLatencyOK describes a response with status code 200, with default header values.

OK
*/
type HandleLatencyOK struct {
	Payload *models.V1TimeSeriesPostResponse
}

func (o *HandleLatencyOK) Error() string {
	return fmt.Sprintf("[POST /v1/timeseries/latency][%d] handleLatencyOK  %+v", 200, o.Payload)
}
func (o *HandleLatencyOK) GetPayload() *models.V1TimeSeriesPostResponse {
	return o.Payload
}

func (o *HandleLatencyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1TimeSeriesPostResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHandleLatencyBadRequest creates a HandleLatencyBadRequest with default headers values
func NewHandleLatencyBadRequest() *HandleLatencyBadRequest {
	return &HandleLatencyBadRequest{}
}

/* HandleLatencyBadRequest describes a response with status code 400, with default header values.

Invalid Parameter
*/
type HandleLatencyBadRequest struct {
	Payload *models.V1ErrorResponse
}

func (o *HandleLatencyBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/timeseries/latency][%d] handleLatencyBadRequest  %+v", 400, o.Payload)
}
func (o *HandleLatencyBadRequest) GetPayload() *models.V1ErrorResponse {
	return o.Payload
}

func (o *HandleLatencyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
