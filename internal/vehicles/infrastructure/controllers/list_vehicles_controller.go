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

type ListVehiclesController struct {
	command commands.ListVehicles
}

func NewListVehiclesController(command commands.ListVehicles) *ListVehiclesController {
	return &ListVehiclesController{command}
}

func (ctrl *ListVehiclesController) GetBind() shared.ControllerBind {
	return shared.ControllerBind{
		Method:       http.MethodGet,
		Version:      "v1",
		RelativePath: "/vehicles",
	}
}

// Execute
// ListVehicles handles the retrieval and filtering of a paginated list of vehicles.
//
// @Summary      List and Filter Vehicles
// @Description  Retrieves a paginated list of vehicles. This endpoint supports filtering by brand, model, and status, as well as sorting and pagination.
// @ID           list-vehicles
// @Tags         vehicles
// @Produce      json
// @Param        brand     query     string  false  "Filter by vehicle brand (e.g., 'Ford')"
// @Param        model     query     string  false  "Filter by vehicle model (e.g., 'Mustang')"
// @Param        status    query     string  false  "Filter by vehicle status" enums(available, sold)
// @Param        sortBy    query     string  false  "Field to sort by" enums(price, year, createdAt) default(createdAt)
// @Param        sortOrder query     string  false  "Sort order ('asc' or 'desc')" enums(asc, desc) default(desc)
// @Param        page      query     int     false  "Page number for pagination" default(1)
// @Param        pageSize  query     int     false  "Number of items per page" default(10)
// @Success      200       {object}  controllers.PaginatedResponse[responses.VehicleViewResponse] "A paginated list of vehicles that match the criteria"
// @Failure      400       {object}  controllers.ErrorResponse "Bad Request (e.g., invalid filter or pagination parameters)"
// @Failure      500       {object}  controllers.ErrorResponse "Internal Server Error"
// @Router       /v1/vehicles [get]
func (ctrl *ListVehiclesController) Execute(ectx echo.Context) error {
	searchParams, err := ctrl.GetQueryParams(ectx)
	if err != nil {
		return ectx.JSON(http.StatusBadRequest, "bad request")
	}

	entity, err := searchParams.ToDomain()
	if err != nil {
		return ectx.JSON(http.StatusBadRequest, err)
	}

	var handler error
	listeners := commands.ListVehiclesListeners{
		OnSuccess: func(vehicles *shared.PaginatedEntity[entities.Vehicle]) {
			handler = ctrl.onSuccess(ectx, vehicles)
		},
		OnInternalServerError: func(err error) {
			handler = ctrl.onInternalServerError(ectx, err)
		},
	}
	ctrl.command.Execute(*entity, listeners)
	return handler
}

func (ctrl *ListVehiclesController) GetQueryParams(ectx echo.Context) (*requests.ListVehiclesQueryParams, error) {
	searchParams := &requests.ListVehiclesQueryParams{}
	if err := ectx.Bind(searchParams); err != nil {
		return nil, err
	}
	if searchParams.Page == 0 {
		searchParams.Page = 1
	}
	if searchParams.Size == 0 {
		searchParams.Size = 10
	}
	return searchParams, nil
}

func (ctrl *ListVehiclesController) onSuccess(
	ectx echo.Context, vehicles *shared.PaginatedEntity[entities.Vehicle]) error {
	var responseList []responses.VehicleViewResponse
	for _, vehicle := range vehicles.Content {
		entry := responses.NewVehicleViewResponse(vehicle)
		responseList = append(responseList, entry)
	}

	response := controllers.NewPaginatedResponse(responseList, vehicles.Pagination)
	return ectx.JSON(http.StatusOK, response)
}

func (ctrl *ListVehiclesController) onInternalServerError(ectx echo.Context, err error) error {
	logger.Errorf("Error occured %v", err)
	response := controllers.NewErrorResponse(
		http.StatusInternalServerError,
		nil,
	)
	return ectx.JSON(http.StatusInternalServerError, response)
}
