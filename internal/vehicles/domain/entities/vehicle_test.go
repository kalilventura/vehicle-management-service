//go:build unit

package entities_test

import (
	"testing"

	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities/dtos"
	"github.com/stretchr/testify/assert"
)

func TestVehicleMethods(t *testing.T) {

	newValidVehicle := func() entities.Vehicle {
		price, _ := dtos.NewPrice(25000.50)
		mileage, _ := dtos.NewMileage(10000)
		condition, _ := dtos.NewCondition("used")
		year, _ := dtos.NewYear(2020)
		status, _ := dtos.NewStatus("available")
		return entities.Vehicle{
			ID:          "VH123456",
			Brand:       "Toyota",
			Model:       "Corolla",
			Color:       "Red",
			Description: "A reliable sedan",
			Price:       price,
			Features:    entities.Features{},
			Specification: entities.Specification{
				Mileage: mileage,
			},
			Status:    status,
			Condition: condition,
			Year:      year,
		}
	}

	t.Run("IsValid - valid used vehicle", func(t *testing.T) {
		v := newValidVehicle()
		assert.NoError(t, v.IsValid())
	})

	t.Run("IsValid - valid new vehicle with zero mileage", func(t *testing.T) {
		v := newValidVehicle()
		cond, _ := dtos.NewCondition("new")
		v.Condition = cond
		mileage, _ := dtos.NewMileage(0)
		v.Specification.Mileage = mileage
		assert.NoError(t, v.IsValid())
	})

	t.Run("IsValid - invalid new vehicle with mileage > 0", func(t *testing.T) {
		v := newValidVehicle()
		cond, _ := dtos.NewCondition("new")
		v.Condition = cond
		mileage, _ := dtos.NewMileage(100)
		v.Specification.Mileage = mileage
		assert.EqualError(t, v.IsValid(), "a new vehicle must have a mileage greater than zero")
	})

	t.Run("GetPrice", func(t *testing.T) {
		v := newValidVehicle()
		assert.Equal(t, 25000.50, v.GetPrice())
	})

	t.Run("GetStatus", func(t *testing.T) {
		v := newValidVehicle()
		assert.Equal(t, "available", v.GetStatus())
	})

	t.Run("GetCondition", func(t *testing.T) {
		v := newValidVehicle()
		assert.Equal(t, "used", v.GetCondition())
	})

	t.Run("GetYear", func(t *testing.T) {
		v := newValidVehicle()
		assert.Equal(t, 2020, v.GetYear())
	})

	t.Run("ZeroValueVehicle", func(t *testing.T) {
		var v entities.Vehicle
		assert.Equal(t, 0.0, v.GetPrice())
		assert.Equal(t, "", v.GetStatus())
		assert.Equal(t, "", v.GetCondition())
		assert.Equal(t, 0, v.GetYear())
		assert.NoError(t, v.IsValid()) // Zero value vehicle should be considered valid
	})
}
