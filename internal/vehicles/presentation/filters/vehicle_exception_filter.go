package filters

import (
	"net/http"

	"github.com/kalilventura/vehicle-management/internal/shared/infrastructure/controllers"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/exceptions"
	domainerr "github.com/kalilventura/vehicle-management/internal/shared/domain/errors"
	"github.com/labstack/echo/v4"
	logger "github.com/sirupsen/logrus"
)

// VehicleExceptionFilter handles domain exceptions and maps them to HTTP status codes
type VehicleExceptionFilter struct{}

// NewVehicleExceptionFilter creates a new VehicleExceptionFilter
func NewVehicleExceptionFilter() *VehicleExceptionFilter {
	return &VehicleExceptionFilter{}
}

// HandleError handles errors and returns appropriate HTTP response
func (f *VehicleExceptionFilter) HandleError(ectx echo.Context, err error) error {
	if err == nil {
		return nil
	}

	var statusCode int
	var message string
	var details interface{}

	// Map domain exceptions to HTTP status codes
	switch e := err.(type) {
	case exceptions.VehicleNotFoundException:
		statusCode = http.StatusNotFound
		message = e.Error()
		details = map[string]string{
			"vehicle_id": e.VehicleID,
		}

	case exceptions.InvalidVehicleStatusException:
		statusCode = http.StatusBadRequest
		message = e.Error()
		details = map[string]string{
			"vehicle_id":     e.VehicleID,
			"current_status": e.CurrentStatus,
			"operation":      e.Operation,
		}

	case exceptions.VehicleValidationException:
		statusCode = http.StatusBadRequest
		message = e.Error()
		details = map[string]string{
			"field":   e.Field,
			"message": e.Message,
		}

	default:
		// Check for shared domain errors
		if err == domainerr.ErrRecordNotFound {
			statusCode = http.StatusNotFound
			message = "resource not found"
		} else {
			// Unknown error - log and return 500
			logger.Errorf("Unhandled error: %v", err)
			statusCode = http.StatusInternalServerError
			message = "internal server error"
		}
	}

	response := controllers.NewErrorResponse(statusCode, details)
	if message != "" {
		// Add message to response if needed
		response.Error = message
	}

	return ectx.JSON(statusCode, response)
}

