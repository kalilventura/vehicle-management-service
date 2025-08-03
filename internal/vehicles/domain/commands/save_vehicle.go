package commands

import "github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"

type SaveVehicle interface {
	Execute(vehicle *entities.Vehicle)
}
