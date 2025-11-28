package exceptions

import "fmt"

// VehicleNotFoundException is raised when a vehicle is not found
type VehicleNotFoundException struct {
	VehicleID string
}

func (e VehicleNotFoundException) Error() string {
	return fmt.Sprintf("vehicle with ID %s not found", e.VehicleID)
}

// NewVehicleNotFoundException creates a new VehicleNotFoundException
func NewVehicleNotFoundException(vehicleID string) VehicleNotFoundException {
	return VehicleNotFoundException{VehicleID: vehicleID}
}

