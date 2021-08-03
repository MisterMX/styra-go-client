// Code generated by go-swagger; DO NOT EDIT.

package notifications_install

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mistermx/styra-go-client/pkg/models"
)

// InitiateNotificationInstallReader is a Reader for the InitiateNotificationInstall structure.
type InitiateNotificationInstallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InitiateNotificationInstallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInitiateNotificationInstallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInitiateNotificationInstallOK creates a InitiateNotificationInstallOK with default headers values
func NewInitiateNotificationInstallOK() *InitiateNotificationInstallOK {
	return &InitiateNotificationInstallOK{}
}

/* InitiateNotificationInstallOK describes a response with status code 200, with default header values.

OK
*/
type InitiateNotificationInstallOK struct {
	Payload *models.V1NotificationInstallNewStateResponse
}

func (o *InitiateNotificationInstallOK) Error() string {
	return fmt.Sprintf("[GET /v1/notifications-install/state/{type}][%d] initiateNotificationInstallOK  %+v", 200, o.Payload)
}
func (o *InitiateNotificationInstallOK) GetPayload() *models.V1NotificationInstallNewStateResponse {
	return o.Payload
}

func (o *InitiateNotificationInstallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1NotificationInstallNewStateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}