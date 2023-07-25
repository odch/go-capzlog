// Code generated by go-swagger; DO NOT EDIT.

package external_system_balloon_flights

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
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
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExternalSystemBalloonFlightsGetOK creates a ExternalSystemBalloonFlightsGetOK with default headers values
func NewExternalSystemBalloonFlightsGetOK() *ExternalSystemBalloonFlightsGetOK {
	return &ExternalSystemBalloonFlightsGetOK{}
}

/* ExternalSystemBalloonFlightsGetOK describes a response with status code 200, with default header values.

ExternalSystemBalloonFlightsGetOK external system balloon flights get o k
*/
type ExternalSystemBalloonFlightsGetOK struct {
	Payload *models.ExternalSystemBalloonFlightOutput
}

func (o *ExternalSystemBalloonFlightsGetOK) Error() string {
	return fmt.Sprintf("[GET /api/pel/public/externalsystemballoonflights/{id}][%d] externalSystemBalloonFlightsGetOK  %+v", 200, o.Payload)
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
