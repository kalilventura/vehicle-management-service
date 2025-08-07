package dtos

import "errors"

type Price float64

func NewPrice(value float64) (Price, error) {
	if value < 0 {
		return 0, errors.New("price cannot be negative")
	}
	return Price(value), nil
}

func (p Price) Value() float64 {
	return float64(p)
}
