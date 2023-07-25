// Code generated by go-swagger; DO NOT EDIT.

package external_system_simulator_sessions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new external system simulator sessions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for external system simulator sessions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ExternalSystemSimulatorSessionsDelete(params *ExternalSystemSimulatorSessionsDeleteParams, authInfo runtime.ClientAuthInfoWriter, writer io.Writer, opts ...ClientOption) (*ExternalSystemSimulatorSessionsDeleteOK, error)

	ExternalSystemSimulatorSessionsGet(params *ExternalSystemSimulatorSessionsGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExternalSystemSimulatorSessionsGetOK, error)

	ExternalSystemSimulatorSessionsPost(params *ExternalSystemSimulatorSessionsPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExternalSystemSimulatorSessionsPostOK, error)

	ExternalSystemSimulatorSessionsPostMultiple(params *ExternalSystemSimulatorSessionsPostMultipleParams, authInfo runtime.ClientAuthInfoWriter, writer io.Writer, opts ...ClientOption) (*ExternalSystemSimulatorSessionsPostMultipleOK, error)

	ExternalSystemSimulatorSessionsPut(params *ExternalSystemSimulatorSessionsPutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExternalSystemSimulatorSessionsPutOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ExternalSystemSimulatorSessionsDelete deletes an existing simulator session in capzlog aero
*/
func (a *Client) ExternalSystemSimulatorSessionsDelete(params *ExternalSystemSimulatorSessionsDeleteParams, authInfo runtime.ClientAuthInfoWriter, writer io.Writer, opts ...ClientOption) (*ExternalSystemSimulatorSessionsDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExternalSystemSimulatorSessionsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ExternalSystemSimulatorSessions_Delete",
		Method:             "DELETE",
		PathPattern:        "/api/pel/public/externalsystemsimulatorsessions/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExternalSystemSimulatorSessionsDeleteReader{formats: a.formats, writer: writer},
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
	success, ok := result.(*ExternalSystemSimulatorSessionsDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ExternalSystemSimulatorSessions_Delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExternalSystemSimulatorSessionsGet gets a simulator session from capzlog aero based on the external system unique ID
*/
func (a *Client) ExternalSystemSimulatorSessionsGet(params *ExternalSystemSimulatorSessionsGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExternalSystemSimulatorSessionsGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExternalSystemSimulatorSessionsGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ExternalSystemSimulatorSessions_Get",
		Method:             "GET",
		PathPattern:        "/api/pel/public/externalsystemsimulatorsessions/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExternalSystemSimulatorSessionsGetReader{formats: a.formats},
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
	success, ok := result.(*ExternalSystemSimulatorSessionsGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ExternalSystemSimulatorSessions_Get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExternalSystemSimulatorSessionsPost creates a simulator session in capzlog aero
*/
func (a *Client) ExternalSystemSimulatorSessionsPost(params *ExternalSystemSimulatorSessionsPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExternalSystemSimulatorSessionsPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExternalSystemSimulatorSessionsPostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ExternalSystemSimulatorSessions_Post",
		Method:             "POST",
		PathPattern:        "/api/pel/public/externalsystemsimulatorsessions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExternalSystemSimulatorSessionsPostReader{formats: a.formats},
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
	success, ok := result.(*ExternalSystemSimulatorSessionsPostOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ExternalSystemSimulatorSessions_Post: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExternalSystemSimulatorSessionsPostMultiple creates multiple simulator sessions in capzlog aero
*/
func (a *Client) ExternalSystemSimulatorSessionsPostMultiple(params *ExternalSystemSimulatorSessionsPostMultipleParams, authInfo runtime.ClientAuthInfoWriter, writer io.Writer, opts ...ClientOption) (*ExternalSystemSimulatorSessionsPostMultipleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExternalSystemSimulatorSessionsPostMultipleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ExternalSystemSimulatorSessions_PostMultiple",
		Method:             "POST",
		PathPattern:        "/api/pel/public/externalsystemsimulatorSessions/PostMultiple",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExternalSystemSimulatorSessionsPostMultipleReader{formats: a.formats, writer: writer},
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
	success, ok := result.(*ExternalSystemSimulatorSessionsPostMultipleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ExternalSystemSimulatorSessions_PostMultiple: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExternalSystemSimulatorSessionsPut modifies an existing simulator session in capzlog aero
*/
func (a *Client) ExternalSystemSimulatorSessionsPut(params *ExternalSystemSimulatorSessionsPutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ExternalSystemSimulatorSessionsPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExternalSystemSimulatorSessionsPutParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ExternalSystemSimulatorSessions_Put",
		Method:             "PUT",
		PathPattern:        "/api/pel/public/externalsystemsimulatorsessions/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExternalSystemSimulatorSessionsPutReader{formats: a.formats},
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
	success, ok := result.(*ExternalSystemSimulatorSessionsPutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ExternalSystemSimulatorSessions_Put: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
