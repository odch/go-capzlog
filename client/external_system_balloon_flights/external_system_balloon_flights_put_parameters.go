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

// NewExternalSystemBalloonFlightsPutParams creates a new ExternalSystemBalloonFlightsPutParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExternalSystemBalloonFlightsPutParams() *ExternalSystemBalloonFlightsPutParams {
	return &ExternalSystemBalloonFlightsPutParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExternalSystemBalloonFlightsPutParamsWithTimeout creates a new ExternalSystemBalloonFlightsPutParams object
// with the ability to set a timeout on a request.
func NewExternalSystemBalloonFlightsPutParamsWithTimeout(timeout time.Duration) *ExternalSystemBalloonFlightsPutParams {
	return &ExternalSystemBalloonFlightsPutParams{
		timeout: timeout,
	}
}

// NewExternalSystemBalloonFlightsPutParamsWithContext creates a new ExternalSystemBalloonFlightsPutParams object
// with the ability to set a context for a request.
func NewExternalSystemBalloonFlightsPutParamsWithContext(ctx context.Context) *ExternalSystemBalloonFlightsPutParams {
	return &ExternalSystemBalloonFlightsPutParams{
		Context: ctx,
	}
}

// NewExternalSystemBalloonFlightsPutParamsWithHTTPClient creates a new ExternalSystemBalloonFlightsPutParams object
// with the ability to set a custom HTTPClient for a request.
func NewExternalSystemBalloonFlightsPutParamsWithHTTPClient(client *http.Client) *ExternalSystemBalloonFlightsPutParams {
	return &ExternalSystemBalloonFlightsPutParams{
		HTTPClient: client,
	}
}

/* ExternalSystemBalloonFlightsPutParams contains all the parameters to send to the API endpoint
   for the external system balloon flights put operation.

   Typically these are written to a http.Request.
*/
type ExternalSystemBalloonFlightsPutParams struct {

	/* ID.

	   The ExternalSystemUniqueID of the balloonFlight to be modified
	*/
	ID string

	/* InputEntity.

	   Full data of the balloonFlight to be edited
	*/
	InputEntity *models.ExternalSystemBalloonFlightInput

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the external system balloon flights put params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExternalSystemBalloonFlightsPutParams) WithDefaults() *ExternalSystemBalloonFlightsPutParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the external system balloon flights put params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExternalSystemBalloonFlightsPutParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the external system balloon flights put params
func (o *ExternalSystemBalloonFlightsPutParams) WithTimeout(timeout time.Duration) *ExternalSystemBalloonFlightsPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the external system balloon flights put params
func (o *ExternalSystemBalloonFlightsPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the external system balloon flights put params
func (o *ExternalSystemBalloonFlightsPutParams) WithContext(ctx context.Context) *ExternalSystemBalloonFlightsPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the external system balloon flights put params
func (o *ExternalSystemBalloonFlightsPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the external system balloon flights put params
func (o *ExternalSystemBalloonFlightsPutParams) WithHTTPClient(client *http.Client) *ExternalSystemBalloonFlightsPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the external system balloon flights put params
func (o *ExternalSystemBalloonFlightsPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the external system balloon flights put params
func (o *ExternalSystemBalloonFlightsPutParams) WithID(id string) *ExternalSystemBalloonFlightsPutParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the external system balloon flights put params
func (o *ExternalSystemBalloonFlightsPutParams) SetID(id string) {
	o.ID = id
}

// WithInputEntity adds the inputEntity to the external system balloon flights put params
func (o *ExternalSystemBalloonFlightsPutParams) WithInputEntity(inputEntity *models.ExternalSystemBalloonFlightInput) *ExternalSystemBalloonFlightsPutParams {
	o.SetInputEntity(inputEntity)
	return o
}

// SetInputEntity adds the inputEntity to the external system balloon flights put params
func (o *ExternalSystemBalloonFlightsPutParams) SetInputEntity(inputEntity *models.ExternalSystemBalloonFlightInput) {
	o.InputEntity = inputEntity
}

// WriteToRequest writes these params to a swagger request
func (o *ExternalSystemBalloonFlightsPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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