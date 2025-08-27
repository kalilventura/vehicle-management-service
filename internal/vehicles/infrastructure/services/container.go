package services

import (
	"github.com/google/wire"

	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/services"
)

//nolint:gochecknoglobals // requirement for container
var Container = wire.NewSet(
	NewPaymentsService,
	wire.Bind(new(services.PaymentsService), new(*PaymentsService)),
)
