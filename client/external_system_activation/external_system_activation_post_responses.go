// Code generated by go-swagger; DO NOT EDIT.

package external_system_activation

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

// ExternalSystemActivationPostReader is a Reader for the ExternalSystemActivationPost structure.
type ExternalSystemActivationPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExternalSystemActivationPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExternalSystemActivationPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewExternalSystemActivationPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/pel/public/externalsystemactivation] ExternalSystemActivation_Post", response, response.Code())
	}
}

// NewExternalSystemActivationPostOK creates a ExternalSystemActivationPostOK with default headers values
func NewExternalSystemActivationPostOK() *ExternalSystemActivationPostOK {
	return &ExternalSystemActivationPostOK{}
}

/*
ExternalSystemActivationPostOK describes a response with status code 200, with default header values.

ExternalSystemActivationPostOK external system activation post o k
*/
type ExternalSystemActivationPostOK struct {
	Payload *models.ExternalSystemConnectionActivationOutput
}

// IsSuccess returns true when this external system activation post o k response has a 2xx status code
func (o *ExternalSystemActivationPostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this external system activation post o k response has a 3xx status code
func (o *ExternalSystemActivationPostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this external system activation post o k response has a 4xx status code
func (o *ExternalSystemActivationPostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this external system activation post o k response has a 5xx status code
func (o *ExternalSystemActivationPostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this external system activation post o k response a status code equal to that given
func (o *ExternalSystemActivationPostOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the external system activation post o k response
func (o *ExternalSystemActivationPostOK) Code() int {
	return 200
}

func (o *ExternalSystemActivationPostOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/pel/public/externalsystemactivation][%d] externalSystemActivationPostOK %s", 200, payload)
}

func (o *ExternalSystemActivationPostOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/pel/public/externalsystemactivation][%d] externalSystemActivationPostOK %s", 200, payload)
}

func (o *ExternalSystemActivationPostOK) GetPayload() *models.ExternalSystemConnectionActivationOutput {
	return o.Payload
}

func (o *ExternalSystemActivationPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExternalSystemConnectionActivationOutput)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExternalSystemActivationPostBadRequest creates a ExternalSystemActivationPostBadRequest with default headers values
func NewExternalSystemActivationPostBadRequest() *ExternalSystemActivationPostBadRequest {
	return &ExternalSystemActivationPostBadRequest{}
}

/*
ExternalSystemActivationPostBadRequest describes a response with status code 400, with default header values.

ExternalSystemActivationPostBadRequest external system activation post bad request
*/
type ExternalSystemActivationPostBadRequest struct {
	Payload string
}

// IsSuccess returns true when this external system activation post bad request response has a 2xx status code
func (o *ExternalSystemActivationPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this external system activation post bad request response has a 3xx status code
func (o *ExternalSystemActivationPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this external system activation post bad request response has a 4xx status code
func (o *ExternalSystemActivationPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this external system activation post bad request response has a 5xx status code
func (o *ExternalSystemActivationPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this external system activation post bad request response a status code equal to that given
func (o *ExternalSystemActivationPostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the external system activation post bad request response
func (o *ExternalSystemActivationPostBadRequest) Code() int {
	return 400
}

func (o *ExternalSystemActivationPostBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/pel/public/externalsystemactivation][%d] externalSystemActivationPostBadRequest %s", 400, payload)
}

func (o *ExternalSystemActivationPostBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/pel/public/externalsystemactivation][%d] externalSystemActivationPostBadRequest %s", 400, payload)
}

func (o *ExternalSystemActivationPostBadRequest) GetPayload() string {
	return o.Payload
}

func (o *ExternalSystemActivationPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
