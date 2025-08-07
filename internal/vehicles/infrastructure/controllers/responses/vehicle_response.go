package responses

import (
	"time"

	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
)

type VehicleResponse struct {
	ID                 string     `json:"id"`
	Brand              string     `json:"brand"`
	Model              string     `json:"model"`
	Color              string     `json:"color"`
	Description        string     `json:"description"`
	Price              float64    `json:"price"`
	Status             string     `json:"status"`
	Condition          string     `json:"condition"`
	Year               int        `json:"year"`
	HasAirConditioning bool       `json:"has_air_conditioning"`
	HasAirbag          bool       `json:"has_airbag"`
	HasAbsBrakes       bool       `json:"has_abs_brakes"`
	HasPowerSteering   bool       `json:"has_power_steering"`
	HasPowerWindows    bool       `json:"has_power_windows"`
	HasPowerLocks      bool       `json:"has_power_locks"`
	HasMultimedia      bool       `json:"has_multimedia"`
	HasAlarm           bool       `json:"has_alarm"`
	HasTractionControl bool       `json:"has_traction_control"`
	HasRearCamera      bool       `json:"has_rear_camera"`
	HasParkingSensors  bool       `json:"has_parking_sensors"`
	BodyType           string     `json:"body_type"`
	Transmission       string     `json:"transmission"`
	FuelType           string     `json:"fuel_type"`
	Mileage            int        `json:"mileage"`
	Doors              int        `json:"doors"`
	Engine             string     `json:"engine"`
	CreatedAt          time.Time  `json:"created_at,omitempty"`
	UpdatedAt          *time.Time `json:"updated_at,omitempty"`
}

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
