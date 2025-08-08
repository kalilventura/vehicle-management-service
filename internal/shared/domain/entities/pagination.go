package entities

type Pagination struct {
	Page          int
	Size          int
	TotalElements int64
}

func (itself *Pagination) FinalIndex() int {
	return itself.Page*itself.Size - 1
}

func (itself *Pagination) Offset() int {
	return (itself.Page - 1) * itself.Size
}

func (itself *Pagination) IsFirstPage() bool {
	return itself.Page == 1
}

func (itself *Pagination) IsLastPage() bool {
	return int64(itself.Page*itself.Size) >= itself.TotalElements
}

func (itself *Pagination) TotalPages() int64 {
	limit := int64(itself.Size)

	if limit == 0 || itself.TotalElements == 0 {
		return 0
	}
	totalPages := itself.TotalElements / limit

	// If there's a remainder, add 1 to account for the extra page
	if itself.TotalElements%limit != 0 {
		totalPages++
	}
	return totalPages
}
