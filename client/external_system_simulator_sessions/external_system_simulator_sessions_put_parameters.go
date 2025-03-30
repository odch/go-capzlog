// Code generated by go-swagger; DO NOT EDIT.

package external_system_simulator_sessions

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

// NewExternalSystemSimulatorSessionsPutParams creates a new ExternalSystemSimulatorSessionsPutParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExternalSystemSimulatorSessionsPutParams() *ExternalSystemSimulatorSessionsPutParams {
	return &ExternalSystemSimulatorSessionsPutParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExternalSystemSimulatorSessionsPutParamsWithTimeout creates a new ExternalSystemSimulatorSessionsPutParams object
// with the ability to set a timeout on a request.
func NewExternalSystemSimulatorSessionsPutParamsWithTimeout(timeout time.Duration) *ExternalSystemSimulatorSessionsPutParams {
	return &ExternalSystemSimulatorSessionsPutParams{
		timeout: timeout,
	}
}

// NewExternalSystemSimulatorSessionsPutParamsWithContext creates a new ExternalSystemSimulatorSessionsPutParams object
// with the ability to set a context for a request.
func NewExternalSystemSimulatorSessionsPutParamsWithContext(ctx context.Context) *ExternalSystemSimulatorSessionsPutParams {
	return &ExternalSystemSimulatorSessionsPutParams{
		Context: ctx,
	}
}

// NewExternalSystemSimulatorSessionsPutParamsWithHTTPClient creates a new ExternalSystemSimulatorSessionsPutParams object
// with the ability to set a custom HTTPClient for a request.
func NewExternalSystemSimulatorSessionsPutParamsWithHTTPClient(client *http.Client) *ExternalSystemSimulatorSessionsPutParams {
	return &ExternalSystemSimulatorSessionsPutParams{
		HTTPClient: client,
	}
}

/*
ExternalSystemSimulatorSessionsPutParams contains all the parameters to send to the API endpoint

	for the external system simulator sessions put operation.

	Typically these are written to a http.Request.
*/
type ExternalSystemSimulatorSessionsPutParams struct {

	/* ID.

	   The ExternalSystemUniqueID of the simulatorSession to be modified
	*/
	ID string

	/* InputEntity.

	   Full data of the simulatorSession to be edited
	*/
	InputEntity *models.ExternalSystemSimulatorSessionInput

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the external system simulator sessions put params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExternalSystemSimulatorSessionsPutParams) WithDefaults() *ExternalSystemSimulatorSessionsPutParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the external system simulator sessions put params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExternalSystemSimulatorSessionsPutParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the external system simulator sessions put params
func (o *ExternalSystemSimulatorSessionsPutParams) WithTimeout(timeout time.Duration) *ExternalSystemSimulatorSessionsPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the external system simulator sessions put params
func (o *ExternalSystemSimulatorSessionsPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the external system simulator sessions put params
func (o *ExternalSystemSimulatorSessionsPutParams) WithContext(ctx context.Context) *ExternalSystemSimulatorSessionsPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the external system simulator sessions put params
func (o *ExternalSystemSimulatorSessionsPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the external system simulator sessions put params
func (o *ExternalSystemSimulatorSessionsPutParams) WithHTTPClient(client *http.Client) *ExternalSystemSimulatorSessionsPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the external system simulator sessions put params
func (o *ExternalSystemSimulatorSessionsPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the external system simulator sessions put params
func (o *ExternalSystemSimulatorSessionsPutParams) WithID(id string) *ExternalSystemSimulatorSessionsPutParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the external system simulator sessions put params
func (o *ExternalSystemSimulatorSessionsPutParams) SetID(id string) {
	o.ID = id
}

// WithInputEntity adds the inputEntity to the external system simulator sessions put params
func (o *ExternalSystemSimulatorSessionsPutParams) WithInputEntity(inputEntity *models.ExternalSystemSimulatorSessionInput) *ExternalSystemSimulatorSessionsPutParams {
	o.SetInputEntity(inputEntity)
	return o
}

// SetInputEntity adds the inputEntity to the external system simulator sessions put params
func (o *ExternalSystemSimulatorSessionsPutParams) SetInputEntity(inputEntity *models.ExternalSystemSimulatorSessionInput) {
	o.InputEntity = inputEntity
}

// WriteToRequest writes these params to a swagger request
func (o *ExternalSystemSimulatorSessionsPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}
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
