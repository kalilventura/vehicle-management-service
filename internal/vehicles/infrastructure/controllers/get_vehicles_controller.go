package controllers

import (
	"github.com/kalilventura/vehicle-management/internal/shared/domain/entities"
	"github.com/labstack/echo/v4"
	"net/http"
)

type GetVehiclesController struct{}

func NewGetVehiclesController() *GetVehiclesController {
	return &GetVehiclesController{}
}

func (ctrl *GetVehiclesController) GetBind() entities.ControllerBind {
	return entities.ControllerBind{
		Method:       http.MethodGet,
		Version:      "v1",
		RelativePath: "/vehicles",
	}
}

func (ctrl *GetVehiclesController) Execute(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Hello World")
}
