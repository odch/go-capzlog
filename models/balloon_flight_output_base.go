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

// BalloonFlightOutputBase Represents a balloon flight in capzlog.aero
//
// swagger:model BalloonFlightOutputBase
type BalloonFlightOutputBase struct {

	// aircraft
	Aircraft *BalloonOutput `json:"Aircraft,omitempty"`

	// are times local
	// Required: true
	AreTimesLocal bool `json:"AreTimesLocal"`

	// arrival
	// Required: true
	// Min Length: 1
	Arrival *string `json:"Arrival"`

	// date
	// Required: true
	// Min Length: 1
	// Pattern: ^\d{4}\-(0[1-9]|1[0-2])\-(0[1-9]|[1-2][0-9]|3[0-1])$
	Date *string `json:"Date"`

	// day landings
	// Required: true
	// Maximum: 100
	// Minimum: 0
	DayLandings *int32 `json:"DayLandings"`

	// The day/night sequence of the BalloonFlight
	// Required: true
	DayNightSequence struct {
		DayNightSequences
	} `json:"DayNightSequence"`

	// day takeoffs
	// Required: true
	// Maximum: 100
	// Minimum: 0
	DayTakeoffs *int32 `json:"DayTakeoffs"`

	// day time
	// Required: true
	// Min Length: 1
	// Pattern: ^\d{1-5}:[0-5][0-9]$
	DayTime *string `json:"DayTime"`

	// departure
	// Required: true
	// Min Length: 1
	Departure *string `json:"Departure"`

	// distance in kilometers
	DistanceInKilometers int32 `json:"DistanceInKilometers,omitempty"`

	// dual time
	// Required: true
	// Min Length: 1
	// Pattern: ^\d{1-5}:[0-5][0-9]$
	DualTime *string `json:"DualTime"`

	// evening twilight time
	// Pattern: ^([0-1][0-9]|2[0-3]):[0-5][0-9]$
	EveningTwilightTime string `json:"EveningTwilightTime,omitempty"`

	// flight time
	// Required: true
	// Min Length: 1
	// Pattern: ^\d{1-5}:[0-5][0-9]$
	FlightTime *string `json:"FlightTime"`

	// gas time
	// Required: true
	// Min Length: 1
	// Pattern: ^\d{1-5}:[0-5][0-9]$
	GasTime *string `json:"GasTime"`

	// group a time
	// Required: true
	// Min Length: 1
	// Pattern: ^\d{1-5}:[0-5][0-9]$
	GroupATime *string `json:"GroupATime"`

	// group b time
	// Required: true
	// Min Length: 1
	// Pattern: ^\d{1-5}:[0-5][0-9]$
	GroupBTime *string `json:"GroupBTime"`

	// group c time
	// Required: true
	// Min Length: 1
	// Pattern: ^\d{1-5}:[0-5][0-9]$
	GroupCTime *string `json:"GroupCTime"`

	// group d time
	// Required: true
	// Min Length: 1
	// Pattern: ^\d{1-5}:[0-5][0-9]$
	GroupDTime *string `json:"GroupDTime"`

	// inflations
	Inflations int32 `json:"Inflations,omitempty"`

	// instructor time
	// Required: true
	// Min Length: 1
	// Pattern: ^\d{1-5}:[0-5][0-9]$
	InstructorTime *string `json:"InstructorTime"`

	// landing time
	// Required: true
	// Min Length: 1
	// Pattern: ^([0-1][0-9]|2[0-3]):[0-5][0-9]$
	LandingTime *string `json:"LandingTime"`

	// markers
	Markers []*Marker `json:"Markers"`

	// morning twilight time
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

	// night time
	// Required: true
	// Min Length: 1
	// Pattern: ^\d{1-5}:[0-5][0-9]$
	NightTime *string `json:"NightTime"`

	// number of passengers
	NumberOfPassengers int32 `json:"NumberOfPassengers,omitempty"`

	// p i c name
	// Required: true
	// Min Length: 1
	PICName *string `json:"PICName"`

	// p i c time
	// Required: true
	// Min Length: 1
	// Pattern: ^\d{1-5}:[0-5][0-9]$
	PICTime *string `json:"PICTime"`

	// pilot function
	// Required: true
	PilotFunction *PilotFunctions `json:"PilotFunction"`

	// remark
	Remark string `json:"Remark,omitempty"`

	// status
	// Required: true
	Status *UpdateStatus `json:"Status"`

	// takeoff time
	// Required: true
	// Min Length: 1
	// Pattern: ^([0-1][0-9]|2[0-3]):[0-5][0-9]$
	TakeoffTime *string `json:"TakeoffTime"`
}

