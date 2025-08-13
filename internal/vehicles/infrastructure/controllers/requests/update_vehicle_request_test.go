//go:build unit

package requests_test

import (
	"testing"

	"github.com/kalilventura/vehicle-management/internal/vehicles/infrastructure/controllers/requests"
	"github.com/stretchr/testify/assert"
)

func TestListVehiclesRequest(t *testing.T) {
	validStatus := "available"
	invalidStatus := "invalid_status"
	validSortBy := "price"
	validSortOrder := "asc"
	minPrice := 10000.0
	maxPrice := 50000.0
	invalidPrice := -100.0

	t.Run("successful conversion with all fields", func(t *testing.T) {
		qp := requests.ListVehiclesQueryParams{
			Status:    &validStatus,
			MinPrice:  &minPrice,
			MaxPrice:  &maxPrice,
			SortBy:    &validSortBy,
			SortOrder: &validSortOrder,
			Page:      1,
			Size:      10,
		}

		result, err := qp.ToDomain()

		assert.NoError(t, err)
		assert.Equal(t, validStatus, (*result.Status).Value())
		assert.Equal(t, minPrice, (*result.MinPrice).Value())
		assert.Equal(t, maxPrice, (*result.MaxPrice).Value())
		assert.Equal(t, validSortBy, result.SortBy)
		assert.Equal(t, validSortOrder, result.SortOrder)
		assert.Equal(t, 1, result.Pagination.Page)
		assert.Equal(t, 10, result.Pagination.Size)
	})

	t.Run("successful conversion with partial fields", func(t *testing.T) {
		qp := requests.ListVehiclesQueryParams{
			SortBy: &validSortBy,
			Page:   2,
			Size:   20,
		}

		result, err := qp.ToDomain()

		assert.NoError(t, err)
		assert.Nil(t, result.Status)
		assert.Nil(t, result.MinPrice)
		assert.Nil(t, result.MaxPrice)
		assert.Equal(t, validSortBy, result.SortBy)
		assert.Equal(t, "", result.SortOrder) // Not set
		assert.Equal(t, 2, result.Pagination.Page)
		assert.Equal(t, 20, result.Pagination.Size)
	})

	t.Run("invalid status", func(t *testing.T) {
		qp := requests.ListVehiclesQueryParams{
			Status: &invalidStatus,
			Page:   1,
			Size:   10,
		}

		_, err := qp.ToDomain()

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid status")
	})

	t.Run("invalid min price", func(t *testing.T) {
		qp := requests.ListVehiclesQueryParams{
			MinPrice: &invalidPrice,
			Page:     1,
			Size:     10,
		}

		_, err := qp.ToDomain()

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid min_price")
	})

	t.Run("invalid max price", func(t *testing.T) {
		qp := requests.ListVehiclesQueryParams{
			MaxPrice: &invalidPrice,
			Page:     1,
			Size:     10,
		}

		_, err := qp.ToDomain()

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid max_price")
	})

	t.Run("empty status string", func(t *testing.T) {
		emptyStatus := ""
		qp := requests.ListVehiclesQueryParams{
			Status: &emptyStatus,
			Page:   1,
			Size:   10,
		}

		result, err := qp.ToDomain()

		assert.NoError(t, err)
		assert.Nil(t, result.Status)
	})

	t.Run("zero pagination values", func(t *testing.T) {
		qp := requests.ListVehiclesQueryParams{
			Page: 0,
			Size: 0,
		}

		result, err := qp.ToDomain()

		assert.NoError(t, err)
		assert.Equal(t, 0, result.Pagination.Page)
		assert.Equal(t, 0, result.Pagination.Size)
	})

	t.Run("negative pagination values", func(t *testing.T) {
		qp := requests.ListVehiclesQueryParams{
			Page: -1,
			Size: -10,
		}

		result, err := qp.ToDomain()

		assert.NoError(t, err)
		assert.Equal(t, -1, result.Pagination.Page)
		assert.Equal(t, -10, result.Pagination.Size)
	})

	t.Run("nil pointers", func(t *testing.T) {
		qp := requests.ListVehiclesQueryParams{
			Status:    nil,
			MinPrice:  nil,
			MaxPrice:  nil,
			SortBy:    nil,
			SortOrder: nil,
			Page:      1,
			Size:      10,
		}

		result, err := qp.ToDomain()

		assert.NoError(t, err)
		assert.Nil(t, result.Status)
		assert.Nil(t, result.MinPrice)
		assert.Nil(t, result.MaxPrice)
		assert.Equal(t, "", result.SortBy)
		assert.Equal(t, "", result.SortOrder)
		assert.Equal(t, 1, result.Pagination.Page)
		assert.Equal(t, 10, result.Pagination.Size)
	})

	t.Run("empty string pointers", func(t *testing.T) {
		emptyString := ""
		qp := requests.ListVehiclesQueryParams{
			Status:    &emptyString,
			SortBy:    &emptyString,
			SortOrder: &emptyString,
			Page:      1,
			Size:      10,
		}

		result, err := qp.ToDomain()

		assert.NoError(t, err)
		assert.Nil(t, result.Status)
		assert.Equal(t, "", result.SortBy)
		assert.Equal(t, "", result.SortOrder)
	})
}
