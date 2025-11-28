package events

import (
	"github.com/kalilventura/vehicle-management/internal/shared/domain"
)

// VehicleCreatedEvent is raised when a vehicle is created
type VehicleCreatedEvent struct {
	domain.BaseDomainEvent
	Brand string
	Model string
	Price float64
}

// NewVehicleCreatedEvent creates a new VehicleCreatedEvent
func NewVehicleCreatedEvent(vehicleID, brand, model string, price float64) VehicleCreatedEvent {
	return VehicleCreatedEvent{
		BaseDomainEvent: domain.NewBaseDomainEvent(vehicleID, "VehicleCreatedEvent"),
		Brand:           brand,
		Model:           model,
		Price:           price,
	}
}

