package entities

import (
	"errors"

	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities/dtos"
)

type Vehicle struct {
	ID            string
	Brand         string
	Model         string
	Color         string
	Description   string
	Price         dtos.Price
	Features      Features
	Specification Specification
	Status        dtos.Status
	Condition     dtos.Condition
	Year          dtos.Year
}

func (v Vehicle) IsValid() error {
	if v.Condition == dtos.New && v.Specification.MileageIsGreaterThan(0) {
		return errors.New("a new vehicle must have a mileage greater than zero")
	}
	return nil
}

func (v Vehicle) GetPrice() float64 {
	return v.Price.Value()
}

func (v Vehicle) GetStatus() string {
	return v.Status.Value()
}

func (v Vehicle) GetCondition() string {
	return v.Condition.Value()
}

func (v Vehicle) GetYear() int {
	return v.Year.Value()
}
