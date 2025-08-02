//go:build !test && wireinject

package main

import (
	"github.com/google/wire"
	"github.com/kalilventura/vehicle-management/internal/shared/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/vehicles"
	"os"
	"strconv"
)

func injectModules() []entities.HTTPModule {
	wire.Build(
		vehicles.Container,
		newModules,
	)
	return nil
}

func injectSettings() *entities.Settings {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	return entities.NewSettings(port)
}

func newModules(vehiclesModule *vehicles.Module) []entities.HTTPModule {
	return []entities.HTTPModule{
		vehiclesModule,
	}
}
