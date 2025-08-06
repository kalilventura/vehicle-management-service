//go:build unit

package dtos_test

import (
	"strings"
	"testing"

	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities/dtos"
	"github.com/stretchr/testify/assert"
)

func TestStatus(t *testing.T) {
	t.Run("should create valid status values", func(t *testing.T) {
		tests := []struct {
			name  string
			input string
			want  dtos.Status
		}{
			{"available status", "available", dtos.Available},
			{"reserved status", "reserved", dtos.Reserved},
			{"sold status", "sold", dtos.Sold},
			{"maintenance status", "maintenance", dtos.Maintenance},

			{"uppercase AVAILABLE", "AVAILABLE", dtos.Available},
			{"mixed case SoLd", "SoLd", dtos.Sold},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// given & when
				status, err := dtos.NewStatus(tt.input)

				// then
				assert.NoError(t, err)
				assert.Equal(t, tt.want, status)
				assert.Equal(t, strings.ToLower(tt.input), status.Value())
			})
		}
	})

	t.Run("should return error for invalid status values", func(t *testing.T) {
		tests := []struct {
			name  string
			input string
		}{
			{"empty string", ""},
			{"unknown status", "pending"},
			{"invalid variation", "available2"},
			{"space padded", " reserved "},
			{"numeric", "123"},
			{"special chars", "sold!"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// given & when
				_, err := dtos.NewStatus(tt.input)

				// then
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "invalid status")
			})
		}
	})

	t.Run("should expose correct constant values", func(t *testing.T) {
		assert.Equal(t, "available", dtos.Available.Value())
		assert.Equal(t, "reserved", dtos.Reserved.Value())
		assert.Equal(t, "sold", dtos.Sold.Value())
		assert.Equal(t, "maintenance", dtos.Maintenance.Value())
	})

	t.Run("should be case insensitive when creating", func(t *testing.T) {
		// given & when
		s1, _ := dtos.NewStatus("AVAILABLE")
		s2, _ := dtos.NewStatus("available")
		s3, _ := dtos.NewStatus("Available")

		// then
		assert.Equal(t, dtos.Available, s1)
		assert.Equal(t, s1, s2)
		assert.Equal(t, s2, s3)
		assert.Equal(t, "available", s1.Value())
	})

	t.Run("should maintain original case in constants", func(t *testing.T) {
		assert.Equal(t, "available", string(dtos.Available))
		assert.Equal(t, "reserved", string(dtos.Reserved))
		assert.Equal(t, "sold", string(dtos.Sold))
		assert.Equal(t, "maintenance", string(dtos.Maintenance))
	})
}
