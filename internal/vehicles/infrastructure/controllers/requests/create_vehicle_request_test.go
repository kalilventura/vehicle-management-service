//go:build unit

package requests_test

import (
	"testing"

	"github.com/kalilventura/vehicle-management/internal/vehicles/infrastructure/controllers/requests"
	"github.com/stretchr/testify/assert"
)

func TestCreateVehicleRequest(t *testing.T) {
	newValidRequest := func() *requests.CreateVehicleRequest {
		status := "available"
		return &requests.CreateVehicleRequest{
			Price:              25000.50,
			Brand:              "Toyota",
			Model:              "Corolla",
			Year:               2020,
			BodyType:           "sedan",
			Transmission:       "automatic",
			FuelType:           "gasoline",
			Color:              "Red",
			Mileage:            15000,
			Engine:             "2.0L",
			Doors:              4,
			Condition:          "used",
			Description:        "Well maintained vehicle",
			Status:             &status,
			HasAirConditioning: true,
			HasPowerWindows:    true,
		}
	}

	t.Run("successful conversion with all fields", func(t *testing.T) {
		req := newValidRequest()
		vehicle, err := req.ToDomain()

		assert.NoError(t, err)
		assert.Equal(t, req.Brand, vehicle.Brand)
		assert.Equal(t, req.Model, vehicle.Model)
		assert.Equal(t, req.Color, vehicle.Color)
		assert.Equal(t, req.Description, vehicle.Description)
		assert.Equal(t, req.Price, vehicle.Price.Value())
		assert.Equal(t, req.Year, vehicle.Year.Value())
		assert.Equal(t, req.Condition, vehicle.Condition.Value())
		assert.Equal(t, *req.Status, vehicle.Status.Value())
		assert.Equal(t, req.BodyType, vehicle.Specification.BodyType.Value())
		assert.Equal(t, req.Transmission, vehicle.Specification.Transmission.Value())
		assert.Equal(t, req.FuelType, vehicle.Specification.FuelType.Value())
		assert.Equal(t, req.Mileage, vehicle.Specification.Mileage.Value())
		assert.Equal(t, req.Doors, vehicle.Specification.Doors.Value())
		assert.Equal(t, req.Engine, vehicle.Specification.Engine)
		assert.Equal(t, req.HasAirConditioning, vehicle.Features.HasAirConditioning)
		assert.Equal(t, req.HasPowerWindows, vehicle.Features.HasPowerWindows)
	})

	t.Run("successful conversion with nil status", func(t *testing.T) {
		req := newValidRequest()
		req.Status = nil
		vehicle, err := req.ToDomain()

		assert.NoError(t, err)
		assert.Equal(t, "available", vehicle.Status.Value())
	})

	t.Run("invalid condition", func(t *testing.T) {
		req := newValidRequest()
		req.Condition = "invalid"
		_, err := req.ToDomain()

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid condition")
	})

	t.Run("invalid year", func(t *testing.T) {
		req := newValidRequest()
		req.Year = 1899
		_, err := req.ToDomain()

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid year")
	})

	t.Run("invalid price", func(t *testing.T) {
		req := newValidRequest()
		req.Price = -100
		_, err := req.ToDomain()

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid price")
	})

	t.Run("invalid body type", func(t *testing.T) {
		req := newValidRequest()
		req.BodyType = ""
		_, err := req.ToDomain()

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid body type")
	})

	t.Run("invalid transmission", func(t *testing.T) {
		req := newValidRequest()
		req.Transmission = ""
		_, err := req.ToDomain()

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid transmission")
	})

	t.Run("invalid fuel type", func(t *testing.T) {
		req := newValidRequest()
		req.FuelType = ""
		_, err := req.ToDomain()

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid fuel type")
	})

	t.Run("invalid mileage", func(t *testing.T) {
		req := newValidRequest()
		req.Mileage = -100
		_, err := req.ToDomain()

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid mileage")
	})

	t.Run("invalid doors", func(t *testing.T) {
		req := newValidRequest()
		req.Doors = 1
		_, err := req.ToDomain()

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid door count")
	})

	t.Run("invalid status", func(t *testing.T) {
		req := newValidRequest()
		invalidStatus := "invalid"
		req.Status = &invalidStatus
		_, err := req.ToDomain()

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid status")
	})

	t.Run("new vehicle with mileage > 0", func(t *testing.T) {
		req := newValidRequest()
		req.Condition = "new"
		req.Mileage = 100
		vehicle, err := req.ToDomain()

		assert.NoError(t, err)
		assert.Equal(t, "new", vehicle.Condition.Value())
		assert.Equal(t, 100, vehicle.Specification.Mileage.Value())
		// Note: The validation of new vehicle with mileage > 0 should be handled by the Vehicle entity
	})

	t.Run("zero value request", func(t *testing.T) {
		req := &requests.CreateVehicleRequest{}
		_, err := req.ToDomain()

		assert.Error(t, err)
		// Multiple errors expected, but the first one will be about condition
		assert.Contains(t, err.Error(), "invalid condition")
	})
}
