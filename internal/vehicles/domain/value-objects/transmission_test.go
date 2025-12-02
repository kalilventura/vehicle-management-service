//go:build unit

package valueobjects_test

import (
	"strings"
	"testing"

	valueobjects "github.com/kalilventura/vehicle-management/internal/vehicles/domain/value-objects"
	"github.com/stretchr/testify/assert"
)

func TestTransmission(t *testing.T) {
	t.Run("should create valid transmission types", func(t *testing.T) {
		tests := []struct {
			name  string
			input string
		}{
			{"manual transmission", "manual"},
			{"automatic transmission", "automatic"},
			{"CVT transmission", "cvt"},
			{"uppercase MANUAL", "MANUAL"},
			{"mixed case AuToMaTiC", "AuToMaTiC"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// given & when
				transmission, err := valueobjects.NewTransmission(tt.input)

				// then
				assert.NoError(t, err)
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
				_, err := valueobjects.NewTransmission(tt.input)

				// then
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "invalid transmission")
			})
		}
	})

	t.Run("should expose correct constant values", func(t *testing.T) {
		assert.Equal(t, "manual", valueobjects.TransmissionManual)
		assert.Equal(t, "automatic", valueobjects.TransmissionAutomatic)
		assert.Equal(t, "cvt", valueobjects.TransmissionCVT)
	})

	t.Run("should be case insensitive when creating", func(t *testing.T) {
		// given & when
		t1, _ := valueobjects.NewTransmission("MANUAL")
		t2, _ := valueobjects.NewTransmission("manual")
		t3, _ := valueobjects.NewTransmission("Manual")

		// then
		assert.Equal(t, "manual", t1.Value())
		assert.Equal(t, t1, t2)
		assert.Equal(t, t2, t3)
	})
}
