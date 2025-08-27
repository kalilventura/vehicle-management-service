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

type UpdateVehicleController struct {
	command commands.UpdateVehicle
}

func NewUpdateVehicleController(command commands.UpdateVehicle) *UpdateVehicleController {
	return &UpdateVehicleController{command}
}

func (ctrl *UpdateVehicleController) GetBind() shared.ControllerBind {
	return shared.ControllerBind{
		Method:       http.MethodPatch,
		Version:      "v1",
		RelativePath: "/vehicles/:id",
	}
}

// Execute handles the partial update of an existing vehicle's attributes.
//
// UpdateVehicle handles the partial update of an existing vehicle's attributes.
//
// @Summary      Update Vehicle Details (Partial)
// @Description  Partially updates the data for a specific vehicle. Only the fields provided in the JSON request body will be modified. All other fields will remain unchanged.
// @ID           update-vehicle-by-id
// @Tags         vehicles
// @Accept       json
// @Produce      json
// @Param        id      path      string                         true  "The unique identifier (UUID) of the vehicle to update" format(uuid)
// @Param        vehicle body      requests.UpdateVehicleRequest  true  "Payload with the vehicle fields to be updated"
// @Success      200     {object}  controllers.SuccessResponse{data=responses.VehicleResponse} "Vehicle updated successfully"
// @Failure      400     {object}  controllers.ErrorResponse "Bad Request (e.g., invalid data format or validation error)"
// @Failure      404     {object}  controllers.ErrorResponse "The vehicle with the specified ID was not found"
// @Failure      409     {object}  controllers.ErrorResponse "Conflict (e.g., updating a unique field to a value that already exists)"
// @Failure      500     {object}  controllers.ErrorResponse "Internal Server Error"
// @Router       /v1/vehicles/{id} [patch]
func (ctrl *UpdateVehicleController) Execute(ectx echo.Context) error {
	var handler error
	id := ectx.Param("id")

	request := new(requests.UpdateVehicleRequest)
	if err := ectx.Bind(request); err != nil {
		return ctrl.onInvalid(ectx, err)
	}

	domain, domainErr := request.ToDomain(id)
	if domainErr != nil {
		return ctrl.onInvalid(ectx, domainErr)
	}

	listeners := commands.UpdateVehicleListeners{
		OnSuccess: func(vehicle *entities.UpdateVehicleInput) {
			handler = ctrl.onSuccess(ectx, vehicle)
		},
		OnNotFound: func() {
			handler = ctrl.onNotFound(ectx)
		},
		OnInternalServerError: func(err error) {
			handler = ctrl.onError(ectx, err)
		},
	}
	ctrl.command.Execute(domain, listeners)
	return handler
}

func (ctrl *UpdateVehicleController) onSuccess(ectx echo.Context, vehicle *entities.UpdateVehicleInput) error {
	updateResp := responses.NewUpdateResponse(vehicle)
	response := controllers.NewSuccessResponse(http.StatusOK, updateResp)
	return ectx.JSON(http.StatusOK, response)
}

func (ctrl *UpdateVehicleController) onNotFound(ectx echo.Context) error {
	validationErrors := map[string]string{"message": "The requested vehicle was not found"}
	response := controllers.NewErrorResponse(
		http.StatusNotFound,
		validationErrors,
	)
	return ectx.JSON(http.StatusNotFound, response)
}

func (ctrl *UpdateVehicleController) onInvalid(ectx echo.Context, err error) error {
	validationErrors := helpers.ExtractValidationErrors(err)
	response := controllers.NewErrorResponse(
		http.StatusBadRequest,
		validationErrors,
	)
	return ectx.JSON(http.StatusBadRequest, response)
}

func (ctrl *UpdateVehicleController) onError(ectx echo.Context, err error) error {
	logger.Errorf("Error occured %v", err)
	response := controllers.NewErrorResponse(
		http.StatusInternalServerError,
		nil,
	)
	return ectx.JSON(http.StatusInternalServerError, response)
}
