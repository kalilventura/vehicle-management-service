package services

import "github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"

type PaymentsService interface {
	Pay(sellRequest *entities.SellVehicle, listeners PaymentsServiceListeners)
}

type PaymentsServiceListeners struct {
	OnSuccess             func(sellRequest *entities.SellVehicle)
	OnBadRequest          func(err error)
	OnInternalServerError func(err error)
}
