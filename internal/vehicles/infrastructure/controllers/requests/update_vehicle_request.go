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
	HasAirConditioning *bool    `json:"hasAirConditioning,omitempty"`
	HasAirbag          *bool    `json:"hasAirbag,omitempty"`
	HasAbsBrakes       *bool    `json:"hasAbsBrakes,omitempty"`
	HasPowerSteering   *bool    `json:"hasPowerSteering,omitempty"`
	HasPowerWindows    *bool    `json:"hasPowerWindows,omitempty"`
	HasPowerLocks      *bool    `json:"hasPowerLocks,omitempty"`
	HasMultimedia      *bool    `json:"hasMultimedia,omitempty"`
	HasAlarm           *bool    `json:"hasAlarm,omitempty"`
	HasTractionControl *bool    `json:"hasTractionControl,omitempty"`
	HasRearCamera      *bool    `json:"hasRearCamera,omitempty"`
	HasParkingSensors  *bool    `json:"hasParkingSensors,omitempty"`
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
