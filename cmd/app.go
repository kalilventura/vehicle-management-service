package main

import (
	"fmt"

	"github.com/kalilventura/vehicle-management/internal/shared/domain/entities"
	"github.com/labstack/echo/v4"
	logger "github.com/sirupsen/logrus"
)

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
			logger.Infof(fmt.Sprintf("router registered: [%s] %s", bind.Method, bind.GetFullPath()))
			group.Add(bind.Method, bind.RelativePath, controller.Execute)
		}
	}
}

func StartServer(modules []entities.HTTPModule, globalSettings *entities.Settings) {
	application := echo.New()
	handleRoutes(application, modules)

	port := globalSettings.GetPort()
	application.Logger.Fatal(application.Start(port))
}
