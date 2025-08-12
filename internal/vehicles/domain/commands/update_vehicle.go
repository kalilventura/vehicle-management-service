package commands

import "github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"

type UpdateVehicle interface {
	Execute(vehicle *entities.UpdateVehicleInput, listeners UpdateVehicleListeners)
}

type UpdateVehicleListeners struct {
	OnSuccess             func(vehicle *entities.UpdateVehicleInput)
	OnNotFound            func()
	OnInternalServerError func(err error)
}
