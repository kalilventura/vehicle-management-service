package controllers

import (
	"net/http"

	shared "github.com/kalilventura/vehicle-management/internal/shared/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/shared/infrastructure/controllers"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/commands"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/vehicles/infrastructure/controllers/requests"
	"github.com/kalilventura/vehicle-management/internal/vehicles/infrastructure/controllers/responses"
	"github.com/labstack/echo/v4"
	logger "github.com/sirupsen/logrus"
)

type SaveVehiclesController struct {
	command commands.SaveVehicle
}

func NewSaveVehiclesController(command commands.SaveVehicle) *SaveVehiclesController {
	return &SaveVehiclesController{command}
}

func (ctrl *SaveVehiclesController) GetBind() shared.ControllerBind {
	return shared.ControllerBind{
		Method:       http.MethodPost,
		Version:      "v1",
		RelativePath: "/vehicles",
	}
}

func (ctrl *SaveVehiclesController) Execute(ectx echo.Context) error {
	vehicleRequest := new(requests.CreateVehicleRequest)
	if err := ectx.Bind(vehicleRequest); err != nil {
		return ctrl.onInvalid(ectx, err)
	}

	entity, domainErr := vehicleRequest.ToDomain()
	if domainErr != nil {
		return ctrl.onInvalid(ectx, domainErr)
	}

	var handlerErr error
	listeners := commands.SaveVehicleListeners{
		OnSuccess: func(vehicle *entities.Vehicle) {
			handlerErr = ctrl.onSuccess(ectx, vehicle)
		},
		OnNotValid: func(err error) {
			handlerErr = ctrl.onInvalid(ectx, err)
		},
		OnInternalServerError: func(err error) {
			handlerErr = ctrl.onError(ectx, err)
		},
	}
	ctrl.command.Execute(entity, listeners)
	return handlerErr
}

func (ctrl *SaveVehiclesController) onSuccess(ectx echo.Context, vehicle *entities.Vehicle) error {
	response := controllers.NewSuccessResponse(http.StatusCreated, responses.NewVehicleResponse(vehicle))

	ectx.Response().Header().Set(echo.HeaderLocation, "/v1/vehicles/"+vehicle.ID)
	ectx.Response().Header().Set("X-Resource-ID", vehicle.ID)

	return ectx.JSON(http.StatusCreated, response)
}

func (ctrl *SaveVehiclesController) onInvalid(ectx echo.Context, err error) error {
	validationErrors := ctrl.extractValidationErrors(err)
	response := controllers.NewErrorResponse(
		http.StatusBadRequest,
		validationErrors,
	)
	return ectx.JSON(http.StatusBadRequest, response)
}

func (ctrl *SaveVehiclesController) onError(ectx echo.Context, err error) error {
	logger.Errorf("Error occured %v", err)
	response := controllers.NewErrorResponse(
		http.StatusInternalServerError,
		nil,
	)
	return ectx.JSON(http.StatusInternalServerError, response)
}

func (ctrl *SaveVehiclesController) extractValidationErrors(err error) map[string]string {
	return map[string]string{
		"generic": err.Error(),
	}
}
