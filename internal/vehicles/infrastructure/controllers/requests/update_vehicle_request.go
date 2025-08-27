package requests

import (
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities/dtos"
)

// UpdateVehicleRequest
// @Description Object that represents a vehicle to update
type UpdateVehicleRequest struct {
	Price              *float64 `json:"price,omitempty"`
	Mileage            *int     `json:"mileage,omitempty"`
	Status             *string  `json:"status,omitempty"`
	Color              *string  `json:"color,omitempty"`
	Description        *string  `json:"description,omitempty"`
	Condition          *string  `json:"condition,omitempty"`
	HasAirConditioning *bool    `json:"has_air_conditioning,omitempty"`
	HasAirbag          *bool    `json:"has_airbag,omitempty"`
	HasAbsBrakes       *bool    `json:"has_abs_brakes,omitempty"`
	HasPowerSteering   *bool    `json:"has_power_steering,omitempty"`
	HasPowerWindows    *bool    `json:"has_power_windows,omitempty"`
	HasPowerLocks      *bool    `json:"has_power_locks,omitempty"`
	HasMultimedia      *bool    `json:"has_multimedia,omitempty"`
	HasAlarm           *bool    `json:"has_alarm,omitempty"`
	HasTractionControl *bool    `json:"has_traction_control,omitempty"`
	HasRearCamera      *bool    `json:"has_rear_camera,omitempty"`
	HasParkingSensors  *bool    `json:"has_parking_sensors,omitempty"`
} // @name UpdateVehicleRequest

func (r UpdateVehicleRequest) ToDomain(id string) (*entities.UpdateVehicleInput, error) {
	input := &entities.UpdateVehicleInput{
		ID:          id,
		Color:       r.Color,
		Description: r.Description,
	}

	if r.Price != nil {
		price, _ := dtos.NewPrice(*r.Price)
		input.Price = &price
	}
	if r.Status != nil {
		status, _ := dtos.NewStatus(*r.Status)
		input.Status = &status
	}
	if r.Condition != nil {
		condition, _ := dtos.NewCondition(*r.Condition)
		input.Condition = &condition
	}

	if r.hasAnyFeature() {
		input.Features = &entities.UpdateFeaturesInput{
			HasAirConditioning: r.HasAirConditioning,
			HasAirbag:          r.HasAirbag,
			HasAbsBrakes:       r.HasAbsBrakes,
			HasPowerSteering:   r.HasPowerSteering,
			HasPowerWindows:    r.HasPowerWindows,
			HasPowerLocks:      r.HasPowerLocks,
			HasMultimedia:      r.HasMultimedia,
			HasAlarm:           r.HasAlarm,
			HasTractionControl: r.HasTractionControl,
			HasRearCamera:      r.HasRearCamera,
			HasParkingSensors:  r.HasParkingSensors,
		}
	}
	return input, nil
}

func (r UpdateVehicleRequest) hasAnyFeature() bool {
	return r.HasAirConditioning != nil ||
		r.HasAirbag != nil ||
		r.HasAbsBrakes != nil ||
		r.HasPowerSteering != nil ||
		r.HasPowerWindows != nil ||
		r.HasPowerLocks != nil ||
		r.HasMultimedia != nil ||
		r.HasAlarm != nil ||
		r.HasTractionControl != nil ||
		r.HasRearCamera != nil ||
		r.HasParkingSensors != nil
}
