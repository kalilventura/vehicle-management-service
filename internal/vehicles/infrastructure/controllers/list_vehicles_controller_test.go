//go:build unit

package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kalilventura/vehicle-management/internal/vehicles/infrastructure/controllers"
	"github.com/kalilventura/vehicle-management/test/vehicles/domain/builders"
	"github.com/kalilventura/vehicle-management/test/vehicles/domain/commands"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestListVehiclesController(t *testing.T) {
	const route = "/v1/vehicles"
	t.Run("should return OK when there are vehicles to show", func(t *testing.T) {
		// given
		page := builders.NewVehicleBuilder().BuildPagination()
		command := commands.NewListVehiclesCommandStub().WithOnSuccess(page)
		controller := controllers.NewListVehiclesController(command)

		router := echo.New()
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, route, nil)
		router.GET(route, controller.Execute)

		// when
		router.ServeHTTP(recorder, request)

		// then
		assert.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("should return InternalServerError due an unexpected error", func(t *testing.T) {
		// given
		command := commands.NewListVehiclesCommandStub().WithOnInternalServerError()
		controller := controllers.NewListVehiclesController(command)

		router := echo.New()
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, route, nil)
		router.GET(route, controller.Execute)

		// when
		router.ServeHTTP(recorder, request)

		// then
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})
}
