// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AirportOutput airport output
//
// swagger:model AirportOutput
type AirportOutput struct {

	// A free text name of the airport or landing location. Shall only be used if the ICAO code is a no location indicator ZZZZ or abXX where ab is the country prefix, e.g. LSXX for no location indicator in Switzerland
	FreeText string `json:"FreeText,omitempty"`

	// i c a o code
	// Required: true
	// Min Length: 1
	ICAOCode *string `json:"ICAOCode"`
}

// Validate validates this airport output
func (m *AirportOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateICAOCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AirportOutput) validateICAOCode(formats strfmt.Registry) error {

	if err := validate.Required("ICAOCode", "body", m.ICAOCode); err != nil {
		return err
	}

	if err := validate.MinLength("ICAOCode", "body", *m.ICAOCode, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this airport output based on context it is used
func (m *AirportOutput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AirportOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AirportOutput) UnmarshalBinary(b []byte) error {
	var res AirportOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
