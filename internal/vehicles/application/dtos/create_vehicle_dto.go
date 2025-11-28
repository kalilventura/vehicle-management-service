package dtos

// CreateVehicleDTO represents the create vehicle request
type CreateVehicleDTO struct {
	Brand         string
	Model         string
	Color         string
	Description   string
	Price         float64
	BodyType      string
	Transmission  string
	FuelType      string
	Mileage       int
	Doors         int
	Engine        string
	Year          int
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
}

