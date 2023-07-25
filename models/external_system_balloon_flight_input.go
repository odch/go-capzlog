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

// ExternalSystemBalloonFlightInput external system balloon flight input
//
// swagger:model ExternalSystemBalloonFlightInput
type ExternalSystemBalloonFlightInput struct {
	BalloonFlightInputBase

	// external system unique ID
	// Required: true
	// Min Length: 1
	// Pattern: ^\w{1-100}$
	ExternalSystemUniqueID *string `json:"ExternalSystemUniqueID"`

	// is last flight of day
	// Required: true
	IsLastFlightOfDay *bool `json:"IsLastFlightOfDay"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ExternalSystemBalloonFlightInput) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BalloonFlightInputBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BalloonFlightInputBase = aO0

	// AO1
	var dataAO1 struct {
		ExternalSystemUniqueID *string `json:"ExternalSystemUniqueID"`

		IsLastFlightOfDay *bool `json:"IsLastFlightOfDay"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ExternalSystemUniqueID = dataAO1.ExternalSystemUniqueID

	m.IsLastFlightOfDay = dataAO1.IsLastFlightOfDay

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ExternalSystemBalloonFlightInput) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.BalloonFlightInputBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		ExternalSystemUniqueID *string `json:"ExternalSystemUniqueID"`

		IsLastFlightOfDay *bool `json:"IsLastFlightOfDay"`
	}

	dataAO1.ExternalSystemUniqueID = m.ExternalSystemUniqueID

	dataAO1.IsLastFlightOfDay = m.IsLastFlightOfDay

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this external system balloon flight input
func (m *ExternalSystemBalloonFlightInput) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BalloonFlightInputBase
	if err := m.BalloonFlightInputBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalSystemUniqueID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsLastFlightOfDay(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExternalSystemBalloonFlightInput) validateExternalSystemUniqueID(formats strfmt.Registry) error {

	if err := validate.Required("ExternalSystemUniqueID", "body", m.ExternalSystemUniqueID); err != nil {
		return err
	}

	if err := validate.MinLength("ExternalSystemUniqueID", "body", *m.ExternalSystemUniqueID, 1); err != nil {
		return err
	}

	if err := validate.Pattern("ExternalSystemUniqueID", "body", *m.ExternalSystemUniqueID, `^\w{1-100}$`); err != nil {
		return err
	}

	return nil
}

func (m *ExternalSystemBalloonFlightInput) validateIsLastFlightOfDay(formats strfmt.Registry) error {

	if err := validate.Required("IsLastFlightOfDay", "body", m.IsLastFlightOfDay); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this external system balloon flight input based on the context it is used
func (m *ExternalSystemBalloonFlightInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BalloonFlightInputBase
	if err := m.BalloonFlightInputBase.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ExternalSystemBalloonFlightInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExternalSystemBalloonFlightInput) UnmarshalBinary(b []byte) error {
	var res ExternalSystemBalloonFlightInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
