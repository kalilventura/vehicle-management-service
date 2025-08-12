package controllers

import (
	"github.com/google/wire"
)

//nolint:gochecknoglobals // requirement for container
var Container = wire.NewSet(
	NewSaveVehicleController,
	NewGetVehicleByIdController,
	NewListVehiclesController,
	NewUpdateVehicleController,
)
