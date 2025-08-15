//go:build unit

package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/kalilventura/vehicle-management/internal/vehicles/infrastructure/controllers"
	"github.com/kalilventura/vehicle-management/test/vehicles/domain/builders"
	"github.com/kalilventura/vehicle-management/test/vehicles/domain/commands"
	builders2 "github.com/kalilventura/vehicle-management/test/vehicles/infrastructure/builders"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestUpdateVehicleController(t *testing.T) {
	const route = "/v1/vehicles/%s"
	t.Run("should respond OK when the vehicle was updated", func(t *testing.T) {
		// given
		output := builders.NewUpdateVehicleInputBuilder().Build()
		updateRequest := builders2.NewUpdateVehicleRequestBuilder().WithValidDefaults().BuildRequest()

		router := echo.New()
		endpoint := fmt.Sprintf(route, gofakeit.UUID())
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPatch, endpoint, updateRequest)
		request.Header.Set("Content-Type", "application/json")

		command := commands.NewUpdateVehicleCommandStub().WithOnSuccess(output)
		controller := controllers.NewUpdateVehicleController(command)
		router.PATCH(endpoint, controller.Execute)

		// when
		router.ServeHTTP(recorder, request)

		// then
		assert.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("should respond Bad Request when the request is invalid", func(t *testing.T) {
		// given
		data := "{invalid"
		requestBodyBytes, _ := json.Marshal(data)
		reqBody := bytes.NewBuffer(requestBodyBytes)

		router := echo.New()
		endpoint := fmt.Sprintf(route, gofakeit.UUID())
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPatch, endpoint, reqBody)
		request.Header.Set("Content-Type", "application/json")

		command := commands.NewUpdateVehicleCommandStub()
		controller := controllers.NewUpdateVehicleController(command)
		router.PATCH(endpoint, controller.Execute)

		// when
		router.ServeHTTP(recorder, request)

		// then
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("should respond Not Found when the vehicle not exists", func(t *testing.T) {
		// given
		updateRequest := builders2.NewUpdateVehicleRequestBuilder().WithValidDefaults().BuildRequest()

		router := echo.New()
		endpoint := fmt.Sprintf(route, gofakeit.UUID())
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPatch, endpoint, updateRequest)
		request.Header.Set("Content-Type", "application/json")

		command := commands.NewUpdateVehicleCommandStub().WithOnNotFound()
		controller := controllers.NewUpdateVehicleController(command)
		router.PATCH(endpoint, controller.Execute)

		// when
		router.ServeHTTP(recorder, request)

		// then
		assert.Equal(t, http.StatusNotFound, recorder.Code)
	})

	t.Run("should respond Internal Server Error due an unexpected error", func(t *testing.T) {
		// given
		updateRequest := builders2.NewUpdateVehicleRequestBuilder().WithValidDefaults().BuildRequest()

		router := echo.New()
		endpoint := fmt.Sprintf(route, gofakeit.UUID())
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPatch, endpoint, updateRequest)
		request.Header.Set("Content-Type", "application/json")

		command := commands.NewUpdateVehicleCommandStub().WithOnInternalServerError()
		controller := controllers.NewUpdateVehicleController(command)
		router.PATCH(endpoint, controller.Execute)

		// when
		router.ServeHTTP(recorder, request)

		// then
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("should return the metadata", func(t *testing.T) {
		// given
		controller := controllers.NewUpdateVehicleController(nil)

		// when
		metadata := controller.GetBind()

		// then
		assert.Equal(t, "/vehicles/:id", metadata.RelativePath)
		assert.Equal(t, http.MethodPatch, metadata.Method)
	})
}
