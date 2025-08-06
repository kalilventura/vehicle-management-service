package dtos

import (
	"fmt"
	"time"
)

type Year int

func NewYear(value int) (Year, error) {
	currentYear := time.Now().Year()
	if value < 1900 || value > currentYear+1 {
		return 0, fmt.Errorf("year must be between 1900 and %d", currentYear+1)
	}
	return Year(value), nil
}

func (y Year) Value() int {
	return int(y)
}
