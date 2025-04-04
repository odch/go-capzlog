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

// ExternalSystemBalloonFlightsGetReader is a Reader for the ExternalSystemBalloonFlightsGet structure.
type ExternalSystemBalloonFlightsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExternalSystemBalloonFlightsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExternalSystemBalloonFlightsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /api/pel/public/externalsystemballoonflights/{id}] ExternalSystemBalloonFlights_Get", response, response.Code())
	}
}

// NewExternalSystemBalloonFlightsGetOK creates a ExternalSystemBalloonFlightsGetOK with default headers values
func NewExternalSystemBalloonFlightsGetOK() *ExternalSystemBalloonFlightsGetOK {
	return &ExternalSystemBalloonFlightsGetOK{}
}

/*
ExternalSystemBalloonFlightsGetOK describes a response with status code 200, with default header values.

ExternalSystemBalloonFlightsGetOK external system balloon flights get o k
*/
type ExternalSystemBalloonFlightsGetOK struct {
	Payload *models.ExternalSystemBalloonFlightOutput
}

// IsSuccess returns true when this external system balloon flights get o k response has a 2xx status code
func (o *ExternalSystemBalloonFlightsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this external system balloon flights get o k response has a 3xx status code
func (o *ExternalSystemBalloonFlightsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this external system balloon flights get o k response has a 4xx status code
func (o *ExternalSystemBalloonFlightsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this external system balloon flights get o k response has a 5xx status code
func (o *ExternalSystemBalloonFlightsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this external system balloon flights get o k response a status code equal to that given
func (o *ExternalSystemBalloonFlightsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the external system balloon flights get o k response
func (o *ExternalSystemBalloonFlightsGetOK) Code() int {
	return 200
}

func (o *ExternalSystemBalloonFlightsGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/pel/public/externalsystemballoonflights/{id}][%d] externalSystemBalloonFlightsGetOK %s", 200, payload)
}

func (o *ExternalSystemBalloonFlightsGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/pel/public/externalsystemballoonflights/{id}][%d] externalSystemBalloonFlightsGetOK %s", 200, payload)
}

func (o *ExternalSystemBalloonFlightsGetOK) GetPayload() *models.ExternalSystemBalloonFlightOutput {
	return o.Payload
}

func (o *ExternalSystemBalloonFlightsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExternalSystemBalloonFlightOutput)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
