package controllers

import (
	"net/http"

	shared "github.com/kalilventura/vehicle-management/internal/shared/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/shared/infrastructure/controllers"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/commands"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/vehicles/infrastructure/controllers/responses"
	"github.com/labstack/echo/v4"
	logger "github.com/sirupsen/logrus"
)

type GetVehicleByIdController struct {
	command commands.GetVehicleByID
}

func NewGetVehicleByIdController(command commands.GetVehicleByID) *GetVehicleByIdController {
	return &GetVehicleByIdController{command}
}

func (ctrl *GetVehicleByIdController) GetBind() shared.ControllerBind {
	return shared.ControllerBind{
		Method:       http.MethodGet,
		Version:      "v1",
		RelativePath: "/vehicles/:id",
	}
}

func (ctrl *GetVehicleByIdController) Execute(ectx echo.Context) error {
	id := ectx.Param("id")

	var handler error
	listeners := commands.GetVehicleByIDListeners{
		OnSuccess: func(vehicle *entities.Vehicle) {
			handler = ctrl.onSuccess(ectx, vehicle)
		},
		OnNotFound: func() {
			handler = ctrl.onNotFound(ectx)
		},
		OnInternalServerError: func(err error) {
			handler = ctrl.onError(ectx, err)
		},
	}
	ctrl.command.Execute(id, listeners)
	return handler
}

func (ctrl *GetVehicleByIdController) onSuccess(ectx echo.Context, vehicle *entities.Vehicle) error {
	response := controllers.NewSuccessResponse(http.StatusOK, responses.NewVehicleResponse(vehicle))
	return ectx.JSON(http.StatusOK, response)
}

func (ctrl *GetVehicleByIdController) onNotFound(ectx echo.Context) error {
	validationErrors := map[string]string{
		"message": "The requested vehicle was not found",
	}
	response := controllers.NewErrorResponse(
		http.StatusNotFound,
		validationErrors,
	)
	return ectx.JSON(http.StatusNotFound, response)
}

func (ctrl *GetVehicleByIdController) onError(ectx echo.Context, err error) error {
	logger.Errorf("Error occured %v", err)
	response := controllers.NewErrorResponse(
		http.StatusInternalServerError,
		nil,
	)
	return ectx.JSON(http.StatusInternalServerError, response)
}
