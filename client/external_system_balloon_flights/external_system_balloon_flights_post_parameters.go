// Code generated by go-swagger; DO NOT EDIT.

package external_system_balloon_flights

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

// NewExternalSystemBalloonFlightsPostParams creates a new ExternalSystemBalloonFlightsPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExternalSystemBalloonFlightsPostParams() *ExternalSystemBalloonFlightsPostParams {
	return &ExternalSystemBalloonFlightsPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExternalSystemBalloonFlightsPostParamsWithTimeout creates a new ExternalSystemBalloonFlightsPostParams object
// with the ability to set a timeout on a request.
func NewExternalSystemBalloonFlightsPostParamsWithTimeout(timeout time.Duration) *ExternalSystemBalloonFlightsPostParams {
	return &ExternalSystemBalloonFlightsPostParams{
		timeout: timeout,
	}
}

// NewExternalSystemBalloonFlightsPostParamsWithContext creates a new ExternalSystemBalloonFlightsPostParams object
// with the ability to set a context for a request.
func NewExternalSystemBalloonFlightsPostParamsWithContext(ctx context.Context) *ExternalSystemBalloonFlightsPostParams {
	return &ExternalSystemBalloonFlightsPostParams{
		Context: ctx,
	}
}

// NewExternalSystemBalloonFlightsPostParamsWithHTTPClient creates a new ExternalSystemBalloonFlightsPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewExternalSystemBalloonFlightsPostParamsWithHTTPClient(client *http.Client) *ExternalSystemBalloonFlightsPostParams {
	return &ExternalSystemBalloonFlightsPostParams{
		HTTPClient: client,
	}
}

/*
ExternalSystemBalloonFlightsPostParams contains all the parameters to send to the API endpoint

	for the external system balloon flights post operation.

	Typically these are written to a http.Request.
*/
type ExternalSystemBalloonFlightsPostParams struct {

	/* InputEntity.

	   BalloonFlight to be created
	*/
	InputEntity *models.ExternalSystemBalloonFlightInput

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the external system balloon flights post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExternalSystemBalloonFlightsPostParams) WithDefaults() *ExternalSystemBalloonFlightsPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the external system balloon flights post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExternalSystemBalloonFlightsPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the external system balloon flights post params
func (o *ExternalSystemBalloonFlightsPostParams) WithTimeout(timeout time.Duration) *ExternalSystemBalloonFlightsPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the external system balloon flights post params
func (o *ExternalSystemBalloonFlightsPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the external system balloon flights post params
func (o *ExternalSystemBalloonFlightsPostParams) WithContext(ctx context.Context) *ExternalSystemBalloonFlightsPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the external system balloon flights post params
func (o *ExternalSystemBalloonFlightsPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the external system balloon flights post params
func (o *ExternalSystemBalloonFlightsPostParams) WithHTTPClient(client *http.Client) *ExternalSystemBalloonFlightsPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the external system balloon flights post params
func (o *ExternalSystemBalloonFlightsPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInputEntity adds the inputEntity to the external system balloon flights post params
func (o *ExternalSystemBalloonFlightsPostParams) WithInputEntity(inputEntity *models.ExternalSystemBalloonFlightInput) *ExternalSystemBalloonFlightsPostParams {
	o.SetInputEntity(inputEntity)
	return o
}

// SetInputEntity adds the inputEntity to the external system balloon flights post params
func (o *ExternalSystemBalloonFlightsPostParams) SetInputEntity(inputEntity *models.ExternalSystemBalloonFlightInput) {
	o.InputEntity = inputEntity
}

// WriteToRequest writes these params to a swagger request
func (o *ExternalSystemBalloonFlightsPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.InputEntity != nil {
		if err := r.SetBodyParam(o.InputEntity); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
