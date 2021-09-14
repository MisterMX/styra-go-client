// Code generated by go-swagger; DO NOT EDIT.

package relay_server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RelayPUTReader is a Reader for the RelayPUT structure.
type RelayPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RelayPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRelayPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRelayPUTOK creates a RelayPUTOK with default headers values
func NewRelayPUTOK() *RelayPUTOK {
	return &RelayPUTOK{}
}

/* RelayPUTOK describes a response with status code 200, with default header values.

OK
*/
type RelayPUTOK struct {
}

func (o *RelayPUTOK) Error() string {
	return fmt.Sprintf("[PUT /v1/relay/{key}/{path}][%d] relayPUTOK ", 200)
}

func (o *RelayPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
