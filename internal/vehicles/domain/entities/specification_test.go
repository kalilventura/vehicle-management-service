//go:build unit

package entities_test

import (
	"testing"

	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities/dtos"
	"github.com/stretchr/testify/assert"
)

func TestSpecificationMethods(t *testing.T) {
	t.Run("MileageIsGreaterThan - when greater", func(t *testing.T) {
		mileage, _ := dtos.NewMileage(10000)
		spec := entities.Specification{
			Mileage: mileage,
		}
		assert.True(t, spec.MileageIsGreaterThan(5000))
	})

	t.Run("MileageIsGreaterThan - when equal", func(t *testing.T) {
		mileage, _ := dtos.NewMileage(10000)
		spec := entities.Specification{
			Mileage: mileage,
		}
		assert.False(t, spec.MileageIsGreaterThan(10000))
	})

	t.Run("MileageIsGreaterThan - when less", func(t *testing.T) {
		mileage, _ := dtos.NewMileage(5000)
		spec := entities.Specification{
			Mileage: mileage,
		}
		assert.False(t, spec.MileageIsGreaterThan(10000))
	})

	t.Run("GetBodyType", func(t *testing.T) {
		expected := "suv"
		bodyType, _ := dtos.NewBodyType(expected)
		spec := entities.Specification{
			BodyType: bodyType,
		}
		assert.Equal(t, expected, spec.GetBodyType())
	})

	t.Run("GetFuelType", func(t *testing.T) {
		expected := "electric"
		fuelType, _ := dtos.NewFuelType(expected)
		spec := entities.Specification{
			FuelType: fuelType,
		}
		assert.Equal(t, expected, spec.GetFuelType())
	})

	t.Run("GetMileage", func(t *testing.T) {
		expected := 25000
		mileage, _ := dtos.NewMileage(expected)
		spec := entities.Specification{
			Mileage: mileage,
		}
		assert.Equal(t, expected, spec.GetMileage())
	})

	t.Run("GetDoors", func(t *testing.T) {
		expected := 4
		doors, _ := dtos.NewDoors(expected)
		spec := entities.Specification{
			Doors: doors,
		}
		assert.Equal(t, expected, spec.GetDoors())
	})

	t.Run("GetEngine", func(t *testing.T) {
		expected := "2.0L Turbo"
		spec := entities.Specification{
			Engine: expected,
		}
		assert.Equal(t, expected, spec.GetEngine())
	})

	t.Run("GetTransmission", func(t *testing.T) {
		expected := "automatic"
		transmission, _ := dtos.NewTransmission(expected)
		spec := entities.Specification{
			Transmission: transmission,
		}
		assert.Equal(t, expected, spec.GetTransmission())
	})

	t.Run("ZeroValueSpecification", func(t *testing.T) {
		var spec entities.Specification
		assert.Equal(t, "", spec.GetBodyType())
		assert.Equal(t, "", spec.GetFuelType())
		assert.Equal(t, 0, spec.GetMileage())
		assert.Equal(t, 0, spec.GetDoors())
		assert.Equal(t, "", spec.GetEngine())
		assert.Equal(t, "", spec.GetTransmission())
		assert.False(t, spec.MileageIsGreaterThan(0))
	})
}
