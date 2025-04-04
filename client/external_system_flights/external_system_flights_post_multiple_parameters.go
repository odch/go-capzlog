// Code generated by go-swagger; DO NOT EDIT.

package external_system_flights

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/odch/go-capzlog/models"
)

// NewExternalSystemFlightsPostMultipleParams creates a new ExternalSystemFlightsPostMultipleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExternalSystemFlightsPostMultipleParams() *ExternalSystemFlightsPostMultipleParams {
	return &ExternalSystemFlightsPostMultipleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExternalSystemFlightsPostMultipleParamsWithTimeout creates a new ExternalSystemFlightsPostMultipleParams object
// with the ability to set a timeout on a request.
func NewExternalSystemFlightsPostMultipleParamsWithTimeout(timeout time.Duration) *ExternalSystemFlightsPostMultipleParams {
	return &ExternalSystemFlightsPostMultipleParams{
		timeout: timeout,
	}
}

// NewExternalSystemFlightsPostMultipleParamsWithContext creates a new ExternalSystemFlightsPostMultipleParams object
// with the ability to set a context for a request.
func NewExternalSystemFlightsPostMultipleParamsWithContext(ctx context.Context) *ExternalSystemFlightsPostMultipleParams {
	return &ExternalSystemFlightsPostMultipleParams{
		Context: ctx,
	}
}

// NewExternalSystemFlightsPostMultipleParamsWithHTTPClient creates a new ExternalSystemFlightsPostMultipleParams object
// with the ability to set a custom HTTPClient for a request.
func NewExternalSystemFlightsPostMultipleParamsWithHTTPClient(client *http.Client) *ExternalSystemFlightsPostMultipleParams {
	return &ExternalSystemFlightsPostMultipleParams{
		HTTPClient: client,
	}
}

/*
ExternalSystemFlightsPostMultipleParams contains all the parameters to send to the API endpoint

	for the external system flights post multiple operation.

	Typically these are written to a http.Request.
*/
type ExternalSystemFlightsPostMultipleParams struct {

	/* InputEntities.

	   Flights to be created
	*/
	InputEntities []*models.ExternalSystemFlightInput

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the external system flights post multiple params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExternalSystemFlightsPostMultipleParams) WithDefaults() *ExternalSystemFlightsPostMultipleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the external system flights post multiple params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExternalSystemFlightsPostMultipleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the external system flights post multiple params
func (o *ExternalSystemFlightsPostMultipleParams) WithTimeout(timeout time.Duration) *ExternalSystemFlightsPostMultipleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the external system flights post multiple params
func (o *ExternalSystemFlightsPostMultipleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the external system flights post multiple params
func (o *ExternalSystemFlightsPostMultipleParams) WithContext(ctx context.Context) *ExternalSystemFlightsPostMultipleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the external system flights post multiple params
func (o *ExternalSystemFlightsPostMultipleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the external system flights post multiple params
func (o *ExternalSystemFlightsPostMultipleParams) WithHTTPClient(client *http.Client) *ExternalSystemFlightsPostMultipleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the external system flights post multiple params
func (o *ExternalSystemFlightsPostMultipleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInputEntities adds the inputEntities to the external system flights post multiple params
func (o *ExternalSystemFlightsPostMultipleParams) WithInputEntities(inputEntities []*models.ExternalSystemFlightInput) *ExternalSystemFlightsPostMultipleParams {
	o.SetInputEntities(inputEntities)
	return o
}

// SetInputEntities adds the inputEntities to the external system flights post multiple params
func (o *ExternalSystemFlightsPostMultipleParams) SetInputEntities(inputEntities []*models.ExternalSystemFlightInput) {
	o.InputEntities = inputEntities
}

// WriteToRequest writes these params to a swagger request
func (o *ExternalSystemFlightsPostMultipleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.InputEntities != nil {
		if err := r.SetBodyParam(o.InputEntities); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
