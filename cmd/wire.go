//go:build !test && wireinject

package main

import (
	"os"
	"strconv"

	"github.com/google/wire"
	"github.com/kalilventura/vehicle-management/internal/shared/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/shared/infrastructure/configuration"
	"github.com/kalilventura/vehicle-management/internal/vehicles"
)

func InjectModules() []entities.HTTPModule {
	wire.Build(
		InjectSettings,
		injectDatabaseSettings,
		configuration.NewDatabaseClient,
		vehicles.Container,
		newModules,
	)
	return nil
}

func InjectSettings() *entities.Settings {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	paymentsAPI := os.Getenv("PAYMENTS_API")

	return entities.NewSettings(port, paymentsAPI)
}

func injectDatabaseSettings() *entities.DatabaseSettings {
	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbSSL := os.Getenv("DB_SSL")
	return entities.NewDatabaseSettings(
		host,
		name,
		port,
		user,
		password,
		dbSSL,
	)
}

func newModules(vehiclesModule *vehicles.Module) []entities.HTTPModule {
	return []entities.HTTPModule{
		vehiclesModule,
	}
}
