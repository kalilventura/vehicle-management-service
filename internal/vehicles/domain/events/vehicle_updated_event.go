package events

import (
	"github.com/kalilventura/vehicle-management/internal/shared/domain"
)

// VehicleUpdatedEvent is raised when a vehicle is updated
type VehicleUpdatedEvent struct {
	domain.BaseDomainEvent
}

// NewVehicleUpdatedEvent creates a new VehicleUpdatedEvent
func NewVehicleUpdatedEvent(vehicleID string) VehicleUpdatedEvent {
	return VehicleUpdatedEvent{
		BaseDomainEvent: domain.NewBaseDomainEvent(vehicleID, "VehicleUpdatedEvent"),
	}
}

