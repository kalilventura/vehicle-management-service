package dtos

import "time"

// VehicleResponseDTO represents the vehicle response
type VehicleResponseDTO struct {
	ID            string
	Brand         string
	Model         string
	Color         string
	Description   string
	Price         float64
	Currency      string
	BodyType      string
	Transmission  string
	FuelType      string
	Mileage       int
	Doors         int
	Engine        string
	Year          int
	Status        string
	Condition     string
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
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

