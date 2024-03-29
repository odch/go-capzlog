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

// MarkerType marker type
//
// swagger:model MarkerType
type MarkerType string

func NewMarkerType(value MarkerType) *MarkerType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated MarkerType.
func (m MarkerType) Pointer() *MarkerType {
	return &m
}

const (

	// MarkerTypeCustom captures enum value "Custom"
	MarkerTypeCustom MarkerType = "Custom"

	// MarkerTypeSkillTest captures enum value "SkillTest"
	MarkerTypeSkillTest MarkerType = "SkillTest"

	// MarkerTypeProficiencyCheck captures enum value "ProficiencyCheck"
	MarkerTypeProficiencyCheck MarkerType = "ProficiencyCheck"

	// MarkerTypeAssessmentOfCompetence captures enum value "AssessmentOfCompetence"
	MarkerTypeAssessmentOfCompetence MarkerType = "AssessmentOfCompetence"

	// MarkerTypeOperatorProficiencyCheck captures enum value "OperatorProficiencyCheck"
	MarkerTypeOperatorProficiencyCheck MarkerType = "OperatorProficiencyCheck"

	// MarkerTypeOperatorLineCheck captures enum value "OperatorLineCheck"
	MarkerTypeOperatorLineCheck MarkerType = "OperatorLineCheck"

	// MarkerTypeExaminerFlight captures enum value "ExaminerFlight"
	MarkerTypeExaminerFlight MarkerType = "ExaminerFlight"

	// MarkerTypeLanguageProficiencyCheck captures enum value "LanguageProficiencyCheck"
	MarkerTypeLanguageProficiencyCheck MarkerType = "LanguageProficiencyCheck"

	// MarkerTypeRefresherTraining captures enum value "RefresherTraining"
	MarkerTypeRefresherTraining MarkerType = "RefresherTraining"

	// MarkerTypeDemonstrationFlight captures enum value "DemonstrationFlight"
	MarkerTypeDemonstrationFlight MarkerType = "DemonstrationFlight"

	// MarkerTypeTrainingFlight captures enum value "TrainingFlight"
	MarkerTypeTrainingFlight MarkerType = "TrainingFlight"

	// MarkerTypeFamiliarizationFlight captures enum value "FamiliarizationFlight"
	MarkerTypeFamiliarizationFlight MarkerType = "FamiliarizationFlight"

	// MarkerTypeDifferenceTraining captures enum value "DifferenceTraining"
	MarkerTypeDifferenceTraining MarkerType = "DifferenceTraining"

	// MarkerTypeLineFlyingUnderSupervision captures enum value "LineFlyingUnderSupervision"
	MarkerTypeLineFlyingUnderSupervision MarkerType = "LineFlyingUnderSupervision"

	// MarkerTypeLandingTraining captures enum value "LandingTraining"
	MarkerTypeLandingTraining MarkerType = "LandingTraining"

	// MarkerTypeZeroFlightTimeTraining captures enum value "ZeroFlightTimeTraining"
	MarkerTypeZeroFlightTimeTraining MarkerType = "ZeroFlightTimeTraining"

	// MarkerTypeLaunchPrivilege captures enum value "LaunchPrivilege"
	MarkerTypeLaunchPrivilege MarkerType = "LaunchPrivilege"

	// MarkerTypeAerobaticPrivilege captures enum value "AerobaticPrivilege"
	MarkerTypeAerobaticPrivilege MarkerType = "AerobaticPrivilege"

	// MarkerTypeCloudFlyingPrivilege captures enum value "CloudFlyingPrivilege"
	MarkerTypeCloudFlyingPrivilege MarkerType = "CloudFlyingPrivilege"

	// MarkerTypeCourseCompleted captures enum value "CourseCompleted"
	MarkerTypeCourseCompleted MarkerType = "CourseCompleted"

	// MarkerTypeAdditionalRatingTrainingCourse captures enum value "AdditionalRatingTrainingCourse"
	MarkerTypeAdditionalRatingTrainingCourse MarkerType = "AdditionalRatingTrainingCourse"

	// MarkerTypeInstructorTrainingCourse captures enum value "InstructorTrainingCourse"
	MarkerTypeInstructorTrainingCourse MarkerType = "InstructorTrainingCourse"

	// MarkerTypeDemonstrationTheAbilityToInstruct captures enum value "DemonstrationTheAbilityToInstruct"
	MarkerTypeDemonstrationTheAbilityToInstruct MarkerType = "DemonstrationTheAbilityToInstruct"

	// MarkerTypeWithStudent captures enum value "WithStudent"
	MarkerTypeWithStudent MarkerType = "WithStudent"

	// MarkerTypeWithInstructor captures enum value "WithInstructor"
	MarkerTypeWithInstructor MarkerType = "WithInstructor"

	// MarkerTypeWithExaminer captures enum value "WithExaminer"
	MarkerTypeWithExaminer MarkerType = "WithExaminer"

	// MarkerTypeLesson captures enum value "Lesson"
	MarkerTypeLesson MarkerType = "Lesson"

	// MarkerTypeEitherSeatQualification captures enum value "EitherSeatQualification"
	MarkerTypeEitherSeatQualification MarkerType = "EitherSeatQualification"

	// MarkerTypeRecurrentMountainPaxCheckHelicopter captures enum value "RecurrentMountainPaxCheckHelicopter"
	MarkerTypeRecurrentMountainPaxCheckHelicopter MarkerType = "RecurrentMountainPaxCheckHelicopter"

	// MarkerTypeLOFT captures enum value "LOFT"
	MarkerTypeLOFT MarkerType = "LOFT"

	// MarkerTypeOperatorProficiencyTraining captures enum value "OperatorProficiencyTraining"
	MarkerTypeOperatorProficiencyTraining MarkerType = "OperatorProficiencyTraining"

	// MarkerTypeTrainingPhase captures enum value "TrainingPhase"
	MarkerTypeTrainingPhase MarkerType = "TrainingPhase"

	// MarkerTypeSolo captures enum value "Solo"
	MarkerTypeSolo MarkerType = "Solo"

	// MarkerTypeCrossCountry captures enum value "CrossCountry"
	MarkerTypeCrossCountry MarkerType = "CrossCountry"

	// MarkerTypeSeaTakeoff captures enum value "SeaTakeoff"
	MarkerTypeSeaTakeoff MarkerType = "SeaTakeoff"

	// MarkerTypeSeaLanding captures enum value "SeaLanding"
	MarkerTypeSeaLanding MarkerType = "SeaLanding"

	// MarkerTypeSeriesOfFlights captures enum value "SeriesOfFlights"
	MarkerTypeSeriesOfFlights MarkerType = "SeriesOfFlights"

	// MarkerTypeReducedTimeOnControls captures enum value "ReducedTimeOnControls"
	MarkerTypeReducedTimeOnControls MarkerType = "ReducedTimeOnControls"

	// MarkerTypeSailplaneTowing captures enum value "SailplaneTowing"
	MarkerTypeSailplaneTowing MarkerType = "SailplaneTowing"

	// MarkerTypeBannerTowing captures enum value "BannerTowing"
	MarkerTypeBannerTowing MarkerType = "BannerTowing"

	// MarkerTypeAerobatic captures enum value "Aerobatic"
	MarkerTypeAerobatic MarkerType = "Aerobatic"

	// MarkerTypeSelfSustainable captures enum value "SelfSustainable"
	MarkerTypeSelfSustainable MarkerType = "SelfSustainable"

	// MarkerTypeCloudFlying captures enum value "CloudFlying"
	MarkerTypeCloudFlying MarkerType = "CloudFlying"

	// MarkerTypeTetheredFlight captures enum value "TetheredFlight"
	MarkerTypeTetheredFlight MarkerType = "TetheredFlight"

	// MarkerTypeNightVisionImagingSystem captures enum value "NightVisionImagingSystem"
	MarkerTypeNightVisionImagingSystem MarkerType = "NightVisionImagingSystem"

	// MarkerTypeFlightTest captures enum value "FlightTest"
	MarkerTypeFlightTest MarkerType = "FlightTest"

	// MarkerTypeParadropping captures enum value "Paradropping"
	MarkerTypeParadropping MarkerType = "Paradropping"

	// MarkerTypeLowVisibilityTakeoff captures enum value "LowVisibilityTakeoff"
	MarkerTypeLowVisibilityTakeoff MarkerType = "LowVisibilityTakeoff"

	// MarkerTypeApproachType captures enum value "ApproachType"
	MarkerTypeApproachType MarkerType = "ApproachType"

	// MarkerTypeLowVisibilityLandingAirplane captures enum value "LowVisibilityLandingAirplane"
	MarkerTypeLowVisibilityLandingAirplane MarkerType = "LowVisibilityLandingAirplane"

	// MarkerTypeHelicopterDepartureInFog captures enum value "HelicopterDepartureInFog"
	MarkerTypeHelicopterDepartureInFog MarkerType = "HelicopterDepartureInFog"

	// MarkerTypeMountainTakeoff captures enum value "MountainTakeoff"
	MarkerTypeMountainTakeoff MarkerType = "MountainTakeoff"

	// MarkerTypeMountainTakeoffHelicopter captures enum value "MountainTakeoffHelicopter"
	MarkerTypeMountainTakeoffHelicopter MarkerType = "MountainTakeoffHelicopter"

	// MarkerTypeMountainLanding captures enum value "MountainLanding"
	MarkerTypeMountainLanding MarkerType = "MountainLanding"

	// MarkerTypeMountainLandingHelicopter captures enum value "MountainLandingHelicopter"
	MarkerTypeMountainLandingHelicopter MarkerType = "MountainLandingHelicopter"

	// MarkerTypeMountainLandingOfficialSiteHelicopter captures enum value "MountainLandingOfficialSiteHelicopter"
	MarkerTypeMountainLandingOfficialSiteHelicopter MarkerType = "MountainLandingOfficialSiteHelicopter"

	// MarkerTypeMountainLandingAbove2000mHelicopter captures enum value "MountainLandingAbove2000mHelicopter"
	MarkerTypeMountainLandingAbove2000mHelicopter MarkerType = "MountainLandingAbove2000mHelicopter"

	// MarkerTypeMountainLandingAbove2700mHelicopter captures enum value "MountainLandingAbove2700mHelicopter"
	MarkerTypeMountainLandingAbove2700mHelicopter MarkerType = "MountainLandingAbove2700mHelicopter"

	// MarkerTypeTakeoffRunway captures enum value "TakeoffRunway"
	MarkerTypeTakeoffRunway MarkerType = "TakeoffRunway"

	// MarkerTypeLandingRunway captures enum value "LandingRunway"
	MarkerTypeLandingRunway MarkerType = "LandingRunway"

	// MarkerTypeHoldingPattern captures enum value "HoldingPattern"
	MarkerTypeHoldingPattern MarkerType = "HoldingPattern"

	// MarkerTypeGoAround captures enum value "GoAround"
	MarkerTypeGoAround MarkerType = "GoAround"

	// MarkerTypeRunwaySwingOver captures enum value "RunwaySwingOver"
	MarkerTypeRunwaySwingOver MarkerType = "RunwaySwingOver"

	// MarkerTypeSteepApproach captures enum value "SteepApproach"
	MarkerTypeSteepApproach MarkerType = "SteepApproach"

	// MarkerTypeGlacierLanding captures enum value "GlacierLanding"
	MarkerTypeGlacierLanding MarkerType = "GlacierLanding"

	// MarkerTypeTouchAndGo captures enum value "TouchAndGo"
	MarkerTypeTouchAndGo MarkerType = "TouchAndGo"

	// MarkerTypeEVS captures enum value "EVS"
	MarkerTypeEVS MarkerType = "EVS"

	// MarkerTypeHESLO1 captures enum value "HESLO1"
	MarkerTypeHESLO1 MarkerType = "HESLO1"

	// MarkerTypeHESLO2 captures enum value "HESLO2"
	MarkerTypeHESLO2 MarkerType = "HESLO2"

	// MarkerTypeHESLO3 captures enum value "HESLO3"
	MarkerTypeHESLO3 MarkerType = "HESLO3"

	// MarkerTypeHESLO4 captures enum value "HESLO4"
	MarkerTypeHESLO4 MarkerType = "HESLO4"

	// MarkerTypeHEC1 captures enum value "HEC1"
	MarkerTypeHEC1 MarkerType = "HEC1"

	// MarkerTypeHEC2 captures enum value "HEC2"
	MarkerTypeHEC2 MarkerType = "HEC2"

	// MarkerTypeHHO captures enum value "HHO"
	MarkerTypeHHO MarkerType = "HHO"

	// MarkerTypeReliefPilot captures enum value "ReliefPilot"
	MarkerTypeReliefPilot MarkerType = "ReliefPilot"

	// MarkerTypeExaminer captures enum value "Examiner"
	MarkerTypeExaminer MarkerType = "Examiner"

	// MarkerTypeLanguageAssessor captures enum value "LanguageAssessor"
	MarkerTypeLanguageAssessor MarkerType = "LanguageAssessor"

	// MarkerTypePilotRole captures enum value "PilotRole"
	MarkerTypePilotRole MarkerType = "PilotRole"

	// MarkerTypeWithMate captures enum value "WithMate"
	MarkerTypeWithMate MarkerType = "WithMate"

	// MarkerTypeWithPassengers captures enum value "WithPassengers"
	MarkerTypeWithPassengers MarkerType = "WithPassengers"

	// MarkerTypeNumberOfPassengers captures enum value "NumberOfPassengers"
	MarkerTypeNumberOfPassengers MarkerType = "NumberOfPassengers"

	// MarkerTypeCargoWeight captures enum value "CargoWeight"
	MarkerTypeCargoWeight MarkerType = "CargoWeight"

	// MarkerTypeFlightNumber captures enum value "FlightNumber"
	MarkerTypeFlightNumber MarkerType = "FlightNumber"

	// MarkerTypeCrew captures enum value "Crew"
	MarkerTypeCrew MarkerType = "Crew"

	// MarkerTypeBirdStrike captures enum value "BirdStrike"
	MarkerTypeBirdStrike MarkerType = "BirdStrike"
)

