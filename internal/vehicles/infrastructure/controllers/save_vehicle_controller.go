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

type SaveVehicleController struct {
	command commands.SaveVehicle
}

func NewSaveVehicleController(command commands.SaveVehicle) *SaveVehicleController {
	return &SaveVehicleController{command}
}

func (ctrl *SaveVehicleController) GetBind() shared.ControllerBind {
	return shared.ControllerBind{
		Method:       http.MethodPost,
		Version:      "v1",
		RelativePath: "/vehicles",
	}
}

// Execute godoc
// @Summary Save a new vehicle
// @Description Save a new vehicle
// @BasePath /v1/vehicles
// @Tags vehicles
// @Accept application/json
// @Produce application/json
// @Param request body requests.CreateVehicleRequest true "Request body"
// @Success 200 {object} controllers.SuccessResponse
// @Failure 400 {object} controllers.ErrorResponse
// @Failure 500 {object} controllers.ErrorResponse
// @Router /v1/vehicles [post]
func (ctrl *SaveVehicleController) Execute(ectx echo.Context) error {
	vehicleRequest := new(requests.CreateVehicleRequest)
	if err := ectx.Bind(vehicleRequest); err != nil {
		return ctrl.onInvalid(ectx, err)
	}

	entity, domainErr := vehicleRequest.ToDomain()
	if domainErr != nil {
		return ctrl.onInvalid(ectx, domainErr)
	}

	var handler error
	listeners := commands.SaveVehicleListeners{
		OnSuccess: func(vehicle *entities.Vehicle) {
			handler = ctrl.onSuccess(ectx, vehicle)
		},
		OnNotValid: func(err error) {
			handler = ctrl.onInvalid(ectx, err)
		},
		OnInternalServerError: func(err error) {
			handler = ctrl.onError(ectx, err)
		},
	}
	ctrl.command.Execute(entity, listeners)
	return handler
}

func (ctrl *SaveVehicleController) onSuccess(ectx echo.Context, vehicle *entities.Vehicle) error {
	response := controllers.NewSuccessResponse(http.StatusCreated, responses.NewVehicleResponse(vehicle))

	ectx.Response().Header().Set(echo.HeaderLocation, "/v1/vehicles/"+vehicle.ID)
	ectx.Response().Header().Set("X-Resource-ID", vehicle.ID)

	return ectx.JSON(http.StatusCreated, response)
}

func (ctrl *SaveVehicleController) onInvalid(ectx echo.Context, err error) error {
	validationErrors := ctrl.extractValidationErrors(err)
	response := controllers.NewErrorResponse(
		http.StatusBadRequest,
		validationErrors,
	)
	return ectx.JSON(http.StatusBadRequest, response)
}

func (ctrl *SaveVehicleController) onError(ectx echo.Context, err error) error {
	logger.Errorf("Error occured %v", err)
	response := controllers.NewErrorResponse(
		http.StatusInternalServerError,
		nil,
	)
	return ectx.JSON(http.StatusInternalServerError, response)
}

func (ctrl *SaveVehicleController) extractValidationErrors(err error) map[string]string {
	return map[string]string{
		"generic": err.Error(),
	}
}
