//go:build unit

package valueobjects_test

import (
	"strings"
	"testing"

	valueobjects "github.com/kalilventura/vehicle-management/internal/vehicles/domain/value-objects"
	"github.com/stretchr/testify/assert"
)

func TestFuelType(t *testing.T) {
	t.Run("should create valid fuel types", func(t *testing.T) {
		tests := []struct {
			name  string
			input string
		}{
			{"gasoline", "gasoline"},
			{"ethanol", "ethanol"},
			{"diesel", "diesel"},
			{"flex", "flex"},
			{"hybrid", "hybrid"},
			{"electric", "electric"},
			{"gnv", "gnv"},
			{"uppercase GASOLINE", "GASOLINE"},
			{"mixed case DiEsEl", "DiEsEl"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// given & when
				fuelType, err := valueobjects.NewFuelType(tt.input)

				// then
				assert.NoError(t, err)
				assert.Equal(t, strings.ToLower(tt.input), fuelType.Value())
			})
		}
	})

	t.Run("should return error for invalid fuel types", func(t *testing.T) {
		tests := []struct {
			name  string
			input string
		}{
			{"empty string", ""},
			{"unknown type", "hydrogen"},
			{"invalid variation", "flexfuel"},
			{"space padded", " diesel "},
			{"numeric", "123"},
			{"special chars", "electric!"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// given & when
				_, err := valueobjects.NewFuelType(tt.input)

				// then
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "invalid fuel type")
			})
		}
	})

	t.Run("should expose correct constant values", func(t *testing.T) {
		assert.Equal(t, "gasoline", valueobjects.FuelTypeGasoline)
		assert.Equal(t, "ethanol", valueobjects.FuelTypeEthanol)
		assert.Equal(t, "diesel", valueobjects.FuelTypeDiesel)
		assert.Equal(t, "flex", valueobjects.FuelTypeFlex)
		assert.Equal(t, "hybrid", valueobjects.FuelTypeHybrid)
		assert.Equal(t, "electric", valueobjects.FuelTypeElectric)
		assert.Equal(t, "gnv", valueobjects.FuelTypeGNV)
	})

	t.Run("should be case insensitive when creating", func(t *testing.T) {
		// given & when
		f1, _ := valueobjects.NewFuelType("GASOLINE")
		f2, _ := valueobjects.NewFuelType("gasoline")
		f3, _ := valueobjects.NewFuelType("Gasoline")

		// then
		assert.Equal(t, valueobjects.FuelTypeGasoline, f1)
		assert.Equal(t, "gasoline", f1.Value())
		assert.Equal(t, f1, f2)
		assert.Equal(t, f2, f3)
		assert.Equal(t, "gasoline", f1.Value())
	})
}
