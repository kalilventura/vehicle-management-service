package controllers

import (
	"net/http"

	shared "github.com/kalilventura/vehicle-management/internal/shared/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/shared/infrastructure/controllers"
	"github.com/kalilventura/vehicle-management/internal/vehicles/application/use-cases/list-vehicles"
	"github.com/kalilventura/vehicle-management/internal/vehicles/presentation/controllers/requests"
	"github.com/kalilventura/vehicle-management/internal/vehicles/presentation/controllers/responses"
	"github.com/kalilventura/vehicle-management/internal/vehicles/presentation/filters"
	"github.com/labstack/echo/v4"
)

type ListVehiclesController struct {
	service         *listvehicles.ListVehiclesService
	exceptionFilter *filters.VehicleExceptionFilter
}

func NewListVehiclesController(
	service *listvehicles.ListVehiclesService,
	exceptionFilter *filters.VehicleExceptionFilter,
) *ListVehiclesController {
	return &ListVehiclesController{
		service:         service,
		exceptionFilter: exceptionFilter,
	}
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
		return ctrl.exceptionFilter.HandleError(ectx, err)
	}

	entity, err := searchParams.ToDomain()
	if err != nil {
		return ctrl.exceptionFilter.HandleError(ectx, err)
	}

	// Execute use case
	result, err := ctrl.service.Execute(*entity)
	if err != nil {
		return ctrl.exceptionFilter.HandleError(ectx, err)
	}

	// Convert DTOs to response format
	var responseList []responses.VehicleViewResponse
	for _, dto := range result.Content {
		responseList = append(responseList, responses.NewVehicleViewResponseFromDTO(dto))
	}

	// Return paginated response
	response := controllers.NewPaginatedResponse(responseList, result.Pagination)
	return ectx.JSON(http.StatusOK, response)
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
