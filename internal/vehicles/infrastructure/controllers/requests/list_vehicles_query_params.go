package requests

import (
	"fmt"

	"github.com/kalilventura/vehicle-management/internal/shared/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities/dtos"
)

type ListVehiclesQueryParams struct {
	Status    *string  `query:"status"`
	MinPrice  *float64 `query:"min_price"`
	MaxPrice  *float64 `query:"max_price"`
	SortBy    *string  `query:"sort_by"`
	SortOrder *string  `query:"sort_order"`
	Page      int      `query:"page"`
	Size      int      `query:"size"`
}

func (qp ListVehiclesQueryParams) ToDomain() (*dtos.ListVehiclesInput, error) {
	input := &dtos.ListVehiclesInput{}
	if qp.SortBy != nil {
		input.SortBy = *qp.SortBy
	}
	if qp.SortOrder != nil {
		input.SortOrder = *qp.SortOrder
	}

	if qp.Status != nil && *qp.Status != "" {
		status, err := dtos.NewStatus(*qp.Status)
		if err != nil {
			return nil, fmt.Errorf("invalid status: %w", err)
		}
		input.Status = &status
	}

	if qp.MinPrice != nil {
		minPrice, err := dtos.NewPrice(*qp.MinPrice)
		if err != nil {
			return nil, fmt.Errorf("invalid min_price: %w", err)
		}
		input.MinPrice = &minPrice
	}

	if qp.MaxPrice != nil {
		maxPrice, err := dtos.NewPrice(*qp.MaxPrice)
		if err != nil {
			return nil, fmt.Errorf("invalid max_price: %w", err)
		}
		input.MaxPrice = &maxPrice
	}

	input.Pagination = entities.Pagination{
		Page: qp.Page,
		Size: qp.Size,
	}

	return input, nil
}
