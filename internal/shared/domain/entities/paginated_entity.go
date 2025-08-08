package entities

type PaginatedEntity[T any] struct {
	Content    []T
	Pagination Pagination
}

func NewPaginatedEntity[T any](content []T, pagination Pagination) PaginatedEntity[T] {
	return PaginatedEntity[T]{
		Content:    content,
		Pagination: pagination,
	}
}

func NewEmptyPaginatedEntity[T any](pagination Pagination) PaginatedEntity[T] {
	return PaginatedEntity[T]{
		Content:    []T{},
		Pagination: pagination,
	}
}
