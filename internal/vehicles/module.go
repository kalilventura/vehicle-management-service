package vehicles

import (
	"github.com/kalilventura/vehicle-management/internal/shared/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/vehicles/infrastructure/controllers"
)

type Module struct {
	vehiclesControllers []entities.Controller
}

func NewModule(
	saveController *controllers.SaveVehicleController,
	getController *controllers.GetVehicleByIdController,
	listController *controllers.ListVehiclesController,
	updateController *controllers.UpdateVehicleController,
	sellController *controllers.SellVehicleController,
) *Module {
	vehiclesControllers := []entities.Controller{
		saveController,
		getController,
		listController,
		updateController,
		sellController,
	}
	return &Module{vehiclesControllers}
}

func (m *Module) GetControllers() []entities.Controller {
	return m.vehiclesControllers
}
