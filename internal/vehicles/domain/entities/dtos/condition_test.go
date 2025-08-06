//go:build unit

package dtos_test

import (
	"strings"
	"testing"

	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities/dtos"
	"github.com/stretchr/testify/assert"
)

func TestCondition(t *testing.T) {
	t.Run("should create valid conditions", func(t *testing.T) {
		tests := []struct {
			name  string
			input string
			want  dtos.Condition
		}{
			{"new condition", "new", dtos.New},
			{"used condition", "used", dtos.Used},
			{"demonstration condition", "demonstration", dtos.Demonstration},

			{"uppercase NEW", "NEW", dtos.New},
			{"mixed case UsEd", "UsEd", dtos.Used},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// given & when
				condition, err := dtos.NewCondition(tt.input)

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
				_, err := dtos.NewCondition(tt.input)

				// then
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "invalid condition")
			})
		}
	})

	t.Run("should expose correct constant values", func(t *testing.T) {
		assert.Equal(t, "new", dtos.New.Value())
		assert.Equal(t, "used", dtos.Used.Value())
		assert.Equal(t, "demonstration", dtos.Demonstration.Value())
	})

	t.Run("should be case insensitive when creating", func(t *testing.T) {
		// given & when
		c1, _ := dtos.NewCondition("NEW")
		c2, _ := dtos.NewCondition("new")
		c3, _ := dtos.NewCondition("New")

		// then
		assert.Equal(t, dtos.New, c1)
		assert.Equal(t, c1, c2)
		assert.Equal(t, c2, c3)
		assert.Equal(t, "new", c1.Value())
	})
}
