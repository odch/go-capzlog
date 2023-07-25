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

// New creates a new external system balloon flights API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for external system balloon flights API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ExternalSystemBalloonFlightsDelete(params *ExternalSystemBalloonFlightsDeleteParams, authInfo runtime.ClientAuthInfoWriter, writer io.Writer, opts ...ClientOption) (*ExternalSystemBalloonFlightsDeleteOK, error)

	ExternalSystemBalloonFlightsGet(params *ExternalSystemBalloonFlightsGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExternalSystemBalloonFlightsGetOK, error)

	ExternalSystemBalloonFlightsPost(params *ExternalSystemBalloonFlightsPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExternalSystemBalloonFlightsPostOK, error)

	ExternalSystemBalloonFlightsPostMultiple(params *ExternalSystemBalloonFlightsPostMultipleParams, authInfo runtime.ClientAuthInfoWriter, writer io.Writer, opts ...ClientOption) (*ExternalSystemBalloonFlightsPostMultipleOK, error)

	ExternalSystemBalloonFlightsPut(params *ExternalSystemBalloonFlightsPutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExternalSystemBalloonFlightsPutOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ExternalSystemBalloonFlightsDelete deletes an existing balloon flight in capzlog aero
*/
func (a *Client) ExternalSystemBalloonFlightsDelete(params *ExternalSystemBalloonFlightsDeleteParams, authInfo runtime.ClientAuthInfoWriter, writer io.Writer, opts ...ClientOption) (*ExternalSystemBalloonFlightsDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExternalSystemBalloonFlightsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ExternalSystemBalloonFlights_Delete",
		Method:             "DELETE",
		PathPattern:        "/api/pel/public/externalsystemballoonflights/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExternalSystemBalloonFlightsDeleteReader{formats: a.formats, writer: writer},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ExternalSystemBalloonFlightsDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ExternalSystemBalloonFlights_Delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExternalSystemBalloonFlightsGet gets a balloon flight from capzlog aero based on the external system unique ID
*/
func (a *Client) ExternalSystemBalloonFlightsGet(params *ExternalSystemBalloonFlightsGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExternalSystemBalloonFlightsGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExternalSystemBalloonFlightsGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ExternalSystemBalloonFlights_Get",
		Method:             "GET",
		PathPattern:        "/api/pel/public/externalsystemballoonflights/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExternalSystemBalloonFlightsGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ExternalSystemBalloonFlightsGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ExternalSystemBalloonFlights_Get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExternalSystemBalloonFlightsPost creates a balloon flight in capzlog aero
*/
func (a *Client) ExternalSystemBalloonFlightsPost(params *ExternalSystemBalloonFlightsPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExternalSystemBalloonFlightsPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExternalSystemBalloonFlightsPostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ExternalSystemBalloonFlights_Post",
		Method:             "POST",
		PathPattern:        "/api/pel/public/externalsystemballoonflights",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExternalSystemBalloonFlightsPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ExternalSystemBalloonFlightsPostOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ExternalSystemBalloonFlights_Post: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExternalSystemBalloonFlightsPostMultiple creates multiple balloon flights in capzlog aero
*/
func (a *Client) ExternalSystemBalloonFlightsPostMultiple(params *ExternalSystemBalloonFlightsPostMultipleParams, authInfo runtime.ClientAuthInfoWriter, writer io.Writer, opts ...ClientOption) (*ExternalSystemBalloonFlightsPostMultipleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExternalSystemBalloonFlightsPostMultipleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ExternalSystemBalloonFlights_PostMultiple",
		Method:             "POST",
		PathPattern:        "/api/pel/public/externalsystemballoonflights/PostMultiple",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExternalSystemBalloonFlightsPostMultipleReader{formats: a.formats, writer: writer},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ExternalSystemBalloonFlightsPostMultipleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ExternalSystemBalloonFlights_PostMultiple: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExternalSystemBalloonFlightsPut modifies an existing balloon flight in capzlog aero
*/
func (a *Client) ExternalSystemBalloonFlightsPut(params *ExternalSystemBalloonFlightsPutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExternalSystemBalloonFlightsPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExternalSystemBalloonFlightsPutParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ExternalSystemBalloonFlights_Put",
		Method:             "PUT",
		PathPattern:        "/api/pel/public/externalsystemballoonflights/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExternalSystemBalloonFlightsPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ExternalSystemBalloonFlightsPutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ExternalSystemBalloonFlights_Put: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
