package controllers

import "github.com/kalilventura/vehicle-management/internal/shared/domain/entities"

// PaginatedResponse Default object to return paginated page
// @Description Default object to return paginated page
type PaginatedResponse[T any] struct {
	Content          []T   `json:"content"`
	IsFirstPage      bool  `json:"first"`
	IsLastPage       bool  `json:"last"`
	PageNumber       int   `json:"pageNumber"`
	Size             int   `json:"pageSize"`
	NumberOfElements int   `json:"numberOfElements"`
	TotalElements    int64 `json:"totalElements"`
	TotalPages       int64 `json:"totalPages"`
} // @name PaginatedResponse

func NewPaginatedResponse[T any](
	content []T,
	pagination entities.Pagination) PaginatedResponse[T] {
	return PaginatedResponse[T]{
		Content:          content,
		IsFirstPage:      pagination.IsFirstPage(),
		IsLastPage:       pagination.IsLastPage(),
		PageNumber:       pagination.Page,
		Size:             pagination.Size,
		NumberOfElements: len(content),
		TotalElements:    pagination.TotalElements,
		TotalPages:       pagination.TotalPages(),
	}
}
