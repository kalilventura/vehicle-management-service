package exceptions

import "fmt"

// InvalidVehicleStatusException is raised when a vehicle operation is invalid for the current status
type InvalidVehicleStatusException struct {
	VehicleID    string
	CurrentStatus string
	Operation    string
}

func (e InvalidVehicleStatusException) Error() string {
	return fmt.Sprintf("cannot %s vehicle %s: current status is %s", e.Operation, e.VehicleID, e.CurrentStatus)
}

// NewInvalidVehicleStatusException creates a new InvalidVehicleStatusException
func NewInvalidVehicleStatusException(vehicleID, currentStatus, operation string) InvalidVehicleStatusException {
	return InvalidVehicleStatusException{
		VehicleID:     vehicleID,
		CurrentStatus: currentStatus,
		Operation:     operation,
	}
}

