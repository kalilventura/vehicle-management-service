package valueobjects

import (
	"github.com/kalilventura/vehicle-management/internal/shared/domain"
)

type FeaturesProps struct {
	HasAirConditioning bool
	HasAirbag          bool
	HasAbsBrakes       bool
	HasPowerSteering   bool
	HasPowerWindows    bool
	HasPowerLocks      bool
	HasMultimedia      bool
	HasAlarm           bool
	HasTractionControl bool
	HasRearCamera      bool
	HasParkingSensors  bool
}

// Features represents vehicle features
type Features struct {
	domain.BaseValueObject[FeaturesProps]
}

// NewFeatures creates a new Features value object
func NewFeatures(props FeaturesProps) Features {
	return Features{
		BaseValueObject: domain.NewBaseValueObject(props),
	}
}

// HasAirConditioning returns if the vehicle has air conditioning
func (f Features) HasAirConditioning() bool {
	return f.Props().HasAirConditioning
}

// HasAirbag returns if the vehicle has airbag
func (f Features) HasAirbag() bool {
	return f.Props().HasAirbag
}

// HasAbsBrakes returns if the vehicle has ABS brakes
func (f Features) HasAbsBrakes() bool {
	return f.Props().HasAbsBrakes
}

// HasPowerSteering returns if the vehicle has power steering
func (f Features) HasPowerSteering() bool {
	return f.Props().HasPowerSteering
}

// HasPowerWindows returns if the vehicle has power windows
func (f Features) HasPowerWindows() bool {
	return f.Props().HasPowerWindows
}

// HasPowerLocks returns if the vehicle has power locks
func (f Features) HasPowerLocks() bool {
	return f.Props().HasPowerLocks
}

// HasMultimedia returns if the vehicle has multimedia
func (f Features) HasMultimedia() bool {
	return f.Props().HasMultimedia
}

// HasAlarm returns if the vehicle has alarm
func (f Features) HasAlarm() bool {
	return f.Props().HasAlarm
}

// HasTractionControl returns if the vehicle has traction control
func (f Features) HasTractionControl() bool {
	return f.Props().HasTractionControl
}

// HasRearCamera returns if the vehicle has rear camera
func (f Features) HasRearCamera() bool {
	return f.Props().HasRearCamera
}

// HasParkingSensors returns if the vehicle has parking sensors
func (f Features) HasParkingSensors() bool {
	return f.Props().HasParkingSensors
}

