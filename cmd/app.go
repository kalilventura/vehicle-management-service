package main

import (
  "fmt"

  _ "github.com/kalilventura/vehicle-management/cmd/docs" // generated docs
  "github.com/kalilventura/vehicle-management/internal/shared/domain/entities"
  "github.com/labstack/echo/v4"
  logger "github.com/sirupsen/logrus"
  "github.com/swaggo/echo-swagger" // echo-swagger middleware
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

func SetupServer(modules []entities.HTTPModule) *echo.Echo {
  application := echo.New()
  handleRoutes(application, modules)
  handleSwagger(application)
  return application
}
