// Code generated by go-swagger; DO NOT EDIT.

package tenants_config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// TackleHandlerReader is a Reader for the TackleHandler structure.
type TackleHandlerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TackleHandlerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTackleHandlerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTackleHandlerOK creates a TackleHandlerOK with default headers values
func NewTackleHandlerOK() *TackleHandlerOK {
	return &TackleHandlerOK{}
}

/* TackleHandlerOK describes a response with status code 200, with default header values.

OK
*/
type TackleHandlerOK struct {
}

func (o *TackleHandlerOK) Error() string {
	return fmt.Sprintf("[POST /v1/tenants-config/tackle][%d] tackleHandlerOK ", 200)
}

func (o *TackleHandlerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}