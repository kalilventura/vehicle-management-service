package services

import (
	"errors"

	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/services"
)

var errPayments = errors.New("errPayments")

type PaymentsServiceStub struct {
	callback func(listeners services.PaymentsServiceListeners)
}

func NewPaymentsServiceStub() *PaymentsServiceStub {
	return &PaymentsServiceStub{}
}

func (s *PaymentsServiceStub) WithOnBadRequest() *PaymentsServiceStub {
	s.callback = func(listeners services.PaymentsServiceListeners) {
		listeners.OnBadRequest(errPayments)
	}
	return s
}

func (s *PaymentsServiceStub) WithOnInternalServerError() *PaymentsServiceStub {
	s.callback = func(listeners services.PaymentsServiceListeners) {
		listeners.OnInternalServerError(errPayments)
	}
	return s
}

func (s *PaymentsServiceStub) WithOnSuccess(entity *entities.SellVehicle) *PaymentsServiceStub {
	s.callback = func(listeners services.PaymentsServiceListeners) {
		listeners.OnSuccess(entity)
	}
	return s
}

func (s *PaymentsServiceStub) Pay(_ *entities.SellVehicle, listeners services.PaymentsServiceListeners) {
	s.callback(listeners)
}
