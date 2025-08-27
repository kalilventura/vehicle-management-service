package main

import (
	"fmt"

	_ "github.com/kalilventura/vehicle-management/cmd/docs" // generated docs
	"github.com/kalilventura/vehicle-management/internal/shared/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/shared/domain/services"
	"github.com/labstack/echo/v4"
	logger "github.com/sirupsen/logrus"
	"github.com/swaggo/echo-swagger" // echo-swagger middleware
)

type App struct {
	migrationService services.MigrationService
	settings         *entities.Settings
	modules          []entities.HTTPModule
}

func NewApp(
	migrationService services.MigrationService,
	settings *entities.Settings,
	modules []entities.HTTPModule,
) *App {
	return &App{
		migrationService,
		settings,
		modules,
	}
}

func (a *App) SetupServer() *echo.Echo {
	application := echo.New()
	handleRoutes(application, a.modules)
	handleSwagger(application)
	return application
}

func (a *App) RunMigrations() {
	logger.Infof("Creating the database structure...")
	err := a.migrationService.Run("db/migrations")
	if err != nil {
		logger.Fatalf("failed to run migrations: %v", err)
		panic(err)
	}
	logger.Infof("Database structure created successfully")
}

func handleRoutes(application *echo.Echo, modules []entities.HTTPModule) {
	logger.Info("🚀 Initializing HTTP router and registering routes...")
	groups := map[string]*echo.Group{}

	for _, module := range modules {
		for _, controller := range module.GetControllers() {
			bind := controller.GetBind()

			group, found := groups[bind.Version]
			if !found {
				group = application.Group(bind.Version)
				groups[bind.Version] = group
			}
			message := fmt.Sprintf("router registered: [%s] %s", bind.Method, bind.GetFullPath())
			logger.Info(message)
			group.Add(bind.Method, bind.RelativePath, controller.Execute)
		}
	}
}

func handleSwagger(application *echo.Echo) {
	logger.Info("registering swagger")
	application.GET("/swagger/*", echoSwagger.WrapHandler)
}
