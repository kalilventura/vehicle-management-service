package controllers

import (
  "github.com/kalilventura/vehicle-management/internal/shared/domain/entities"
  "github.com/kalilventura/vehicle-management/internal/vehicles/domain/commands"
  "github.com/labstack/echo/v4"
  "net/http"
)

type SaveVehiclesController struct {
  command commands.SaveVehicle
}

func NewSaveVehiclesController(command commands.SaveVehicle) *SaveVehiclesController {
  return &SaveVehiclesController{command}
}

func (ctrl *SaveVehiclesController) GetBind() entities.ControllerBind {
  return entities.ControllerBind{
    Method:       http.MethodPost,
    Version:      "v1",
    RelativePath: "/vehicles",
  }
}

func (ctrl *SaveVehiclesController) Execute(ctx echo.Context) error {
  return ctx.JSON(http.StatusOK, "Hello World")
}
