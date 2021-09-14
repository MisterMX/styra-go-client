// Code generated by go-swagger; DO NOT EDIT.

package relay_server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RelayPOSTReader is a Reader for the RelayPOST structure.
type RelayPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RelayPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRelayPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRelayPOSTOK creates a RelayPOSTOK with default headers values
func NewRelayPOSTOK() *RelayPOSTOK {
	return &RelayPOSTOK{}
}

/* RelayPOSTOK describes a response with status code 200, with default header values.

OK
*/
type RelayPOSTOK struct {
}

func (o *RelayPOSTOK) Error() string {
	return fmt.Sprintf("[POST /v1/relay/{key}/{path}][%d] relayPOSTOK ", 200)
}

func (o *RelayPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
