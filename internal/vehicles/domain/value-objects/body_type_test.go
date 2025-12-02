//go:build unit

package valueobjects_test

import (
	"strings"
	"testing"

	valueobjects "github.com/kalilventura/vehicle-management/internal/vehicles/domain/value-objects"
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
				bodyType, err := valueobjects.NewBodyType(tt.input)

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
				_, err := valueobjects.NewBodyType(tt.input)

				// then
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "invalid body type")
			})
		}
	})

	t.Run("should expose correct constant values", func(t *testing.T) {
		assert.Equal(t, "sedan", valueobjects.BodyTypeSedan)
		assert.Equal(t, "hatchback", valueobjects.BodyTypeHatchback)
		assert.Equal(t, "suv", valueobjects.BodyTypeSUV)
		assert.Equal(t, "pickup", valueobjects.BodyTypePickup)
		assert.Equal(t, "coupe", valueobjects.BodyTypeCoupe)
		assert.Equal(t, "convertible", valueobjects.BodyTypeConvertible)
		assert.Equal(t, "wagon", valueobjects.BodyTypeWagon)
		assert.Equal(t, "minivan", valueobjects.BodyTypeMinivan)
		assert.Equal(t, "fastback", valueobjects.BodyTypeFastback)
	})

	t.Run("should be case insensitive when creating", func(t *testing.T) {
		// given & when
		bt1, _ := valueobjects.NewBodyType("SUV")
		bt2, _ := valueobjects.NewBodyType("suv")
		bt3, _ := valueobjects.NewBodyType("Suv")

		// then
		assert.Equal(t, "suv", bt1.Value())
		assert.Equal(t, bt1, bt2)
		assert.Equal(t, bt2, bt3)
	})
}
