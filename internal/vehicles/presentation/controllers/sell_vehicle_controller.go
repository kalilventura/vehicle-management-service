package controllers

import (
	"net/http"

	shared "github.com/kalilventura/vehicle-management/internal/shared/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/shared/infrastructure/controllers"
	"github.com/kalilventura/vehicle-management/internal/vehicles/application/use-cases/sell-vehicle"
	"github.com/kalilventura/vehicle-management/internal/vehicles/presentation/controllers/requests"
	"github.com/kalilventura/vehicle-management/internal/vehicles/presentation/filters"
	"github.com/labstack/echo/v4"
)

type SellVehicleController struct {
	service        *sellvehicle.SellVehicleService
	exceptionFilter *filters.VehicleExceptionFilter
}

func NewSellVehicleController(
	service *sellvehicle.SellVehicleService,
	exceptionFilter *filters.VehicleExceptionFilter,
) *SellVehicleController {
	return &SellVehicleController{
		service:        service,
		exceptionFilter: exceptionFilter,
	}
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
	vehicleRequest := new(requests.SellVehicleRequest)
	if err := ectx.Bind(vehicleRequest); err != nil {
		return ctrl.exceptionFilter.HandleError(ectx, err)
	}

	vehicleID := ectx.Param("id")
	entity := vehicleRequest.ToDomain(vehicleID)

	// Convert to application DTO
	sellDTO := sellvehicle.SellVehicleDTO{
		VehicleID: entity.VehicleID,
		CPF:       entity.Cpf,
		Amount:    entity.Amount,
	}

	// Execute use case
	err := ctrl.service.Execute(sellDTO)
	if err != nil {
		return ctrl.exceptionFilter.HandleError(ectx, err)
	}

	// Return success response
	response := controllers.NewSuccessResponse(http.StatusCreated, map[string]string{
		"vehicle_id": vehicleID,
		"status":     "sold",
	})
	ectx.Response().Header().Set("X-Resource-ID", vehicleID)

	return ectx.JSON(http.StatusCreated, response)
}
