package valueobjects

import (
	"fmt"
	"time"

	"github.com/kalilventura/vehicle-management/internal/shared/domain"
)

type YearProps struct {
	Value int
}

// Year represents a vehicle year
type Year struct {
	domain.BaseValueObject[YearProps]
}

// NewYear creates a new Year value object
func NewYear(value int) (Year, error) {
	currentYear := time.Now().Year()
	if value < 1900 || value > currentYear+1 {
		return Year{}, fmt.Errorf("year must be between 1900 and %d", currentYear+1)
	}
	return Year{
		BaseValueObject: domain.NewBaseValueObject(YearProps{Value: value}),
	}, nil
}

// Value returns the year value
func (y Year) Value() int {
	return y.Props().Value
}

