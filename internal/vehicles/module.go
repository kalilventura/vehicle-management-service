package vehicles

import (
	"github.com/kalilventura/vehicle-management/internal/shared/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/vehicles/infrastructure/controllers"
)

type Module struct {
	vehiclesControllers []entities.Controller
}

func NewModule(listController *controllers.SaveVehiclesController) *Module {
	vehiclesControllers := []entities.Controller{
		listController,
	}
	return &Module{vehiclesControllers}
}

func (m *Module) GetControllers() []entities.Controller {
	return m.vehiclesControllers
}
