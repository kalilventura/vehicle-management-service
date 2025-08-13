//go:build unit

package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kalilventura/vehicle-management/internal/vehicles/infrastructure/controllers"
	"github.com/kalilventura/vehicle-management/test/vehicles/domain/builders"
	"github.com/kalilventura/vehicle-management/test/vehicles/domain/commands"
	builders2 "github.com/kalilventura/vehicle-management/test/vehicles/infrastructure/builders"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestSaveVehicleController(t *testing.T) {
	const route = "/v1/vehicles"
	t.Run("should return Created when the vehicle was created", func(t *testing.T) {
		// given
		requestBody := builders2.NewCreateVehicleRequestBuilder().
			WithValidDefaults().
			BuildRequest()
		vehicle := builders.NewVehicleBuilder().Build()
		command := commands.NewSaveVehicleCommandStub().WithOnSuccess(vehicle)
		controller := controllers.NewSaveVehicleController(command)

		router := echo.New()
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, route, requestBody)
		request.Header.Set("Content-Type", "application/json")
		router.POST(route, controller.Execute)

		// when
		router.ServeHTTP(recorder, request)

		// then
		assert.Equal(t, http.StatusCreated, recorder.Code)
	})

	t.Run("should return Bad Request when the vehicle is invalid", func(t *testing.T) {
		// given
		createRequest := builders2.NewCreateVehicleRequestBuilder().BuildRequest()
		command := commands.NewSaveVehicleCommandStub().WithOnNotValid()
		controller := controllers.NewSaveVehicleController(command)

		router := echo.New()
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, route, createRequest)
		request.Header.Set("Content-Type", "application/json")
		router.POST(route, controller.Execute)

		// when
		router.ServeHTTP(recorder, request)

		// then
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("should return Bad Request when the body is invalid", func(t *testing.T) {
		// given
		vehicle := builders.NewVehicleBuilder().Build()
		command := commands.NewSaveVehicleCommandStub().WithOnSuccess(vehicle)
		controller := controllers.NewSaveVehicleController(command)

		router := echo.New()
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, route, nil)
		request.Header.Set("Content-Type", "application/json")
		router.POST(route, controller.Execute)

		// when
		router.ServeHTTP(recorder, request)

		// then
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("should return Internal Server Error due an unexpected error", func(t *testing.T) {
		// given
		requestBody := builders2.NewCreateVehicleRequestBuilder().
			WithValidDefaults().
			BuildRequest()
		command := commands.NewSaveVehicleCommandStub().WithOnInternalServerError()
		controller := controllers.NewSaveVehicleController(command)

		router := echo.New()
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, route, requestBody)
		request.Header.Set("Content-Type", "application/json")
		router.POST(route, controller.Execute)

		// when
		router.ServeHTTP(recorder, request)

		// then
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})
}
