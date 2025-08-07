package commands

import "github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"

type SaveVehicle interface {
	Execute(vehicle *entities.Vehicle, listeners SaveVehicleListeners)
}

type SaveVehicleListeners struct {
	OnSuccess             func(vehicle *entities.Vehicle)
	OnNotValid            func(err error)
	OnInternalServerError func(err error)
}
