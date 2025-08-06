//go:build unit

package dtos_test

import (
	"strings"
	"testing"

	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities/dtos"
	"github.com/stretchr/testify/assert"
)

func TestFuelType(t *testing.T) {
	t.Run("should create valid fuel types", func(t *testing.T) {
		tests := []struct {
			name  string
			input string
			want  dtos.FuelType
		}{
			{"gasoline", "gasoline", dtos.Gasoline},
			{"ethanol", "ethanol", dtos.Ethanol},
			{"diesel", "diesel", dtos.Diesel},
			{"flex", "flex", dtos.Flex},
			{"hybrid", "hybrid", dtos.Hybrid},
			{"electric", "electric", dtos.Electric},
			{"gnv", "gnv", dtos.GNV},

			{"uppercase GASOLINE", "GASOLINE", dtos.Gasoline},
			{"mixed case DiEsEl", "DiEsEl", dtos.Diesel},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// given & when
				fuelType, err := dtos.NewFuelType(tt.input)

				// then
				assert.NoError(t, err)
				assert.Equal(t, tt.want, fuelType)
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
				_, err := dtos.NewFuelType(tt.input)

				// then
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "invalid fuel type")
			})
		}
	})

	t.Run("should expose correct constant values", func(t *testing.T) {
		assert.Equal(t, "gasoline", dtos.Gasoline.Value())
		assert.Equal(t, "ethanol", dtos.Ethanol.Value())
		assert.Equal(t, "diesel", dtos.Diesel.Value())
		assert.Equal(t, "flex", dtos.Flex.Value())
		assert.Equal(t, "hybrid", dtos.Hybrid.Value())
		assert.Equal(t, "electric", dtos.Electric.Value())
		assert.Equal(t, "gnv", dtos.GNV.Value())
	})

	t.Run("should be case insensitive when creating", func(t *testing.T) {
		// given & when
		f1, _ := dtos.NewFuelType("GASOLINE")
		f2, _ := dtos.NewFuelType("gasoline")
		f3, _ := dtos.NewFuelType("Gasoline")

		// then
		assert.Equal(t, dtos.Gasoline, f1)
		assert.Equal(t, f1, f2)
		assert.Equal(t, f2, f3)
		assert.Equal(t, "gasoline", f1.Value())
	})
}