// for schema
var markerTypeEnum []interface{}

func init() {
	var res []MarkerType
	if err := json.Unmarshal([]byte(`["Custom","SkillTest","ProficiencyCheck","AssessmentOfCompetence","OperatorProficiencyCheck","OperatorLineCheck","ExaminerFlight","LanguageProficiencyCheck","RefresherTraining","DemonstrationFlight","TrainingFlight","FamiliarizationFlight","DifferenceTraining","LineFlyingUnderSupervision","LandingTraining","ZeroFlightTimeTraining","LaunchPrivilege","AerobaticPrivilege","CloudFlyingPrivilege","CourseCompleted","AdditionalRatingTrainingCourse","InstructorTrainingCourse","DemonstrationTheAbilityToInstruct","WithStudent","WithInstructor","WithExaminer","Lesson","EitherSeatQualification","RecurrentMountainPaxCheckHelicopter","LOFT","OperatorProficiencyTraining","TrainingPhase","Solo","CrossCountry","SeaTakeoff","SeaLanding","SeriesOfFlights","ReducedTimeOnControls","SailplaneTowing","BannerTowing","Aerobatic","SelfSustainable","CloudFlying","TetheredFlight","NightVisionImagingSystem","FlightTest","Paradropping","LowVisibilityTakeoff","ApproachType","LowVisibilityLandingAirplane","HelicopterDepartureInFog","MountainTakeoff","MountainTakeoffHelicopter","MountainLanding","MountainLandingHelicopter","MountainLandingOfficialSiteHelicopter","MountainLandingAbove2000mHelicopter","MountainLandingAbove2700mHelicopter","TakeoffRunway","LandingRunway","HoldingPattern","GoAround","RunwaySwingOver","SteepApproach","GlacierLanding","TouchAndGo","EVS","HESLO1","HESLO2","HESLO3","HESLO4","HEC1","HEC2","HHO","ReliefPilot","Examiner","LanguageAssessor","PilotRole","WithMate","WithPassengers","NumberOfPassengers","CargoWeight","FlightNumber","Crew","BirdStrike"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		markerTypeEnum = append(markerTypeEnum, v)
	}
}

func (m MarkerType) validateMarkerTypeEnum(path, location string, value MarkerType) error {
	if err := validate.EnumCase(path, location, value, markerTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this marker type
func (m MarkerType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateMarkerTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this marker type based on context it is used
func (m MarkerType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
