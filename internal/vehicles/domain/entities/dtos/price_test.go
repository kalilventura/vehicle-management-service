//go:build unit

package dtos_test

import (
	"testing"

	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities/dtos"
	"github.com/stretchr/testify/assert"
)

func TestPrice(t *testing.T) {
	t.Run("should create valid price values", func(t *testing.T) {
		tests := []struct {
			name  string
			input float64
			want  dtos.Price
		}{
			{"zero price", 0, dtos.Price(0)},
			{"low price", 1000.50, dtos.Price(1000.50)},
			{"high price", 500000.99, dtos.Price(500000.99)},
			{"integer price", 25000, dtos.Price(25000)},
			{"max float value", 1.7976931348623157e+308, dtos.Price(1.7976931348623157e+308)},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// given & when
				price, err := dtos.NewPrice(tt.input)

				// then
				assert.NoError(t, err)
				assert.Equal(t, tt.want, price)
				assert.Equal(t, tt.input, price.Value())
			})
		}
	})

	t.Run("should return error for invalid price values", func(t *testing.T) {
		tests := []struct {
			name  string
			input float64
		}{
			{"negative value", -1.50},
			{"large negative value", -1000000.99},
			{"minimum float value", -1.7976931348623157e+308},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// given & when
				_, err := dtos.NewPrice(tt.input)

				// then
				assert.Error(t, err)
				assert.EqualError(t, err, "price cannot be negative")
			})
		}
	})

	t.Run("should correctly return float64 value", func(t *testing.T) {
		// given
		testValue := 37500.75

		// when
		price, _ := dtos.NewPrice(testValue)
		value := price.Value()

		// then
		assert.Equal(t, testValue, value)
		assert.IsType(t, 0.0, value)
	})

	t.Run("should handle edge cases", func(t *testing.T) {
		t.Run("very small positive value", func(t *testing.T) {
			// given
			smallValue := 0.0000001

			// when
			price, err := dtos.NewPrice(smallValue)

			// then
			assert.NoError(t, err)
			assert.Equal(t, smallValue, price.Value())
		})

		t.Run("zero value", func(t *testing.T) {
			// given & when
			price, err := dtos.NewPrice(0)

			// then
			assert.NoError(t, err)
			assert.Equal(t, 0.0, price.Value())
		})
	})
}
