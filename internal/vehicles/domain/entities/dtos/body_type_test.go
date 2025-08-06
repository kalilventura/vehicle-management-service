//go:build unit

package dtos_test

import (
	"strings"
	"testing"

	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities/dtos"
	"github.com/stretchr/testify/assert"
)

func TestBodyType(t *testing.T) {
	t.Run("should create valid body types", func(t *testing.T) {
		tests := []struct {
			name  string
			input string
		}{
			{"sedan", "sedan"},
			{"hatchback", "hatchback"},
			{"suv", "suv"},
			{"pickup", "pickup"},
			{"coupe", "coupe"},
			{"convertible", "convertible"},
			{"wagon", "wagon"},
			{"minivan", "minivan"},
			{"fastback", "fastback"},
			// Teste case insensitive
			{"uppercase SUV", "SUV"},
			{"mixed case SeDaN", "SeDaN"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// given & when
				bodyType, err := dtos.NewBodyType(tt.input)

				// then
				assert.NoError(t, err)
				assert.Equal(t, strings.ToLower(tt.input), bodyType.Value())
			})
		}
	})

	t.Run("should return error for invalid body types", func(t *testing.T) {
		tests := []struct {
			name  string
			input string
		}{
			{"empty string", ""},
			{"unknown type", "truck"},
			{"invalid variation", "sedanx"},
			{"space padded", " suv "},
			{"numeric", "123"},
			{"special chars", "suv!"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// given & when
				_, err := dtos.NewBodyType(tt.input)

				// then
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "invalid body type")
			})
		}
	})

	t.Run("should expose correct constant values", func(t *testing.T) {
		assert.Equal(t, "sedan", dtos.Sedan.Value())
		assert.Equal(t, "hatchback", dtos.Hatchback.Value())
		assert.Equal(t, "suv", dtos.SUV.Value())
		assert.Equal(t, "pickup", dtos.Pickup.Value())
		assert.Equal(t, "coupe", dtos.Coupe.Value())
		assert.Equal(t, "convertible", dtos.Convertible.Value())
		assert.Equal(t, "wagon", dtos.Wagon.Value())
		assert.Equal(t, "minivan", dtos.Minivan.Value())
		assert.Equal(t, "fastback", dtos.Fastback.Value())
	})

	t.Run("should be case insensitive when creating", func(t *testing.T) {
		// given & when
		bt1, _ := dtos.NewBodyType("SUV")
		bt2, _ := dtos.NewBodyType("suv")
		bt3, _ := dtos.NewBodyType("Suv")

		// then
		assert.Equal(t, "suv", bt1.Value())
		assert.Equal(t, bt1, bt2)
		assert.Equal(t, bt2, bt3)
	})
}
