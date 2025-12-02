//go:build unit

package valueobjects_test

import (
	"strings"
	"testing"

	valueobjects "github.com/kalilventura/vehicle-management/internal/vehicles/domain/value-objects"
	"github.com/stretchr/testify/assert"
)

func TestStatus(t *testing.T) {
	t.Run("should create valid status values", func(t *testing.T) {
		tests := []struct {
			name  string
			input string
		}{
			{"available status", "available"},
			{"reserved status", "reserved"},
			{"sold status", "sold"},
			{"maintenance status", "maintenance"},
			{"uppercase AVAILABLE", "AVAILABLE"},
			{"mixed case SoLd", "SoLd"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// given & when
				status, err := valueobjects.NewVehicleStatus(tt.input)

				// then
				assert.NoError(t, err)
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
				_, err := valueobjects.NewVehicleStatus(tt.input)

				// then
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "invalid")
			})
		}
	})

	t.Run("should expose correct constant values", func(t *testing.T) {
		assert.Equal(t, "available", valueobjects.StatusAvailable)
		assert.Equal(t, "reserved", valueobjects.StatusReserved)
		assert.Equal(t, "sold", valueobjects.StatusSold)
		assert.Equal(t, "maintenance", valueobjects.StatusMaintenance)
	})

	t.Run("should be case insensitive when creating", func(t *testing.T) {
		// given & when
		s1, _ := valueobjects.NewVehicleStatus("AVAILABLE")
		s2, _ := valueobjects.NewVehicleStatus("available")
		s3, _ := valueobjects.NewVehicleStatus("Available")

		// then
		assert.Equal(t, "available", s1.Value())
		assert.Equal(t, s1, s2)
		assert.Equal(t, s2, s3)
	})

	t.Run("should use factory methods", func(t *testing.T) {
		// given & when
		available := valueobjects.Available()
		reserved := valueobjects.Reserved()
		sold := valueobjects.Sold()
		maintenance := valueobjects.Maintenance()

		// then
		assert.True(t, available.IsAvailable())
		assert.False(t, available.IsReserved())
		assert.False(t, available.IsSold())
		assert.False(t, available.IsMaintenance())

		assert.True(t, reserved.IsReserved())
		assert.True(t, sold.IsSold())
		assert.True(t, maintenance.IsMaintenance())
	})

	t.Run("should validate business rules for selling", func(t *testing.T) {
		// given
		available := valueobjects.Available()
		reserved := valueobjects.Reserved()
		sold := valueobjects.Sold()
		maintenance := valueobjects.Maintenance()

		// then - can be sold
		assert.True(t, available.CanBeSold())
		assert.True(t, reserved.CanBeSold())
		assert.False(t, sold.CanBeSold())
		assert.False(t, maintenance.CanBeSold())
	})

	t.Run("should validate business rules for reservation", func(t *testing.T) {
		// given
		available := valueobjects.Available()
		reserved := valueobjects.Reserved()
		sold := valueobjects.Sold()
		maintenance := valueobjects.Maintenance()

		// then - can be reserved
		assert.True(t, available.CanBeReserved())
		assert.False(t, reserved.CanBeReserved())
		assert.False(t, sold.CanBeReserved())
		assert.False(t, maintenance.CanBeReserved())
	})
}
