package requests

import (
	"fmt"
	"time"

	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities/dtos"
)

type CreateVehicleRequest struct {
	Price        float64 `json:"price" validate:"required,min=0"`
	Brand        string  `json:"brand" validate:"required,max=30"`
	Model        string  `json:"model" validate:"required,max=30"`
	Year         int     `json:"year" validate:"required,min=1900"`
	BodyType     string  `json:"body_type" validate:"required,max=30"`
	Transmission string  `json:"transmission" validate:"required,max=30"`
	FuelType     string  `json:"fuel_type" validate:"required,max=30"`
	Color        string  `json:"color" validate:"required,max=30"`
	Mileage      int     `json:"mileage" validate:"min=0"`
	Engine       string  `json:"engine" validate:"required"`
	Doors        int     `json:"doors" validate:"required,min=2,max=5"`

	HasAirConditioning bool `json:"has_air_conditioning"`
	HasAirbag          bool `json:"has_airbag"`
	HasAbsBrakes       bool `json:"has_abs_brakes"`
	HasPowerSteering   bool `json:"has_power_steering"`
	HasPowerWindows    bool `json:"has_power_windows"`
	HasPowerLocks      bool `json:"has_power_locks"`
	HasMultimedia      bool `json:"has_multimedia"`
	HasAlarm           bool `json:"has_alarm"`
	HasTractionControl bool `json:"has_traction_control"`
	HasRearCamera      bool `json:"has_rear_camera"`
	HasParkingSensors  bool `json:"has_parking_sensors"`

	Condition   string `json:"condition" validate:"required,oneof=new used demonstration"`
	Description string `json:"description"`

	Status    *string    `json:"status" validate:"omitempty,oneof=available reserved sold maintenance"`
	CreatedAt *time.Time `json:"created_at" validate:"omitempty"`
}

func (req *CreateVehicleRequest) ToDomain() (*entities.Vehicle, error) {
	condition, err := dtos.NewCondition(req.Condition)
	if err != nil {
		return nil, fmt.Errorf("invalid condition: %w", err)
	}

	year, err := dtos.NewYear(req.Year)
	if err != nil {
		return nil, fmt.Errorf("invalid year: %w", err)
	}

	price, err := dtos.NewPrice(req.Price)
	if err != nil {
		return nil, fmt.Errorf("invalid price: %w", err)
	}

	bodyType, err := dtos.NewBodyType(req.BodyType)
	if err != nil {
		return nil, fmt.Errorf("invalid body type: %w", err)
	}

	transmission, err := dtos.NewTransmission(req.Transmission)
	if err != nil {
		return nil, fmt.Errorf("invalid transmission: %w", err)
	}

	fuelType, err := dtos.NewFuelType(req.FuelType)
	if err != nil {
		return nil, fmt.Errorf("invalid fuel type: %w", err)
	}

	mileage, err := dtos.NewMileage(req.Mileage)
	if err != nil {
		return nil, fmt.Errorf("invalid mileage: %w", err)
	}

	doors, err := dtos.NewDoors(req.Doors)
	if err != nil {
		return nil, fmt.Errorf("invalid door count: %w", err)
	}

	specification := entities.Specification{
		BodyType:     bodyType,
		Transmission: transmission,
		FuelType:     fuelType,
		Mileage:      mileage,
		Doors:        doors,
		Engine:       req.Engine,
	}

	features := entities.Features{
		HasAirConditioning: req.HasAirConditioning,
		HasAirbag:          req.HasAirbag,
		HasAbsBrakes:       req.HasAbsBrakes,
		HasPowerSteering:   req.HasPowerSteering,
		HasPowerWindows:    req.HasPowerWindows,
		HasPowerLocks:      req.HasPowerLocks,
		HasMultimedia:      req.HasMultimedia,
		HasAlarm:           req.HasAlarm,
		HasTractionControl: req.HasTractionControl,
		HasRearCamera:      req.HasRearCamera,
		HasParkingSensors:  req.HasParkingSensors,
	}
	vehicle := &entities.Vehicle{
		Brand:         req.Brand,
		Model:         req.Model,
		Color:         req.Color,
		Description:   req.Description,
		Year:          year,
		Price:         price,
		Features:      features,
		Condition:     condition,
		Specification: specification,
	}

	if req.Status != nil {
		status, errStatus := dtos.NewStatus(*req.Status)
		if errStatus != nil {
			return nil, fmt.Errorf("invalid status: %w", errStatus)
		}
		vehicle.Status = status
	} else {
		vehicle.Status = dtos.Available
	}

	return vehicle, nil
}
