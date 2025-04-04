// Code generated by go-swagger; DO NOT EDIT.

package external_system_balloon_flights

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/odch/go-capzlog/models"
)

// ExternalSystemBalloonFlightsPutReader is a Reader for the ExternalSystemBalloonFlightsPut structure.
type ExternalSystemBalloonFlightsPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExternalSystemBalloonFlightsPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExternalSystemBalloonFlightsPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /api/pel/public/externalsystemballoonflights/{id}] ExternalSystemBalloonFlights_Put", response, response.Code())
	}
}

// NewExternalSystemBalloonFlightsPutOK creates a ExternalSystemBalloonFlightsPutOK with default headers values
func NewExternalSystemBalloonFlightsPutOK() *ExternalSystemBalloonFlightsPutOK {
	return &ExternalSystemBalloonFlightsPutOK{}
}

/*
ExternalSystemBalloonFlightsPutOK describes a response with status code 200, with default header values.

ExternalSystemBalloonFlightsPutOK external system balloon flights put o k
*/
type ExternalSystemBalloonFlightsPutOK struct {
	Payload *models.ExternalSystemBalloonFlightOutput
}

// IsSuccess returns true when this external system balloon flights put o k response has a 2xx status code
func (o *ExternalSystemBalloonFlightsPutOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this external system balloon flights put o k response has a 3xx status code
func (o *ExternalSystemBalloonFlightsPutOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this external system balloon flights put o k response has a 4xx status code
func (o *ExternalSystemBalloonFlightsPutOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this external system balloon flights put o k response has a 5xx status code
func (o *ExternalSystemBalloonFlightsPutOK) IsServerError() bool {
	return false
}

// IsCode returns true when this external system balloon flights put o k response a status code equal to that given
func (o *ExternalSystemBalloonFlightsPutOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the external system balloon flights put o k response
func (o *ExternalSystemBalloonFlightsPutOK) Code() int {
	return 200
}

func (o *ExternalSystemBalloonFlightsPutOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/pel/public/externalsystemballoonflights/{id}][%d] externalSystemBalloonFlightsPutOK %s", 200, payload)
}

func (o *ExternalSystemBalloonFlightsPutOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/pel/public/externalsystemballoonflights/{id}][%d] externalSystemBalloonFlightsPutOK %s", 200, payload)
}

func (o *ExternalSystemBalloonFlightsPutOK) GetPayload() *models.ExternalSystemBalloonFlightOutput {
	return o.Payload
}

func (o *ExternalSystemBalloonFlightsPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExternalSystemBalloonFlightOutput)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
