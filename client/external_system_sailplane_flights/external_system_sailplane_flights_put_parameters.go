// Code generated by go-swagger; DO NOT EDIT.

package external_system_sailplane_flights

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

// NewExternalSystemSailplaneFlightsPutParams creates a new ExternalSystemSailplaneFlightsPutParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExternalSystemSailplaneFlightsPutParams() *ExternalSystemSailplaneFlightsPutParams {
	return &ExternalSystemSailplaneFlightsPutParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExternalSystemSailplaneFlightsPutParamsWithTimeout creates a new ExternalSystemSailplaneFlightsPutParams object
// with the ability to set a timeout on a request.
func NewExternalSystemSailplaneFlightsPutParamsWithTimeout(timeout time.Duration) *ExternalSystemSailplaneFlightsPutParams {
	return &ExternalSystemSailplaneFlightsPutParams{
		timeout: timeout,
	}
}

// NewExternalSystemSailplaneFlightsPutParamsWithContext creates a new ExternalSystemSailplaneFlightsPutParams object
// with the ability to set a context for a request.
func NewExternalSystemSailplaneFlightsPutParamsWithContext(ctx context.Context) *ExternalSystemSailplaneFlightsPutParams {
	return &ExternalSystemSailplaneFlightsPutParams{
		Context: ctx,
	}
}

// NewExternalSystemSailplaneFlightsPutParamsWithHTTPClient creates a new ExternalSystemSailplaneFlightsPutParams object
// with the ability to set a custom HTTPClient for a request.
func NewExternalSystemSailplaneFlightsPutParamsWithHTTPClient(client *http.Client) *ExternalSystemSailplaneFlightsPutParams {
	return &ExternalSystemSailplaneFlightsPutParams{
		HTTPClient: client,
	}
}

/* ExternalSystemSailplaneFlightsPutParams contains all the parameters to send to the API endpoint
   for the external system sailplane flights put operation.

   Typically these are written to a http.Request.
*/
type ExternalSystemSailplaneFlightsPutParams struct {

	/* ID.

	   The ExternalSystemUniqueID of the sailplaneFlight to be modified
	*/
	ID string

	/* InputEntity.

	   Full data of the sailplaneFlight to be edited
	*/
	InputEntity *models.ExternalSystemSailplaneFlightInput

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the external system sailplane flights put params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExternalSystemSailplaneFlightsPutParams) WithDefaults() *ExternalSystemSailplaneFlightsPutParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the external system sailplane flights put params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExternalSystemSailplaneFlightsPutParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the external system sailplane flights put params
func (o *ExternalSystemSailplaneFlightsPutParams) WithTimeout(timeout time.Duration) *ExternalSystemSailplaneFlightsPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the external system sailplane flights put params
func (o *ExternalSystemSailplaneFlightsPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the external system sailplane flights put params
func (o *ExternalSystemSailplaneFlightsPutParams) WithContext(ctx context.Context) *ExternalSystemSailplaneFlightsPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the external system sailplane flights put params
func (o *ExternalSystemSailplaneFlightsPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the external system sailplane flights put params
func (o *ExternalSystemSailplaneFlightsPutParams) WithHTTPClient(client *http.Client) *ExternalSystemSailplaneFlightsPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the external system sailplane flights put params
func (o *ExternalSystemSailplaneFlightsPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the external system sailplane flights put params
func (o *ExternalSystemSailplaneFlightsPutParams) WithID(id string) *ExternalSystemSailplaneFlightsPutParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the external system sailplane flights put params
func (o *ExternalSystemSailplaneFlightsPutParams) SetID(id string) {
	o.ID = id
}

// WithInputEntity adds the inputEntity to the external system sailplane flights put params
func (o *ExternalSystemSailplaneFlightsPutParams) WithInputEntity(inputEntity *models.ExternalSystemSailplaneFlightInput) *ExternalSystemSailplaneFlightsPutParams {
	o.SetInputEntity(inputEntity)
	return o
}

// SetInputEntity adds the inputEntity to the external system sailplane flights put params
func (o *ExternalSystemSailplaneFlightsPutParams) SetInputEntity(inputEntity *models.ExternalSystemSailplaneFlightInput) {
	o.InputEntity = inputEntity
}

// WriteToRequest writes these params to a swagger request
func (o *ExternalSystemSailplaneFlightsPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
