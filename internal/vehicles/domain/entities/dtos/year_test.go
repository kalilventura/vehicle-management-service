//go:build unit

package dtos_test

import (
	"testing"
	"time"

	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities/dtos"
	"github.com/stretchr/testify/assert"
)

func TestYear(t *testing.T) {
	currentYear := time.Now().Year()

	t.Run("should create valid year values", func(t *testing.T) {
		tests := []struct {
			name  string
			input int
			want  dtos.Year
		}{
			{"minimum valid year", 1900, dtos.Year(1900)},
			{"current year", currentYear, dtos.Year(currentYear)},
			{"next year", currentYear + 1, dtos.Year(currentYear + 1)},
			{"mid-century year", 1950, dtos.Year(1950)},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// given & when
				year, err := dtos.NewYear(tt.input)

				// then
				assert.NoError(t, err)
				assert.Equal(t, tt.want, year)
				assert.Equal(t, tt.input, year.Value())
			})
		}
	})

	t.Run("should return error for invalid year values", func(t *testing.T) {
		tests := []struct {
			name  string
			input int
			want  string
		}{
			{"year before 1900", 1899, "year must be between 1900 and"},
			{"year too far in future", currentYear + 2, "year must be between 1900 and"},
			{"negative year", -1, "year must be between 1900 and"},
			{"zero year", 0, "year must be between 1900 and"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// given & when
				_, err := dtos.NewYear(tt.input)

				// then
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.want)
			})
		}
	})

	t.Run("should correctly return integer value", func(t *testing.T) {
		// given
		testValue := 2023

		// when
		year, _ := dtos.NewYear(testValue)
		value := year.Value()

		// then
		assert.Equal(t, testValue, value)
		assert.IsType(t, 0, value)
	})

	t.Run("should handle edge cases", func(t *testing.T) {
		t.Run("minimum valid year (1900)", func(t *testing.T) {
			// given & when
			year, err := dtos.NewYear(1900)

			// then
			assert.NoError(t, err)
			assert.Equal(t, 1900, year.Value())
		})

		t.Run("maximum valid year (current+1)", func(t *testing.T) {
			// given & when
			year, err := dtos.NewYear(currentYear + 1)

			// then
			assert.NoError(t, err)
			assert.Equal(t, currentYear+1, year.Value())
		})
	})
}
