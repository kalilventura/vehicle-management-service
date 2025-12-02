package controllers

import (
	"net/http"

	shared "github.com/kalilventura/vehicle-management/internal/shared/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/shared/infrastructure/controllers"
	"github.com/kalilventura/vehicle-management/internal/vehicles/application/use-cases/update-vehicle"
	"github.com/kalilventura/vehicle-management/internal/vehicles/presentation/controllers/requests"
	"github.com/kalilventura/vehicle-management/internal/vehicles/presentation/filters"
	"github.com/kalilventura/vehicle-management/internal/vehicles/presentation/mappers"
	"github.com/labstack/echo/v4"
)

type UpdateVehicleController struct {
	service         *updatevehicle.UpdateVehicleService
	exceptionFilter *filters.VehicleExceptionFilter
	responseMapper  *mappers.VehicleResponseMapper
}

func NewUpdateVehicleController(
	service *updatevehicle.UpdateVehicleService,
	exceptionFilter *filters.VehicleExceptionFilter,
	responseMapper *mappers.VehicleResponseMapper,
) *UpdateVehicleController {
	return &UpdateVehicleController{
		service:         service,
		exceptionFilter: exceptionFilter,
		responseMapper:  responseMapper,
	}
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
	id := ectx.Param("id")

	request := new(requests.UpdateVehicleRequest)
	if err := ectx.Bind(request); err != nil {
		return ctrl.exceptionFilter.HandleError(ectx, err)
	}

	domain, domainErr := request.ToDomain(id)
	if domainErr != nil {
		return ctrl.exceptionFilter.HandleError(ectx, domainErr)
	}

	// Execute use case
	responseDTO, err := ctrl.service.Execute(domain)
	if err != nil {
		return ctrl.exceptionFilter.HandleError(ectx, err)
	}

	// Convert DTO to response format
	response := ctrl.responseMapper.ToResponse(responseDTO)

	// Return success response
	successResponse := controllers.NewSuccessResponse(http.StatusOK, response)
	return ectx.JSON(http.StatusOK, successResponse)
}
