// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// FlightSimulationTrainingDeviceTypes flight simulation training device types
//
// swagger:model FlightSimulationTrainingDeviceTypes
type FlightSimulationTrainingDeviceTypes string

func NewFlightSimulationTrainingDeviceTypes(value FlightSimulationTrainingDeviceTypes) *FlightSimulationTrainingDeviceTypes {
	return &value
}

// Pointer returns a pointer to a freshly-allocated FlightSimulationTrainingDeviceTypes.
func (m FlightSimulationTrainingDeviceTypes) Pointer() *FlightSimulationTrainingDeviceTypes {
	return &m
}

const (

	// FlightSimulationTrainingDeviceTypesFFSA captures enum value "FFS_A"
	FlightSimulationTrainingDeviceTypesFFSA FlightSimulationTrainingDeviceTypes = "FFS_A"

	// FlightSimulationTrainingDeviceTypesFFSB captures enum value "FFS_B"
	FlightSimulationTrainingDeviceTypesFFSB FlightSimulationTrainingDeviceTypes = "FFS_B"

	// FlightSimulationTrainingDeviceTypesFFSC captures enum value "FFS_C"
	FlightSimulationTrainingDeviceTypesFFSC FlightSimulationTrainingDeviceTypes = "FFS_C"

	// FlightSimulationTrainingDeviceTypesFFSD captures enum value "FFS_D"
	FlightSimulationTrainingDeviceTypesFFSD FlightSimulationTrainingDeviceTypes = "FFS_D"

	// FlightSimulationTrainingDeviceTypesFTD1 captures enum value "FTD_1"
	FlightSimulationTrainingDeviceTypesFTD1 FlightSimulationTrainingDeviceTypes = "FTD_1"

	// FlightSimulationTrainingDeviceTypesFTD2 captures enum value "FTD_2"
	FlightSimulationTrainingDeviceTypesFTD2 FlightSimulationTrainingDeviceTypes = "FTD_2"

	// FlightSimulationTrainingDeviceTypesFTD3 captures enum value "FTD_3"
	FlightSimulationTrainingDeviceTypesFTD3 FlightSimulationTrainingDeviceTypes = "FTD_3"

	// FlightSimulationTrainingDeviceTypesFNPTI captures enum value "FNPT_I"
	FlightSimulationTrainingDeviceTypesFNPTI FlightSimulationTrainingDeviceTypes = "FNPT_I"

	// FlightSimulationTrainingDeviceTypesFNPTII captures enum value "FNPT_II"
	FlightSimulationTrainingDeviceTypesFNPTII FlightSimulationTrainingDeviceTypes = "FNPT_II"

	// FlightSimulationTrainingDeviceTypesFNPTIII captures enum value "FNPT_III"
	FlightSimulationTrainingDeviceTypesFNPTIII FlightSimulationTrainingDeviceTypes = "FNPT_III"

	// FlightSimulationTrainingDeviceTypesOTD captures enum value "OTD"
	FlightSimulationTrainingDeviceTypesOTD FlightSimulationTrainingDeviceTypes = "OTD"

	// FlightSimulationTrainingDeviceTypesFS captures enum value "FS"
	FlightSimulationTrainingDeviceTypesFS FlightSimulationTrainingDeviceTypes = "FS"

	// FlightSimulationTrainingDeviceTypesBITD captures enum value "BITD"
	FlightSimulationTrainingDeviceTypesBITD FlightSimulationTrainingDeviceTypes = "BITD"

	// FlightSimulationTrainingDeviceTypesBATD captures enum value "BATD"
	FlightSimulationTrainingDeviceTypesBATD FlightSimulationTrainingDeviceTypes = "BATD"

	// FlightSimulationTrainingDeviceTypesAATD captures enum value "AATD"
	FlightSimulationTrainingDeviceTypesAATD FlightSimulationTrainingDeviceTypes = "AATD"

	// FlightSimulationTrainingDeviceTypesInitialValues captures enum value "InitialValues"
	FlightSimulationTrainingDeviceTypesInitialValues FlightSimulationTrainingDeviceTypes = "InitialValues"

	// FlightSimulationTrainingDeviceTypesUncertified captures enum value "Uncertified"
	FlightSimulationTrainingDeviceTypesUncertified FlightSimulationTrainingDeviceTypes = "Uncertified"
)

// for schema
var flightSimulationTrainingDeviceTypesEnum []interface{}

func init() {
	var res []FlightSimulationTrainingDeviceTypes
	if err := json.Unmarshal([]byte(`["FFS_A","FFS_B","FFS_C","FFS_D","FTD_1","FTD_2","FTD_3","FNPT_I","FNPT_II","FNPT_III","OTD","FS","BITD","BATD","AATD","InitialValues","Uncertified"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		flightSimulationTrainingDeviceTypesEnum = append(flightSimulationTrainingDeviceTypesEnum, v)
	}
}

func (m FlightSimulationTrainingDeviceTypes) validateFlightSimulationTrainingDeviceTypesEnum(path, location string, value FlightSimulationTrainingDeviceTypes) error {
	if err := validate.EnumCase(path, location, value, flightSimulationTrainingDeviceTypesEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this flight simulation training device types
func (m FlightSimulationTrainingDeviceTypes) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateFlightSimulationTrainingDeviceTypesEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this flight simulation training device types based on context it is used
func (m FlightSimulationTrainingDeviceTypes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
