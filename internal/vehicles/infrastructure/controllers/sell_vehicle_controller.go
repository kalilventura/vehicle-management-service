package controllers

import (
	"net/http"

	shared "github.com/kalilventura/vehicle-management/internal/shared/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/shared/infrastructure/controllers"
	"github.com/kalilventura/vehicle-management/internal/shared/infrastructure/controllers/helpers"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/commands"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/vehicles/infrastructure/controllers/requests"
	"github.com/labstack/echo/v4"
	logger "github.com/sirupsen/logrus"
)

type SellVehicleController struct {
	command commands.SellVehicle
}

func NewSellVehicleController(command commands.SellVehicle) *SellVehicleController {
	return &SellVehicleController{command}
}

func (ctrl *SellVehicleController) GetBind() shared.ControllerBind {
	return shared.ControllerBind{
		Method:       http.MethodPost,
		Version:      "v1",
		RelativePath: "/vehicles/:id/sales",
	}
}

// Execute
//
// SellVehicle registers the sale of a specific vehicle.
//
// @Summary      Sell a Vehicle
// @Description  Processes and records the sale of an existing vehicle by its unique ID.
// @Description  Upon successful execution, this changes the vehicle's status to "sold".
// @Description  This operation is not idempotent; attempting to sell the same vehicle more than once will result in a conflict error.
// @ID           sell-vehicle-by-id
// @Tags         vehicles
// @Accept       json
// @Produce      json
// @Param        id   path      string                       true  "Vehicle ID (UUID)" format(uuid)
// @Param        sale body      requests.SellVehicleRequest  true  "Payload with sale details, such as buyer information and final price"
// @Success      201  {object}  controllers.SuccessResponse{data=object} "Vehicle sale registered successfully"
// @Failure      400  {object}  controllers.ErrorResponse "Bad Request (e.g., missing sale data or invalid format)"
// @Failure      404  {object}  controllers.ErrorResponse "Vehicle with the specified ID was not found"
// @Failure      409  {object}  controllers.ErrorResponse "Conflict: The vehicle has already been sold"
// @Failure      500  {object}  controllers.ErrorResponse "Internal Server Error"
// @Router       /v1/vehicles/{id}/sales [post]
func (ctrl *SellVehicleController) Execute(ectx echo.Context) error {
	var handler error
	vehicleRequest := new(requests.SellVehicleRequest)
	if err := ectx.Bind(vehicleRequest); err != nil {
		return ctrl.onInvalid(ectx, err)
	}
	entity := vehicleRequest.ToDomain(ectx.Param("id"))

	listeners := commands.SellVehicleListeners{
		OnSuccess: func(sell *entities.SellVehicle) {
			handler = ctrl.onSuccess(ectx, sell)
		},
		OnBadRequest: func(err error) {
			handler = ctrl.onInvalid(ectx, err)
		},
		OnInternalServerError: func(err error) {
			handler = ctrl.onInternalServerError(ectx, err)
		},
	}
	ctrl.command.Execute(entity, listeners)
	return handler
}

func (ctrl *SellVehicleController) onSuccess(ectx echo.Context, sell *entities.SellVehicle) error {
	response := controllers.NewSuccessResponse(http.StatusCreated, sell)
	ectx.Response().Header().Set("X-Resource-ID", sell.VehicleID)

	return ectx.JSON(http.StatusCreated, response)
}

func (ctrl *SellVehicleController) onInvalid(ectx echo.Context, err error) error {
	validationErrors := helpers.ExtractValidationErrors(err)
	response := controllers.NewErrorResponse(
		http.StatusBadRequest,
		validationErrors,
	)
	return ectx.JSON(http.StatusBadRequest, response)
}

func (ctrl *SellVehicleController) onInternalServerError(ectx echo.Context, err error) error {
	logger.Errorf("Error occured %v", err)
	response := controllers.NewErrorResponse(
		http.StatusInternalServerError,
		nil,
	)
	return ectx.JSON(http.StatusInternalServerError, response)
}
