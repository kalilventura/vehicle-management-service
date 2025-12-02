//go:build unit

package valueobjects_test

import (
	"strings"
	"testing"

	valueobjects "github.com/kalilventura/vehicle-management/internal/vehicles/domain/value-objects"
	"github.com/stretchr/testify/assert"
)

func TestCondition(t *testing.T) {
	t.Run("should create valid conditions", func(t *testing.T) {
		tests := []struct {
			name  string
			input string
			want  valueobjects.VehicleCondition
		}{
			{"new condition", "new", valueobjects.NewCondition()},
			{"used condition", "used", valueobjects.UsedCondition()},
			{"demonstration condition", "demonstration", valueobjects.DemonstrationCondition()},

			{"uppercase NEW", "NEW", valueobjects.NewCondition()},
			{"mixed case UsEd", "UsEd", valueobjects.UsedCondition()},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// given & when
				condition, err := valueobjects.NewVehicleCondition(tt.input)

				// then
				assert.NoError(t, err)
				assert.Equal(t, tt.want, condition)
				assert.Equal(t, strings.ToLower(tt.input), condition.Value())
			})
		}
	})

	t.Run("should return error for invalid conditions", func(t *testing.T) {
		tests := []struct {
			name  string
			input string
		}{
			{"empty string", ""},
			{"unknown condition", "refurbished"},
			{"invalid variation", "new2"},
			{"space padded", " used "},
			{"numeric", "123"},
			{"special chars", "new!"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// given & when
				_, err := valueobjects.NewVehicleCondition(tt.input)

				// then
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "invalid condition")
			})
		}
	})

	t.Run("should expose correct constant values", func(t *testing.T) {
		assert.Equal(t, "new", valueobjects.ConditionNew)
		assert.Equal(t, "used", valueobjects.ConditionUsed)
		assert.Equal(t, "demonstration", valueobjects.ConditionDemonstration)
	})

	t.Run("should be case insensitive when creating", func(t *testing.T) {
		// given & when
		c1, _ := valueobjects.NewVehicleCondition("NEW")
		c2, _ := valueobjects.NewVehicleCondition("new")
		c3, _ := valueobjects.NewVehicleCondition("New")

		// then
		assert.Equal(t, valueobjects.ConditionNew, c1)
		assert.Equal(t, c1, c2)
		assert.Equal(t, c2, c3)
		assert.Equal(t, "new", c1.Value())
	})
}
