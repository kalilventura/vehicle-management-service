package commands

import "github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"

type SellVehicle interface {
	Execute(sell *entities.SellVehicle, listeners SellVehicleListeners)
}

type SellVehicleListeners struct {
	OnSuccess             func(sell *entities.SellVehicle)
	OnBadRequest          func(err error)
	OnInternalServerError func(err error)
}
