//go:build unit

package dtos_test

import (
	"testing"

	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities/dtos"
	"github.com/stretchr/testify/assert"
)

func TestDoors(t *testing.T) {
	t.Run("should return an error when the door count is less than 2", func(t *testing.T) {
		// given & when
		_, err := dtos.NewDoors(0)
		// then
		assert.Error(t, err, "the door count should be greater than 2")
	})

	t.Run("should return an error when the door count is greater than 5", func(t *testing.T) {
		// given & when
		_, err := dtos.NewDoors(6)
		// then
		assert.Error(t, err, "the door count should be greater than 5")
	})

	t.Run("should create a new instance", func(t *testing.T) {
		// given & when
		doors, err := dtos.NewDoors(5)
		// then
		assert.Equal(t, 5, doors.Value(), "the value is correct")
		assert.Nil(t, err, "the door count should be greater than 5")
	})
}
