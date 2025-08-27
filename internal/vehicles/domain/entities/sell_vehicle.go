package entities

type SellVehicle struct {
	VehicleID string  `json:"vehicle_id"`
	Cpf       string  `json:"cpf"`
	Amount    float64 `json:"amount"`
}
