package dtos

import "errors"

type Mileage int

func NewMileage(value int) (Mileage, error) {
	if value < 0 {
		return 0, errors.New("mileage cannot be negative")
	}
	return Mileage(value), nil
}

func (m Mileage) Value() int {
	return int(m)
}
