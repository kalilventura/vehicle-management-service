package responses

import "time"

// VehicleResponse
// @Description Represents a vehicle
type VehicleResponse struct {
	ID                 string    `json:"id,omitempty"`
	Brand              string    `json:"brand,omitempty"`
	Model              string    `json:"model,omitempty"`
	Color              string    `json:"color,omitempty"`
	Description        string    `json:"description,omitempty"`
	Price              float64   `json:"price,omitempty"`
	BodyType           string    `json:"bodyType,omitempty"`
	Transmission       string    `json:"transmission,omitempty"`
	FuelType           string    `json:"fuelType,omitempty"`
	Mileage            int       `json:"mileage,omitempty"`
	Doors              int       `json:"doors,omitempty"`
	Engine             string    `json:"engine,omitempty"`
	Year               int       `json:"year,omitempty"`
	Status             string    `json:"status,omitempty"`
	Condition          string    `json:"condition,omitempty"`
	HasAirConditioning bool     `json:"hasAirConditioning,omitempty"`
	HasAirbag          bool     `json:"hasAirbag,omitempty"`
	HasAbsBrakes       bool     `json:"hasAbsBrakes,omitempty"`
	HasPowerSteering   bool     `json:"hasPowerSteering,omitempty"`
	HasPowerWindows    bool     `json:"hasPowerWindows,omitempty"`
	HasPowerLocks      bool     `json:"hasPowerLocks,omitempty"`
	HasMultimedia      bool     `json:"hasMultimedia,omitempty"`
	HasAlarm           bool     `json:"hasAlarm,omitempty"`
	HasTractionControl bool     `json:"hasTractionControl,omitempty"`
	HasRearCamera      bool     `json:"hasRearCamera,omitempty"`
	HasParkingSensors  bool     `json:"hasParkingSensors,omitempty"`
	CreatedAt          time.Time `json:"createdAt,omitempty"`
	UpdatedAt          time.Time `json:"updatedAt,omitempty"`
} // @name VehicleResponse

