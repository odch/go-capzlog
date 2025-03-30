// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SailplaneFlightInputBase Represents a sailplane flight to be created in capzlog.aero
//
// swagger:model SailplaneFlightInputBase
type SailplaneFlightInputBase struct {

	// aircraft
	Aircraft *SailplaneInput `json:"Aircraft,omitempty"`

	// are times local
	// Required: true
	AreTimesLocal bool `json:"AreTimesLocal"`

	// arrival airport
	// Required: true
	ArrivalAirport *AirportInput `json:"ArrivalAirport"`

	// If set to true, the day/night limits are calculated by capzlog.aero; no twilight values should be provided; works only with UTC times or in CET time zone. If set to false, the day/night limits should be provided as twilight values; if none are provided, the SailplaneFlight is considered day.
	// Required: true
	AutoCalculateDayNightTime bool `json:"AutoCalculateDayNightTime"`

	// date
	// Required: true
	// Min Length: 1
	// Pattern: ^\d{4}\-(0[1-9]|1[0-2])\-(0[1-9]|[1-2][0-9]|3[0-1])$
	Date *string `json:"Date"`

	// The day/night sequence of the SailplaneFlight. Required if AutoCalculateDayNightTime is false
	DayNightSequence DayNightSequences `json:"DayNightSequence,omitempty"`

	// departure airport
	// Required: true
	DepartureAirport *AirportInput `json:"DepartureAirport"`

	// Only flights in the TMG category can have night time and thus also twilight times
	// Pattern: ^([0-1][0-9]|2[0-3]):[0-5][0-9]$
	EveningTwilightTime string `json:"EveningTwilightTime,omitempty"`

	// Identifies if a powered sailplane was flown as TMG. Only flights with powered sailplanes may have this flag set to true
	// Required: true
	IsTouringMotorGlider *bool `json:"IsTouringMotorGlider"`

	// landing time
	// Required: true
	// Min Length: 1
	// Pattern: ^([0-1][0-9]|2[0-3]):[0-5][0-9]$
	LandingTime *string `json:"LandingTime"`

	// landings
	// Required: true
	// Maximum: 100
	// Minimum: 0
	Landings *int32 `json:"Landings"`

	// launch method
	// Required: true
	LaunchMethod *LaunchMethods `json:"LaunchMethod"`

	// markers
	Markers []*Marker `json:"Markers"`

	// Only flights in the TMG category can have night time and thus also twilight times
	// Pattern: ^([0-1][0-9]|2[0-3]):[0-5][0-9]$
	MorningTwilightTime string `json:"MorningTwilightTime,omitempty"`

	// night landings
	// Required: true
	// Maximum: 100
	// Minimum: 0
	NightLandings *int32 `json:"NightLandings"`

	// night takeoffs
	// Required: true
	// Maximum: 100
	// Minimum: 0
	NightTakeoffs *int32 `json:"NightTakeoffs"`

	// p i c name
	// Required: true
	// Min Length: 1
	PICName *string `json:"PICName"`

	// pilot function
	// Required: true
	PilotFunction *PilotFunctions `json:"PilotFunction"`

	// remark
	Remark string `json:"Remark,omitempty"`

	// takeoff time
	// Required: true
	// Min Length: 1
	// Pattern: ^([0-1][0-9]|2[0-3]):[0-5][0-9]$
	TakeoffTime *string `json:"TakeoffTime"`

	// takeoffs
	// Required: true
	// Maximum: 100
	// Minimum: 0
	Takeoffs *int32 `json:"Takeoffs"`
}

// Validate validates this sailplane flight input base
func (m *SailplaneFlightInputBase) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAircraft(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAreTimesLocal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateArrivalAirport(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAutoCalculateDayNightTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDayNightSequence(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDepartureAirport(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEveningTwilightTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsTouringMotorGlider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLandingTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLandings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLaunchMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMarkers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMorningTwilightTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNightLandings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNightTakeoffs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePICName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePilotFunction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTakeoffTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTakeoffs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SailplaneFlightInputBase) validateAircraft(formats strfmt.Registry) error {
	if swag.IsZero(m.Aircraft) { // not required
		return nil
	}

	if m.Aircraft != nil {
		if err := m.Aircraft.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Aircraft")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Aircraft")
			}
			return err
		}
	}

	return nil
}

func (m *SailplaneFlightInputBase) validateAreTimesLocal(formats strfmt.Registry) error {

	if err := validate.Required("AreTimesLocal", "body", bool(m.AreTimesLocal)); err != nil {
		return err
	}

	return nil
}

