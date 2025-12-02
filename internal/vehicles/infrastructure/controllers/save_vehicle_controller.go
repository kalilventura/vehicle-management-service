package controllers

import (
	"net/http"

	shared "github.com/kalilventura/vehicle-management/internal/shared/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/shared/infrastructure/controllers"
	"github.com/kalilventura/vehicle-management/internal/vehicles/application/dtos"
	"github.com/kalilventura/vehicle-management/internal/vehicles/application/use-cases/create-vehicle"
	"github.com/kalilventura/vehicle-management/internal/vehicles/infrastructure/controllers/requests"
	"github.com/kalilventura/vehicle-management/internal/vehicles/infrastructure/controllers/responses"
	"github.com/kalilventura/vehicle-management/internal/vehicles/presentation/filters"
	"github.com/kalilventura/vehicle-management/internal/vehicles/presentation/mappers"
	"github.com/labstack/echo/v4"
)

type SaveVehicleController struct {
	service         *createvehicle.CreateVehicleService
	exceptionFilter *filters.VehicleExceptionFilter
	responseMapper  *mappers.VehicleResponseMapper
}

func NewSaveVehicleController(
	service *createvehicle.CreateVehicleService,
	exceptionFilter *filters.VehicleExceptionFilter,
	responseMapper *mappers.VehicleResponseMapper,
) *SaveVehicleController {
	return &SaveVehicleController{
		service:         service,
		exceptionFilter: exceptionFilter,
		responseMapper:  responseMapper,
	}
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
		return ctrl.exceptionFilter.HandleError(ectx, err)
	}

	// Convert request to application DTO
	createDTO := dtos.CreateVehicleDTO{
		Brand:              vehicleRequest.Brand,
		Model:              vehicleRequest.Model,
		Color:              vehicleRequest.Color,
		Description:        vehicleRequest.Description,
		Price:              vehicleRequest.Price,
		BodyType:           vehicleRequest.BodyType,
		Transmission:       vehicleRequest.Transmission,
		FuelType:           vehicleRequest.FuelType,
		Mileage:            vehicleRequest.Mileage,
		Doors:              vehicleRequest.Doors,
		Engine:             vehicleRequest.Engine,
		Year:               vehicleRequest.Year,
		Condition:          vehicleRequest.Condition,
		HasAirConditioning: vehicleRequest.HasAirConditioning,
		HasAirbag:          vehicleRequest.HasAirbag,
		HasAbsBrakes:       vehicleRequest.HasAbsBrakes,
		HasPowerSteering:   vehicleRequest.HasPowerSteering,
		HasPowerWindows:    vehicleRequest.HasPowerWindows,
		HasPowerLocks:      vehicleRequest.HasPowerLocks,
		HasMultimedia:      vehicleRequest.HasMultimedia,
		HasAlarm:           vehicleRequest.HasAlarm,
		HasTractionControl: vehicleRequest.HasTractionControl,
		HasRearCamera:      vehicleRequest.HasRearCamera,
		HasParkingSensors:  vehicleRequest.HasParkingSensors,
	}

	// Execute use case
	responseDTO, err := ctrl.service.Execute(createDTO)
	if err != nil {
		return ctrl.exceptionFilter.HandleError(ectx, err)
	}

	// Convert DTO to response format
	response := ctrl.responseMapper.ToResponse(responseDTO)

	// Set headers
	ectx.Response().Header().Set(echo.HeaderLocation, "/v1/vehicles/"+responseDTO.ID)
	ectx.Response().Header().Set("X-Resource-ID", responseDTO.ID)

	// Return success response
	successResponse := controllers.NewSuccessResponse(http.StatusCreated, response)
	return ectx.JSON(http.StatusCreated, successResponse)
}
