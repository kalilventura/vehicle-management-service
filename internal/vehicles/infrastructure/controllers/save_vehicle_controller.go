package controllers

import (
	"net/http"

	shared "github.com/kalilventura/vehicle-management/internal/shared/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/shared/infrastructure/controllers"
	"github.com/kalilventura/vehicle-management/internal/shared/infrastructure/controllers/helpers"
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

// Execute handles the creation of a new vehicle record.
//
// CreateVehicle handles the creation of a new vehicle record.
//
// @Summary      Create a New Vehicle
// @Description  Adds a new vehicle to the database. The request body must contain all required vehicle details. Upon successful creation, the full vehicle object, including its server-generated unique ID, is returned.
// @ID           create-vehicle
// @Tags         vehicles
// @Accept       json
// @Produce      json
// @Param        vehicle body      requests.CreateVehicleRequest  true  "Payload containing the new vehicle's data"
// @Success      201     {object}  controllers.SuccessResponse{data=responses.VehicleResponse} "Vehicle created successfully"
// @Failure      400     {object}  controllers.ErrorResponse "Bad Request (e.g., missing required fields or invalid data format)"
// @Failure      409     {object}  controllers.ErrorResponse "Conflict (e.g., a vehicle with the same license plate already exists)"
// @Failure      500     {object}  controllers.ErrorResponse "Internal Server Error"
// @Router       /v1/vehicles [post]
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
	validationErrors := helpers.ExtractValidationErrors(err)
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
