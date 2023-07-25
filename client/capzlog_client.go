// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/odch/go-capzlog/client/external_system_activation"
	"github.com/odch/go-capzlog/client/external_system_balloon_flights"
	"github.com/odch/go-capzlog/client/external_system_flights"
	"github.com/odch/go-capzlog/client/external_system_sailplane_flights"
	"github.com/odch/go-capzlog/client/external_system_simulator_sessions"
)

// Default capzlog HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "test.capzlog.aero"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new capzlog HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Capzlog {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new capzlog HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Capzlog {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new capzlog client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Capzlog {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Capzlog)
	cli.Transport = transport
	cli.ExternalSystemActivation = external_system_activation.New(transport, formats)
	cli.ExternalSystemBalloonFlights = external_system_balloon_flights.New(transport, formats)
	cli.ExternalSystemFlights = external_system_flights.New(transport, formats)
	cli.ExternalSystemSailplaneFlights = external_system_sailplane_flights.New(transport, formats)
	cli.ExternalSystemSimulatorSessions = external_system_simulator_sessions.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Capzlog is a client for capzlog
type Capzlog struct {
	ExternalSystemActivation external_system_activation.ClientService

	ExternalSystemBalloonFlights external_system_balloon_flights.ClientService

	ExternalSystemFlights external_system_flights.ClientService

	ExternalSystemSailplaneFlights external_system_sailplane_flights.ClientService

	ExternalSystemSimulatorSessions external_system_simulator_sessions.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Capzlog) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.ExternalSystemActivation.SetTransport(transport)
	c.ExternalSystemBalloonFlights.SetTransport(transport)
	c.ExternalSystemFlights.SetTransport(transport)
	c.ExternalSystemSailplaneFlights.SetTransport(transport)
	c.ExternalSystemSimulatorSessions.SetTransport(transport)
}