package commands

import "github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"

type GetVehicleByID interface {
	Execute(ID string, listeners GetVehicleByIDListeners)
}

type GetVehicleByIDListeners struct {
	OnSuccess             func(vehicle *entities.Vehicle)
	OnNotFound            func()
	OnInternalServerError func(err error)
}
