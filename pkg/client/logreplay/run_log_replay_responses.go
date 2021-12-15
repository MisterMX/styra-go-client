// Code generated by go-swagger; DO NOT EDIT.

package logreplay

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mistermx/styra-go-client/pkg/models"
)

// RunLogReplayReader is a Reader for the RunLogReplay structure.
type RunLogReplayReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RunLogReplayReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRunLogReplayOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRunLogReplayOK creates a RunLogReplayOK with default headers values
func NewRunLogReplayOK() *RunLogReplayOK {
	return &RunLogReplayOK{}
}

/* RunLogReplayOK describes a response with status code 200, with default header values.

OK
*/
type RunLogReplayOK struct {
	Payload *models.LogreplayV1ReplayResult
}

func (o *RunLogReplayOK) Error() string {
	return fmt.Sprintf("[POST /v1/logreplay][%d] runLogReplayOK  %+v", 200, o.Payload)
}
func (o *RunLogReplayOK) GetPayload() *models.LogreplayV1ReplayResult {
	return o.Payload
}

func (o *RunLogReplayOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LogreplayV1ReplayResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
