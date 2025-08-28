package responses

import (
  "time"

  "github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
)

// VehicleResponse
// @Description Represents a vehicle
type VehicleResponse struct {
  ID                 string     `json:"id,omitempty"`
  Brand              string     `json:"brand,omitempty"`
  Model              string     `json:"model,omitempty"`
  Color              string     `json:"color,omitempty"`
  Description        string     `json:"description,omitempty"`
  Price              float64    `json:"price,omitempty"`
  Status             string     `json:"status,omitempty"`
  Condition          string     `json:"condition,omitempty"`
  Year               int        `json:"year,omitempty"`
  HasAirConditioning bool       `json:"hasAirConditioning,omitempty"`
  HasAirbag          bool       `json:"hasAirbag,omitempty"`
  HasAbsBrakes       bool       `json:"hasAbsBrakes,omitempty"`
  HasPowerSteering   bool       `json:"hasPowerSteering,omitempty"`
  HasPowerWindows    bool       `json:"hasPowerWindows,omitempty"`
  HasPowerLocks      bool       `json:"hasPowerLocks,omitempty"`
  HasMultimedia      bool       `json:"hasMultimedia,omitempty"`
  HasAlarm           bool       `json:"hasAlarm,omitempty"`
  HasTractionControl bool       `json:"hasTractionControl,omitempty"`
  HasRearCamera      bool       `json:"hasRearCamera,omitempty"`
  HasParkingSensors  bool       `json:"hasParkingSensors,omitempty"`
  BodyType           string     `json:"bodyType,omitempty"`
  Transmission       string     `json:"transmission,omitempty"`
  FuelType           string     `json:"fuelType,omitempty"`
  Mileage            int        `json:"mileage,omitempty"`
  Doors              int        `json:"doors,omitempty"`
  Engine             string     `json:"engine,omitempty"`
  CreatedAt          *time.Time `json:"createdAt,omitempty"`
  UpdatedAt          *time.Time `json:"updatedAt,omitempty"`
} // @name VehicleResponse

func NewVehicleResponse(vehicle *entities.Vehicle) *VehicleResponse {
  return &VehicleResponse{
    ID:                 vehicle.ID,
    Brand:              vehicle.Brand,
    Model:              vehicle.Model,
    Color:              vehicle.Color,
    Description:        vehicle.Description,
    Price:              vehicle.GetPrice(),
    Status:             vehicle.GetStatus(),
    Condition:          vehicle.GetCondition(),
    Year:               vehicle.GetYear(),
    HasAirConditioning: vehicle.Features.HasAirConditioning,
    HasAirbag:          vehicle.Features.HasAirbag,
    HasAbsBrakes:       vehicle.Features.HasAbsBrakes,
    HasPowerSteering:   vehicle.Features.HasPowerSteering,
    HasPowerWindows:    vehicle.Features.HasPowerWindows,
    HasPowerLocks:      vehicle.Features.HasPowerLocks,
    HasMultimedia:      vehicle.Features.HasMultimedia,
    HasAlarm:           vehicle.Features.HasAlarm,
    HasTractionControl: vehicle.Features.HasTractionControl,
    HasRearCamera:      vehicle.Features.HasRearCamera,
    HasParkingSensors:  vehicle.Features.HasParkingSensors,
    BodyType:           vehicle.Specification.GetBodyType(),
    Transmission:       vehicle.Specification.GetTransmission(),
    FuelType:           vehicle.Specification.GetFuelType(),
    Mileage:            vehicle.Specification.GetMileage(),
    Doors:              vehicle.Specification.GetDoors(),
    Engine:             vehicle.Specification.GetEngine(),
  }
}

func NewUpdateResponse(vehicle *entities.UpdateVehicleInput) *VehicleResponse {
  output := new(VehicleResponse)
  if vehicle.Price != nil {
    output.Price = vehicle.Price.Value()
  }
  if vehicle.Status != nil {
    output.Status = vehicle.Status.Value()
  }
  if vehicle.Condition != nil {
    output.Condition = vehicle.Condition.Value()
  }
  if vehicle.Description != nil {
    output.Description = *vehicle.Description
  }

  if vehicle.Features != nil {
    output.HasAirConditioning = *vehicle.Features.HasAirConditioning
    output.HasAirbag = *vehicle.Features.HasAirbag
    output.HasAbsBrakes = *vehicle.Features.HasAbsBrakes
    output.HasPowerSteering = *vehicle.Features.HasPowerSteering
    output.HasPowerWindows = *vehicle.Features.HasPowerWindows
    output.HasPowerLocks = *vehicle.Features.HasPowerLocks
    output.HasMultimedia = *vehicle.Features.HasMultimedia
    output.HasAlarm = *vehicle.Features.HasAlarm
    output.HasTractionControl = *vehicle.Features.HasTractionControl
    output.HasRearCamera = *vehicle.Features.HasRearCamera
    output.HasParkingSensors = *vehicle.Features.HasParkingSensors
  }
  return output
}
