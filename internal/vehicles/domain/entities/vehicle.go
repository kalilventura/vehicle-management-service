package entities

import (
	"errors"
	"time"

	"github.com/kalilventura/vehicle-management/internal/shared/domain"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/events"
	valueobjects "github.com/kalilventura/vehicle-management/internal/vehicles/domain/value-objects"
)

// Vehicle represents a vehicle aggregate root
type Vehicle struct {
	domain.AggregateRoot
	brand         string
	model         string
	color         string
	description   string
	price         valueobjects.Price
	features      valueobjects.Features
	specification valueobjects.Specification
	status        valueobjects.VehicleStatus
	condition     valueobjects.VehicleCondition
	year          valueobjects.Year
}

// VehicleProps contains the properties needed to create a vehicle
type VehicleProps struct {
	Brand         string
	Model         string
	Color         string
	Description   string
	Price         valueobjects.Price
	Features      valueobjects.Features
	Specification valueobjects.Specification
	Status        valueobjects.VehicleStatus
	Condition     valueobjects.VehicleCondition
	Year          valueobjects.Year
}

// NewVehicle creates a new Vehicle aggregate root
func NewVehicle(props VehicleProps) (*Vehicle, error) {
	if props.Brand == "" {
		return nil, errors.New("brand is required")
	}
	if props.Model == "" {
		return nil, errors.New("model is required")
	}
	if props.Color == "" {
		return nil, errors.New("color is required")
	}

	vehicle := &Vehicle{
		AggregateRoot: domain.NewAggregateRoot(),
		brand:         props.Brand,
		model:         props.Model,
		color:         props.Color,
		description:   props.Description,
		price:         props.Price,
		features:      props.Features,
		specification: props.Specification,
		status:        props.Status,
		condition:     props.Condition,
		year:          props.Year,
	}

	// Validate business rules
	if err := vehicle.validate(); err != nil {
		return nil, err
	}

	// Add domain event
	vehicle.AddDomainEvent(events.NewVehicleCreatedEvent(
		vehicle.ID(),
		vehicle.brand,
		vehicle.model,
		vehicle.price.Amount(),
	))

	return vehicle, nil
}

// NewVehicleWithID creates a Vehicle with a specific ID (for reconstruction from persistence)
func NewVehicleWithID(
	id string,
	createdAt, updatedAt time.Time,
	props VehicleProps) (*Vehicle, error) {
	if props.Brand == "" {
		return nil, errors.New("brand is required")
	}
	if props.Model == "" {
		return nil, errors.New("model is required")
	}
	if props.Color == "" {
		return nil, errors.New("color is required")
	}

	vehicle := &Vehicle{
		AggregateRoot: domain.NewAggregateRootWithID(id, createdAt, updatedAt),
		brand:         props.Brand,
		model:         props.Model,
		color:         props.Color,
		description:   props.Description,
		price:         props.Price,
		features:      props.Features,
		specification: props.Specification,
		status:        props.Status,
		condition:     props.Condition,
		year:          props.Year,
	}

	// Validate business rules
	if err := vehicle.validate(); err != nil {
		return nil, err
	}

	return vehicle, nil
}

// validate validates business rules
func (v *Vehicle) validate() error {
	// Business rule: A new vehicle must have mileage equal to zero
	if v.condition.IsNew() && v.specification.MileageIsGreaterThan(0) {
		return errors.New("a new vehicle must have a mileage equal to zero")
	}
	return nil
}

// Brand returns the vehicle brand
func (v *Vehicle) Brand() string {
	return v.brand
}

// Model returns the vehicle model
func (v *Vehicle) Model() string {
	return v.model
}

// Color returns the vehicle color
func (v *Vehicle) Color() string {
	return v.color
}

// Description returns the vehicle description
func (v *Vehicle) Description() string {
	return v.description
}

// Price returns the vehicle price
func (v *Vehicle) Price() valueobjects.Price {
	return v.price
}

// Features returns the vehicle features
func (v *Vehicle) Features() valueobjects.Features {
	return v.features
}

// Specification returns the vehicle specification
func (v *Vehicle) Specification() valueobjects.Specification {
	return v.specification
}

// Status returns the vehicle status
func (v *Vehicle) Status() valueobjects.VehicleStatus {
	return v.status
}

// Condition returns the vehicle condition
func (v *Vehicle) Condition() valueobjects.VehicleCondition {
	return v.condition
}

// Year returns the vehicle year
func (v *Vehicle) Year() valueobjects.Year {
	return v.year
}

// Sell marks the vehicle as sold
func (v *Vehicle) Sell() error {
	if !v.status.CanBeSold() {
		return errors.New("vehicle cannot be sold in current status: " + v.status.Value())
	}

	v.status = valueobjects.Sold()
	v.Touch()

	// Add domain event
	v.AddDomainEvent(events.NewVehicleSoldEvent(v.ID(), v.price.Amount()))

	return nil
}

// Reserve reserves the vehicle
func (v *Vehicle) Reserve() error {
	if !v.status.CanBeReserved() {
		return errors.New("vehicle cannot be reserved in current status")
	}

	v.status = valueobjects.Reserved()
	v.Touch()

	return nil
}

// UpdatePrice updates the vehicle price
func (v *Vehicle) UpdatePrice(newPrice valueobjects.Price) error {
	if v.status.IsSold() {
		return errors.New("cannot update price of a sold vehicle")
	}

	v.price = newPrice
	v.Touch()

	v.AddDomainEvent(events.NewVehicleUpdatedEvent(v.ID()))

	return nil
}

// UpdateColor updates the vehicle color
func (v *Vehicle) UpdateColor(newColor string) error {
	if newColor == "" {
		return errors.New("color cannot be empty")
	}
	if v.status.IsSold() {
		return errors.New("cannot update color of a sold vehicle")
	}

	v.color = newColor
	v.Touch()

	v.AddDomainEvent(events.NewVehicleUpdatedEvent(v.ID()))

	return nil
}

// UpdateDescription updates the vehicle description
func (v *Vehicle) UpdateDescription(newDescription string) error {
	if v.status.IsSold() {
		return errors.New("cannot update description of a sold vehicle")
	}

	v.description = newDescription
	v.Touch()

	v.AddDomainEvent(events.NewVehicleUpdatedEvent(v.ID()))

	return nil
}
