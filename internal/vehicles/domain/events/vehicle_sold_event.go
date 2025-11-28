package events

import (
	"github.com/kalilventura/vehicle-management/internal/shared/domain"
)

// VehicleSoldEvent is raised when a vehicle is sold
type VehicleSoldEvent struct {
	domain.BaseDomainEvent
	Price float64
}

// NewVehicleSoldEvent creates a new VehicleSoldEvent
func NewVehicleSoldEvent(vehicleID string, price float64) VehicleSoldEvent {
	return VehicleSoldEvent{
		BaseDomainEvent: domain.NewBaseDomainEvent(vehicleID, "VehicleSoldEvent"),
		Price:           price,
	}
}

