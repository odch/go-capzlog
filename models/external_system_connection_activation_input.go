// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ExternalSystemConnectionActivationInput external system connection activation input
//
// swagger:model ExternalSystemConnectionActivationInput
type ExternalSystemConnectionActivationInput struct {
	ExternalSystemConnectionActivationBase

	ExternalSystemConnectionActivationInputAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ExternalSystemConnectionActivationInput) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ExternalSystemConnectionActivationBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ExternalSystemConnectionActivationBase = aO0

	// AO1
	var aO1 ExternalSystemConnectionActivationInputAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.ExternalSystemConnectionActivationInputAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ExternalSystemConnectionActivationInput) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.ExternalSystemConnectionActivationBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.ExternalSystemConnectionActivationInputAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this external system connection activation input
func (m *ExternalSystemConnectionActivationInput) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this external system connection activation input based on context it is used
func (m *ExternalSystemConnectionActivationInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExternalSystemConnectionActivationInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExternalSystemConnectionActivationInput) UnmarshalBinary(b []byte) error {
	var res ExternalSystemConnectionActivationInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ExternalSystemConnectionActivationInputAllOf1 external system connection activation input all of1
//
// swagger:model ExternalSystemConnectionActivationInputAllOf1
type ExternalSystemConnectionActivationInputAllOf1 interface{}
