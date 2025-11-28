package controllers

import (
	"net/http"

	shared "github.com/kalilventura/vehicle-management/internal/shared/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/shared/infrastructure/controllers"
	"github.com/kalilventura/vehicle-management/internal/vehicles/application/use-cases/get-vehicle"
	"github.com/kalilventura/vehicle-management/internal/vehicles/presentation/filters"
	"github.com/kalilventura/vehicle-management/internal/vehicles/presentation/mappers"
	"github.com/labstack/echo/v4"
)

// GetVehicleController handles GET /vehicles/:id requests
type GetVehicleController struct {
	service      *getvehicle.GetVehicleService
	exceptionFilter *filters.VehicleExceptionFilter
	responseMapper  *mappers.VehicleResponseMapper
}

// NewGetVehicleController creates a new GetVehicleController
func NewGetVehicleController(
	service *getvehicle.GetVehicleService,
	exceptionFilter *filters.VehicleExceptionFilter,
	responseMapper *mappers.VehicleResponseMapper,
) *GetVehicleController {
	return &GetVehicleController{
		service:         service,
		exceptionFilter: exceptionFilter,
		responseMapper:  responseMapper,
	}
}

func (ctrl *GetVehicleController) GetBind() shared.ControllerBind {
	return shared.ControllerBind{
		Method:       http.MethodGet,
		Version:      "v1",
		RelativePath: "/vehicles/:id",
	}
}

// Execute handles the get vehicle by ID request
//
// @Summary      Retrieve a Vehicle by ID
// @Description  Fetches the details of a single vehicle from the database using its unique UUID.
// @ID           get-vehicle-by-id
// @Tags         vehicles
// @Produce      json
// @Param        id   path      string  true  "The unique identifier (UUID) of the vehicle" format(uuid)
// @Success      200  {object}  controllers.SuccessResponse{data=responses.VehicleResponse} "Successfully retrieved the vehicle data"
// @Failure      404  {object}  controllers.ErrorResponse "The vehicle with the specified ID was not found"
// @Failure      500  {object}  controllers.ErrorResponse "Internal Server Error"
// @Router       /v1/vehicles/{id} [get]
func (ctrl *GetVehicleController) Execute(ectx echo.Context) error {
	vehicleID := ectx.Param("id")

	// Execute use case
	responseDTO, err := ctrl.service.Execute(vehicleID)
	if err != nil {
		// Let exception filter handle the error and return appropriate status code
		return ctrl.exceptionFilter.HandleError(ectx, err)
	}

	// Convert DTO to response format
	response := ctrl.responseMapper.ToResponse(responseDTO)

	// Return success response
	successResponse := controllers.NewSuccessResponse(http.StatusOK, response)
	return ectx.JSON(http.StatusOK, successResponse)
}

