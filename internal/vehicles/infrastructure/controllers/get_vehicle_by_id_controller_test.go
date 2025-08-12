//go:build unit

package controllers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/kalilventura/vehicle-management/internal/vehicles/infrastructure/controllers"
	"github.com/kalilventura/vehicle-management/test/vehicles/domain/builders"
	"github.com/kalilventura/vehicle-management/test/vehicles/domain/commands"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetVehicleByIDController(t *testing.T) {
	const route = "/v1/vehicles/%s"
	t.Run("should respond NotFound when the vehicle not exists", func(t *testing.T) {
		// given
		router := echo.New()
		endpoint := fmt.Sprintf(route, gofakeit.UUID())
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, endpoint, nil)

		command := commands.NewGetVehicleByIDCommandStub().WithOnNotFound()
		controller := controllers.NewGetVehicleByIdController(command)
		router.GET(endpoint, controller.Execute)

		// when
		router.ServeHTTP(recorder, request)

		// then
		assert.Equal(t, http.StatusNotFound, recorder.Code)
	})

	t.Run("should respond InternalServerError due an unexpected error", func(t *testing.T) {
		// given
		router := echo.New()
		endpoint := fmt.Sprintf(route, gofakeit.UUID())
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, endpoint, nil)

		command := commands.NewGetVehicleByIDCommandStub().WithOnInternalServerError()
		controller := controllers.NewGetVehicleByIdController(command)
		router.GET(endpoint, controller.Execute)

		// when
		router.ServeHTTP(recorder, request)

		// then
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("should respond OK when the vehicle exists", func(t *testing.T) {
		// given
		vehicle := builders.NewVehicleBuilder().Build()
		router := echo.New()
		endpoint := fmt.Sprintf(route, gofakeit.UUID())
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, endpoint, nil)

		command := commands.NewGetVehicleByIDCommandStub().WithOnSuccess(vehicle)
		controller := controllers.NewGetVehicleByIdController(command)
		router.GET(endpoint, controller.Execute)

		// when
		router.ServeHTTP(recorder, request)

		// then
		assert.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("should return the metadata", func(t *testing.T) {
		// given
		controller := controllers.NewGetVehicleByIdController(nil)

		// when
		metadata := controller.GetBind()

		// then
		assert.Equal(t, "/vehicles/:id", metadata.RelativePath)
		assert.Equal(t, http.MethodGet, metadata.Method)
	})
}
