// Code generated by go-swagger; DO NOT EDIT.

package tokens

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mistermx/styra-go-client/pkg/models"
)

// ListTokensReader is a Reader for the ListTokens structure.
type ListTokensReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListTokensReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListTokensOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListTokensOK creates a ListTokensOK with default headers values
func NewListTokensOK() *ListTokensOK {
	return &ListTokensOK{}
}

/* ListTokensOK describes a response with status code 200, with default header values.

OK
*/
type ListTokensOK struct {
	Payload *models.TokensV1TokensListResponse
}

func (o *ListTokensOK) Error() string {
	return fmt.Sprintf("[GET /v1/tokens][%d] listTokensOK  %+v", 200, o.Payload)
}
func (o *ListTokensOK) GetPayload() *models.TokensV1TokensListResponse {
	return o.Payload
}

func (o *ListTokensOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TokensV1TokensListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