func (m *SailplaneFlightInputBase) validateArrivalAirport(formats strfmt.Registry) error {

	if err := validate.Required("ArrivalAirport", "body", m.ArrivalAirport); err != nil {
		return err
	}

	if m.ArrivalAirport != nil {
		if err := m.ArrivalAirport.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ArrivalAirport")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ArrivalAirport")
			}
			return err
		}
	}

	return nil
}

func (m *SailplaneFlightInputBase) validateAutoCalculateDayNightTime(formats strfmt.Registry) error {

	if err := validate.Required("AutoCalculateDayNightTime", "body", bool(m.AutoCalculateDayNightTime)); err != nil {
		return err
	}

	return nil
}

func (m *SailplaneFlightInputBase) validateDate(formats strfmt.Registry) error {

	if err := validate.Required("Date", "body", m.Date); err != nil {
		return err
	}

	if err := validate.MinLength("Date", "body", *m.Date, 1); err != nil {
		return err
	}

	if err := validate.Pattern("Date", "body", *m.Date, `^\d{4}\-(0[1-9]|1[0-2])\-(0[1-9]|[1-2][0-9]|3[0-1])$`); err != nil {
		return err
	}

	return nil
}

func (m *SailplaneFlightInputBase) validateDayNightSequence(formats strfmt.Registry) error {
	if swag.IsZero(m.DayNightSequence) { // not required
		return nil
	}

	if err := m.DayNightSequence.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("DayNightSequence")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("DayNightSequence")
		}
		return err
	}

	return nil
}

func (m *SailplaneFlightInputBase) validateDepartureAirport(formats strfmt.Registry) error {

	if err := validate.Required("DepartureAirport", "body", m.DepartureAirport); err != nil {
		return err
	}

	if m.DepartureAirport != nil {
		if err := m.DepartureAirport.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DepartureAirport")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("DepartureAirport")
			}
			return err
		}
	}

	return nil
}

func (m *SailplaneFlightInputBase) validateEveningTwilightTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EveningTwilightTime) { // not required
		return nil
	}

	if err := validate.Pattern("EveningTwilightTime", "body", m.EveningTwilightTime, `^([0-1][0-9]|2[0-3]):[0-5][0-9]$`); err != nil {
		return err
	}

	return nil
}

func (m *SailplaneFlightInputBase) validateIsTouringMotorGlider(formats strfmt.Registry) error {

	if err := validate.Required("IsTouringMotorGlider", "body", m.IsTouringMotorGlider); err != nil {
		return err
	}

	return nil
}

func (m *SailplaneFlightInputBase) validateLandingTime(formats strfmt.Registry) error {

	if err := validate.Required("LandingTime", "body", m.LandingTime); err != nil {
		return err
	}

	if err := validate.MinLength("LandingTime", "body", *m.LandingTime, 1); err != nil {
		return err
	}

	if err := validate.Pattern("LandingTime", "body", *m.LandingTime, `^([0-1][0-9]|2[0-3]):[0-5][0-9]$`); err != nil {
		return err
	}

	return nil
}

func (m *SailplaneFlightInputBase) validateLandings(formats strfmt.Registry) error {

	if err := validate.Required("Landings", "body", m.Landings); err != nil {
		return err
	}

	if err := validate.MinimumInt("Landings", "body", int64(*m.Landings), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("Landings", "body", int64(*m.Landings), 100, false); err != nil {
		return err
	}

	return nil
}

func (m *SailplaneFlightInputBase) validateLaunchMethod(formats strfmt.Registry) error {

	if err := validate.Required("LaunchMethod", "body", m.LaunchMethod); err != nil {
		return err
	}

	if err := validate.Required("LaunchMethod", "body", m.LaunchMethod); err != nil {
		return err
	}

	if m.LaunchMethod != nil {
		if err := m.LaunchMethod.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("LaunchMethod")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("LaunchMethod")
			}
			return err
		}
	}

	return nil
}

