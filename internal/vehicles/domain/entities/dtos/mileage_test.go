//go:build unit

package dtos_test

import (
	"testing"

	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities/dtos"
	"github.com/stretchr/testify/assert"
)

func TestMileage(t *testing.T) {
	t.Run("should create valid mileage values", func(t *testing.T) {
		tests := []struct {
			name  string
			input int
			want  dtos.Mileage
		}{
			{"zero mileage", 0, dtos.Mileage(0)},
			{"low mileage", 1000, dtos.Mileage(1000)},
			{"high mileage", 100000, dtos.Mileage(100000)},
			{"max int value", 1<<31 - 1, dtos.Mileage(1<<31 - 1)},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// given & when
				mileage, err := dtos.NewMileage(tt.input)

				// then
				assert.NoError(t, err)
				assert.Equal(t, tt.want, mileage)
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
				_, err := dtos.NewMileage(tt.input)

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
		mileage, _ := dtos.NewMileage(testValue)
		value := mileage.Value()

		// then
		assert.Equal(t, testValue, value)
		assert.IsType(t, 0, value)
	})
}
