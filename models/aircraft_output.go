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

// AircraftOutput aircraft output
//
// swagger:model AircraftOutput
type AircraftOutput struct {

	// i c a o designator
	// Required: true
	// Min Length: 1
	ICAODesignator *string `json:"ICAODesignator"`

	// registration
	// Required: true
	// Min Length: 1
	Registration *string `json:"Registration"`

	// The variant of the aircraft, in case several variants are available for the provided registration in capzlog.aero. If no variant is specified and there are multiple variants for the provided registration in capzlog.aero, one is selected randomly
	Variant string `json:"Variant,omitempty"`
}

// Validate validates this aircraft output
func (m *AircraftOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateICAODesignator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegistration(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AircraftOutput) validateICAODesignator(formats strfmt.Registry) error {

	if err := validate.Required("ICAODesignator", "body", m.ICAODesignator); err != nil {
		return err
	}

	if err := validate.MinLength("ICAODesignator", "body", *m.ICAODesignator, 1); err != nil {
		return err
	}

	return nil
}

func (m *AircraftOutput) validateRegistration(formats strfmt.Registry) error {

	if err := validate.Required("Registration", "body", m.Registration); err != nil {
		return err
	}

	if err := validate.MinLength("Registration", "body", *m.Registration, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this aircraft output based on context it is used
func (m *AircraftOutput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AircraftOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AircraftOutput) UnmarshalBinary(b []byte) error {
	var res AircraftOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
