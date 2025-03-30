// Code generated by go-swagger; DO NOT EDIT.

package external_system_balloon_flights

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExternalSystemBalloonFlightsDeleteReader is a Reader for the ExternalSystemBalloonFlightsDelete structure.
type ExternalSystemBalloonFlightsDeleteReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *ExternalSystemBalloonFlightsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExternalSystemBalloonFlightsDeleteOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /api/pel/public/externalsystemballoonflights/{id}] ExternalSystemBalloonFlights_Delete", response, response.Code())
	}
}

// NewExternalSystemBalloonFlightsDeleteOK creates a ExternalSystemBalloonFlightsDeleteOK with default headers values
func NewExternalSystemBalloonFlightsDeleteOK(writer io.Writer) *ExternalSystemBalloonFlightsDeleteOK {
	return &ExternalSystemBalloonFlightsDeleteOK{

		Payload: writer,
	}
}

/*
ExternalSystemBalloonFlightsDeleteOK describes a response with status code 200, with default header values.

ExternalSystemBalloonFlightsDeleteOK external system balloon flights delete o k
*/
type ExternalSystemBalloonFlightsDeleteOK struct {
	Payload io.Writer
}

// IsSuccess returns true when this external system balloon flights delete o k response has a 2xx status code
func (o *ExternalSystemBalloonFlightsDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this external system balloon flights delete o k response has a 3xx status code
func (o *ExternalSystemBalloonFlightsDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this external system balloon flights delete o k response has a 4xx status code
func (o *ExternalSystemBalloonFlightsDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this external system balloon flights delete o k response has a 5xx status code
func (o *ExternalSystemBalloonFlightsDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this external system balloon flights delete o k response a status code equal to that given
func (o *ExternalSystemBalloonFlightsDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the external system balloon flights delete o k response
func (o *ExternalSystemBalloonFlightsDeleteOK) Code() int {
	return 200
}

func (o *ExternalSystemBalloonFlightsDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /api/pel/public/externalsystemballoonflights/{id}][%d] externalSystemBalloonFlightsDeleteOK", 200)
}

func (o *ExternalSystemBalloonFlightsDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /api/pel/public/externalsystemballoonflights/{id}][%d] externalSystemBalloonFlightsDeleteOK", 200)
}

func (o *ExternalSystemBalloonFlightsDeleteOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *ExternalSystemBalloonFlightsDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