func (m *SailplaneFlightInputBase) validateMarkers(formats strfmt.Registry) error {
	if swag.IsZero(m.Markers) { // not required
		return nil
	}

	for i := 0; i < len(m.Markers); i++ {
		if swag.IsZero(m.Markers[i]) { // not required
			continue
		}

		if m.Markers[i] != nil {
			if err := m.Markers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Markers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Markers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SailplaneFlightInputBase) validateMorningTwilightTime(formats strfmt.Registry) error {
	if swag.IsZero(m.MorningTwilightTime) { // not required
		return nil
	}

	if err := validate.Pattern("MorningTwilightTime", "body", m.MorningTwilightTime, `^([0-1][0-9]|2[0-3]):[0-5][0-9]$`); err != nil {
		return err
	}

	return nil
}

func (m *SailplaneFlightInputBase) validateNightLandings(formats strfmt.Registry) error {

	if err := validate.Required("NightLandings", "body", m.NightLandings); err != nil {
		return err
	}

	if err := validate.MinimumInt("NightLandings", "body", int64(*m.NightLandings), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("NightLandings", "body", int64(*m.NightLandings), 100, false); err != nil {
		return err
	}

	return nil
}

func (m *SailplaneFlightInputBase) validateNightTakeoffs(formats strfmt.Registry) error {

	if err := validate.Required("NightTakeoffs", "body", m.NightTakeoffs); err != nil {
		return err
	}

	if err := validate.MinimumInt("NightTakeoffs", "body", int64(*m.NightTakeoffs), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("NightTakeoffs", "body", int64(*m.NightTakeoffs), 100, false); err != nil {
		return err
	}

	return nil
}

func (m *SailplaneFlightInputBase) validatePICName(formats strfmt.Registry) error {

	if err := validate.Required("PICName", "body", m.PICName); err != nil {
		return err
	}

	if err := validate.MinLength("PICName", "body", *m.PICName, 1); err != nil {
		return err
	}

	return nil
}

func (m *SailplaneFlightInputBase) validatePilotFunction(formats strfmt.Registry) error {

	if err := validate.Required("PilotFunction", "body", m.PilotFunction); err != nil {
		return err
	}

	if err := validate.Required("PilotFunction", "body", m.PilotFunction); err != nil {
		return err
	}

	if m.PilotFunction != nil {
		if err := m.PilotFunction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PilotFunction")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PilotFunction")
			}
			return err
		}
	}

	return nil
}

func (m *SailplaneFlightInputBase) validateTakeoffTime(formats strfmt.Registry) error {

	if err := validate.Required("TakeoffTime", "body", m.TakeoffTime); err != nil {
		return err
	}

	if err := validate.MinLength("TakeoffTime", "body", *m.TakeoffTime, 1); err != nil {
		return err
	}

	if err := validate.Pattern("TakeoffTime", "body", *m.TakeoffTime, `^([0-1][0-9]|2[0-3]):[0-5][0-9]$`); err != nil {
		return err
	}

	return nil
}

func (m *SailplaneFlightInputBase) validateTakeoffs(formats strfmt.Registry) error {

	if err := validate.Required("Takeoffs", "body", m.Takeoffs); err != nil {
		return err
	}

	if err := validate.MinimumInt("Takeoffs", "body", int64(*m.Takeoffs), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("Takeoffs", "body", int64(*m.Takeoffs), 100, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this sailplane flight input base based on the context it is used
func (m *SailplaneFlightInputBase) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAircraft(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateArrivalAirport(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDayNightSequence(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDepartureAirport(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLaunchMethod(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMarkers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePilotFunction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SailplaneFlightInputBase) contextValidateAircraft(ctx context.Context, formats strfmt.Registry) error {

	if m.Aircraft != nil {

		if swag.IsZero(m.Aircraft) { // not required
			return nil
		}

		if err := m.Aircraft.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Aircraft")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Aircraft")
			}
			return err
		}
	}

	return nil
}

func (m *SailplaneFlightInputBase) contextValidateArrivalAirport(ctx context.Context, formats strfmt.Registry) error {

	if m.ArrivalAirport != nil {

		if err := m.ArrivalAirport.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ArrivalAirport")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ArrivalAirport")
			}
			return err
		}
	}

	return nil
}

func (m *SailplaneFlightInputBase) contextValidateDayNightSequence(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.DayNightSequence) { // not required
		return nil
	}

	if err := m.DayNightSequence.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("DayNightSequence")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("DayNightSequence")
		}
		return err
	}

	return nil
}

func (m *SailplaneFlightInputBase) contextValidateDepartureAirport(ctx context.Context, formats strfmt.Registry) error {

	if m.DepartureAirport != nil {

		if err := m.DepartureAirport.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DepartureAirport")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("DepartureAirport")
			}
			return err
		}
	}

	return nil
}

func (m *SailplaneFlightInputBase) contextValidateLaunchMethod(ctx context.Context, formats strfmt.Registry) error {

	if m.LaunchMethod != nil {

		if err := m.LaunchMethod.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("LaunchMethod")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("LaunchMethod")
			}
			return err
		}
	}

	return nil
}

func (m *SailplaneFlightInputBase) contextValidateMarkers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Markers); i++ {

		if m.Markers[i] != nil {

			if swag.IsZero(m.Markers[i]) { // not required
				return nil
			}

			if err := m.Markers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Markers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Markers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SailplaneFlightInputBase) contextValidatePilotFunction(ctx context.Context, formats strfmt.Registry) error {

	if m.PilotFunction != nil {

		if err := m.PilotFunction.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PilotFunction")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PilotFunction")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SailplaneFlightInputBase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SailplaneFlightInputBase) UnmarshalBinary(b []byte) error {
	var res SailplaneFlightInputBase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
