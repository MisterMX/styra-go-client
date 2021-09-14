// Code generated by go-swagger; DO NOT EDIT.

package relay_server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RelayPATCHReader is a Reader for the RelayPATCH structure.
type RelayPATCHReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RelayPATCHReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRelayPATCHOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRelayPATCHOK creates a RelayPATCHOK with default headers values
func NewRelayPATCHOK() *RelayPATCHOK {
	return &RelayPATCHOK{}
}

/* RelayPATCHOK describes a response with status code 200, with default header values.

OK
*/
type RelayPATCHOK struct {
}

func (o *RelayPATCHOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/relay/{key}/{path}][%d] relayPATCHOK ", 200)
}

func (o *RelayPATCHOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
