//go:build unit

package dtos_test

import (
	"strings"
	"testing"

	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities/dtos"
	"github.com/stretchr/testify/assert"
)

func TestTransmission(t *testing.T) {
	t.Run("should create valid transmission types", func(t *testing.T) {
		tests := []struct {
			name  string
			input string
			want  dtos.Transmission
		}{
			{"manual transmission", "manual", dtos.Manual},
			{"automatic transmission", "automatic", dtos.Automatic},
			{"CVT transmission", "cvt", dtos.CVT},

			{"uppercase MANUAL", "MANUAL", dtos.Manual},
			{"mixed case AuToMaTiC", "AuToMaTiC", dtos.Automatic},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// given & when
				transmission, err := dtos.NewTransmission(tt.input)

				// then
				assert.NoError(t, err)
				assert.Equal(t, tt.want, transmission)
				assert.Equal(t, strings.ToLower(tt.input), transmission.Value())
			})
		}
	})

	t.Run("should return error for invalid transmission types", func(t *testing.T) {
		tests := []struct {
			name  string
			input string
		}{
			{"empty string", ""},
			{"unknown type", "semi-automatic"},
			{"invalid variation", "automatic2"},
			{"space padded", " manual "},
			{"numeric", "123"},
			{"special chars", "cvt!"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// given & when
				_, err := dtos.NewTransmission(tt.input)

				// then
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "invalid transmission")
			})
		}
	})

	t.Run("should expose correct constant values", func(t *testing.T) {
		assert.Equal(t, "manual", dtos.Manual.Value())
		assert.Equal(t, "automatic", dtos.Automatic.Value())
		assert.Equal(t, "cvt", dtos.CVT.Value())
	})

	t.Run("should be case insensitive when creating", func(t *testing.T) {
		// given & when
		t1, _ := dtos.NewTransmission("MANUAL")
		t2, _ := dtos.NewTransmission("manual")
		t3, _ := dtos.NewTransmission("Manual")

		// then
		assert.Equal(t, dtos.Manual, t1)
		assert.Equal(t, t1, t2)
		assert.Equal(t, t2, t3)
		assert.Equal(t, "manual", t1.Value())
	})

	t.Run("should maintain original case in constants", func(t *testing.T) {
		assert.Equal(t, "manual", dtos.Manual.Value())
		assert.Equal(t, "automatic", dtos.Automatic.Value())
		assert.Equal(t, "cvt", dtos.CVT.Value())
	})
}
