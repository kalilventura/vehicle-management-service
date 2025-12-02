package requests

// CreateVehicleRequest
// @Description Object that represents a vehicle
type CreateVehicleRequest struct {
	Brand              string  `json:"brand" binding:"required"`
	Model              string  `json:"model" binding:"required"`
	Color              string  `json:"color" binding:"required"`
	Description        string  `json:"description"`
	Price              float64 `json:"price" binding:"required"`
	BodyType           string  `json:"bodyType" binding:"required"`
	Transmission       string  `json:"transmission" binding:"required"`
	FuelType           string  `json:"fuelType" binding:"required"`
	Mileage            int     `json:"mileage"`
	Doors              int     `json:"doors" binding:"required"`
	Engine             string  `json:"engine" binding:"required"`
	Year               int     `json:"year" binding:"required"`
	Condition          string  `json:"condition" binding:"required"`
	Status             *string `json:"status,omitempty"`
	HasAirConditioning bool    `json:"hasAirConditioning"`
	HasAirbag          bool    `json:"hasAirbag"`
	HasAbsBrakes       bool    `json:"hasAbsBrakes"`
	HasPowerSteering   bool    `json:"hasPowerSteering"`
	HasPowerWindows    bool    `json:"hasPowerWindows"`
	HasPowerLocks      bool    `json:"hasPowerLocks"`
	HasMultimedia      bool    `json:"hasMultimedia"`
	HasAlarm           bool    `json:"hasAlarm"`
	HasTractionControl bool    `json:"hasTractionControl"`
	HasRearCamera      bool    `json:"hasRearCamera"`
	HasParkingSensors  bool    `json:"hasParkingSensors"`
} // @name CreateVehicleRequest

