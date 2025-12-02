//go:build unit

package valueobjects_test

import (
	"testing"

	valueobjects "github.com/kalilventura/vehicle-management/internal/vehicles/domain/value-objects"
	"github.com/stretchr/testify/assert"
)

func TestMileage(t *testing.T) {
	t.Run("should create valid mileage values", func(t *testing.T) {
		tests := []struct {
			name  string
			input int
		}{
			{"zero mileage", 0},
			{"low mileage", 1000},
			{"high mileage", 100000},
			{"max int value", 1<<31 - 1},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// given & when
				mileage, err := valueobjects.NewMileage(tt.input)

				// then
				assert.NoError(t, err)
				assert.Equal(t, tt.input, mileage.Value())
			})
		}
	})

	t.Run("should return error for invalid mileage values", func(t *testing.T) {
		tests := []struct {
			name  string
			input int
		}{
			{"negative value", -1},
			{"large negative value", -10000},
			{"minimum int value", -1 << 31},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// given & when
				_, err := valueobjects.NewMileage(tt.input)

				// then
				assert.Error(t, err)
				assert.EqualError(t, err, "mileage cannot be negative")
			})
		}
	})

	t.Run("should correctly return integer value", func(t *testing.T) {
		// given
		testValue := 50000

		// when
		mileage, _ := valueobjects.NewMileage(testValue)
		value := mileage.Value()

		// then
		assert.Equal(t, testValue, value)
		assert.IsType(t, 0, value)
	})

	t.Run("should check if mileage is greater than value", func(t *testing.T) {
		// given
		mileage, _ := valueobjects.NewMileage(10000)

		// then
		assert.True(t, mileage.IsGreaterThan(5000))
		assert.False(t, mileage.IsGreaterThan(10000))
		assert.False(t, mileage.IsGreaterThan(15000))
	})
}