// Validate validates this balloon flight output base
func (m *BalloonFlightOutputBase) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAircraft(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAreTimesLocal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateArrival(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDayLandings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDayNightSequence(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDayTakeoffs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDayTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeparture(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDualTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEveningTwilightTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlightTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGasTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupATime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupBTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupCTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupDTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstructorTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLandingTime(formats); err != nil {
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

	if err := m.validateNightTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePICName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePICTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePilotFunction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTakeoffTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BalloonFlightOutputBase) validateAircraft(formats strfmt.Registry) error {
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

func (m *BalloonFlightOutputBase) validateAreTimesLocal(formats strfmt.Registry) error {

	if err := validate.Required("AreTimesLocal", "body", bool(m.AreTimesLocal)); err != nil {
		return err
	}

	return nil
}

func (m *BalloonFlightOutputBase) validateArrival(formats strfmt.Registry) error {

	if err := validate.Required("Arrival", "body", m.Arrival); err != nil {
		return err
	}

	if err := validate.MinLength("Arrival", "body", *m.Arrival, 1); err != nil {
		return err
	}

	return nil
}

func (m *BalloonFlightOutputBase) validateDate(formats strfmt.Registry) error {

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

func (m *BalloonFlightOutputBase) validateDayLandings(formats strfmt.Registry) error {

	if err := validate.Required("DayLandings", "body", m.DayLandings); err != nil {
		return err
	}

	if err := validate.MinimumInt("DayLandings", "body", int64(*m.DayLandings), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("DayLandings", "body", int64(*m.DayLandings), 100, false); err != nil {
		return err
	}

	return nil
}

func (m *BalloonFlightOutputBase) validateDayNightSequence(formats strfmt.Registry) error {

	return nil
}

func (m *BalloonFlightOutputBase) validateDayTakeoffs(formats strfmt.Registry) error {

	if err := validate.Required("DayTakeoffs", "body", m.DayTakeoffs); err != nil {
		return err
	}

	if err := validate.MinimumInt("DayTakeoffs", "body", int64(*m.DayTakeoffs), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("DayTakeoffs", "body", int64(*m.DayTakeoffs), 100, false); err != nil {
		return err
	}

	return nil
}

func (m *BalloonFlightOutputBase) validateDayTime(formats strfmt.Registry) error {

	if err := validate.Required("DayTime", "body", m.DayTime); err != nil {
		return err
	}

	if err := validate.MinLength("DayTime", "body", *m.DayTime, 1); err != nil {
		return err
	}

	if err := validate.Pattern("DayTime", "body", *m.DayTime, `^\d{1-5}:[0-5][0-9]$`); err != nil {
		return err
	}

	return nil
}

func (m *BalloonFlightOutputBase) validateDeparture(formats strfmt.Registry) error {

	if err := validate.Required("Departure", "body", m.Departure); err != nil {
		return err
	}

	if err := validate.MinLength("Departure", "body", *m.Departure, 1); err != nil {
		return err
	}

	return nil
}

func (m *BalloonFlightOutputBase) validateDualTime(formats strfmt.Registry) error {

	if err := validate.Required("DualTime", "body", m.DualTime); err != nil {
		return err
	}

	if err := validate.MinLength("DualTime", "body", *m.DualTime, 1); err != nil {
		return err
	}

	if err := validate.Pattern("DualTime", "body", *m.DualTime, `^\d{1-5}:[0-5][0-9]$`); err != nil {
		return err
	}

	return nil
}

func (m *BalloonFlightOutputBase) validateEveningTwilightTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EveningTwilightTime) { // not required
		return nil
	}

	if err := validate.Pattern("EveningTwilightTime", "body", m.EveningTwilightTime, `^([0-1][0-9]|2[0-3]):[0-5][0-9]$`); err != nil {
		return err
	}

	return nil
}

func (m *BalloonFlightOutputBase) validateFlightTime(formats strfmt.Registry) error {

	if err := validate.Required("FlightTime", "body", m.FlightTime); err != nil {
		return err
	}

	if err := validate.MinLength("FlightTime", "body", *m.FlightTime, 1); err != nil {
		return err
	}

	if err := validate.Pattern("FlightTime", "body", *m.FlightTime, `^\d{1-5}:[0-5][0-9]$`); err != nil {
		return err
	}

	return nil
}

func (m *BalloonFlightOutputBase) validateGasTime(formats strfmt.Registry) error {

	if err := validate.Required("GasTime", "body", m.GasTime); err != nil {
		return err
	}

	if err := validate.MinLength("GasTime", "body", *m.GasTime, 1); err != nil {
		return err
	}

	if err := validate.Pattern("GasTime", "body", *m.GasTime, `^\d{1-5}:[0-5][0-9]$`); err != nil {
		return err
	}

	return nil
}

func (m *BalloonFlightOutputBase) validateGroupATime(formats strfmt.Registry) error {

	if err := validate.Required("GroupATime", "body", m.GroupATime); err != nil {
		return err
	}

	if err := validate.MinLength("GroupATime", "body", *m.GroupATime, 1); err != nil {
		return err
	}

	if err := validate.Pattern("GroupATime", "body", *m.GroupATime, `^\d{1-5}:[0-5][0-9]$`); err != nil {
		return err
	}

	return nil
}

func (m *BalloonFlightOutputBase) validateGroupBTime(formats strfmt.Registry) error {

	if err := validate.Required("GroupBTime", "body", m.GroupBTime); err != nil {
		return err
	}

	if err := validate.MinLength("GroupBTime", "body", *m.GroupBTime, 1); err != nil {
		return err
	}

	if err := validate.Pattern("GroupBTime", "body", *m.GroupBTime, `^\d{1-5}:[0-5][0-9]$`); err != nil {
		return err
	}

	return nil
}

func (m *BalloonFlightOutputBase) validateGroupCTime(formats strfmt.Registry) error {

	if err := validate.Required("GroupCTime", "body", m.GroupCTime); err != nil {
		return err
	}

	if err := validate.MinLength("GroupCTime", "body", *m.GroupCTime, 1); err != nil {
		return err
	}

	if err := validate.Pattern("GroupCTime", "body", *m.GroupCTime, `^\d{1-5}:[0-5][0-9]$`); err != nil {
		return err
	}

	return nil
}

func (m *BalloonFlightOutputBase) validateGroupDTime(formats strfmt.Registry) error {

	if err := validate.Required("GroupDTime", "body", m.GroupDTime); err != nil {
		return err
	}

	if err := validate.MinLength("GroupDTime", "body", *m.GroupDTime, 1); err != nil {
		return err
	}

	if err := validate.Pattern("GroupDTime", "body", *m.GroupDTime, `^\d{1-5}:[0-5][0-9]$`); err != nil {
		return err
	}

	return nil
}

func (m *BalloonFlightOutputBase) validateInstructorTime(formats strfmt.Registry) error {

	if err := validate.Required("InstructorTime", "body", m.InstructorTime); err != nil {
		return err
	}

	if err := validate.MinLength("InstructorTime", "body", *m.InstructorTime, 1); err != nil {
		return err
	}

	if err := validate.Pattern("InstructorTime", "body", *m.InstructorTime, `^\d{1-5}:[0-5][0-9]$`); err != nil {
		return err
	}

	return nil
}

func (m *BalloonFlightOutputBase) validateLandingTime(formats strfmt.Registry) error {

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

func (m *BalloonFlightOutputBase) validateMarkers(formats strfmt.Registry) error {
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

func (m *BalloonFlightOutputBase) validateMorningTwilightTime(formats strfmt.Registry) error {
	if swag.IsZero(m.MorningTwilightTime) { // not required
		return nil
	}

	if err := validate.Pattern("MorningTwilightTime", "body", m.MorningTwilightTime, `^([0-1][0-9]|2[0-3]):[0-5][0-9]$`); err != nil {
		return err
	}

	return nil
}

func (m *BalloonFlightOutputBase) validateNightLandings(formats strfmt.Registry) error {

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

func (m *BalloonFlightOutputBase) validateNightTakeoffs(formats strfmt.Registry) error {

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

func (m *BalloonFlightOutputBase) validateNightTime(formats strfmt.Registry) error {

	if err := validate.Required("NightTime", "body", m.NightTime); err != nil {
		return err
	}

	if err := validate.MinLength("NightTime", "body", *m.NightTime, 1); err != nil {
		return err
	}

	if err := validate.Pattern("NightTime", "body", *m.NightTime, `^\d{1-5}:[0-5][0-9]$`); err != nil {
		return err
	}

	return nil
}

func (m *BalloonFlightOutputBase) validatePICName(formats strfmt.Registry) error {

	if err := validate.Required("PICName", "body", m.PICName); err != nil {
		return err
	}

	if err := validate.MinLength("PICName", "body", *m.PICName, 1); err != nil {
		return err
	}

	return nil
}

func (m *BalloonFlightOutputBase) validatePICTime(formats strfmt.Registry) error {

	if err := validate.Required("PICTime", "body", m.PICTime); err != nil {
		return err
	}

	if err := validate.MinLength("PICTime", "body", *m.PICTime, 1); err != nil {
		return err
	}

	if err := validate.Pattern("PICTime", "body", *m.PICTime, `^\d{1-5}:[0-5][0-9]$`); err != nil {
		return err
	}

	return nil
}

func (m *BalloonFlightOutputBase) validatePilotFunction(formats strfmt.Registry) error {

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

func (m *BalloonFlightOutputBase) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("Status", "body", m.Status); err != nil {
		return err
	}

	if err := validate.Required("Status", "body", m.Status); err != nil {
		return err
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Status")
			}
			return err
		}
	}

	return nil
}

func (m *BalloonFlightOutputBase) validateTakeoffTime(formats strfmt.Registry) error {

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

// ContextValidate validate this balloon flight output base based on the context it is used
func (m *BalloonFlightOutputBase) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAircraft(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDayNightSequence(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMarkers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePilotFunction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BalloonFlightOutputBase) contextValidateAircraft(ctx context.Context, formats strfmt.Registry) error {

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

func (m *BalloonFlightOutputBase) contextValidateDayNightSequence(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *BalloonFlightOutputBase) contextValidateMarkers(ctx context.Context, formats strfmt.Registry) error {

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

func (m *BalloonFlightOutputBase) contextValidatePilotFunction(ctx context.Context, formats strfmt.Registry) error {

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

func (m *BalloonFlightOutputBase) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {

		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BalloonFlightOutputBase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BalloonFlightOutputBase) UnmarshalBinary(b []byte) error {
	var res BalloonFlightOutputBase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
