package factories

import (
	"time"

	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
	valueobjects "github.com/kalilventura/vehicle-management/internal/vehicles/domain/value-objects"
	"github.com/kalilventura/vehicle-management/internal/vehicles/infrastructure/persistence/models"
)

// ToOrmEntity converts a Vehicle domain entity to GormVehicle
func ToOrmEntity(vehicle *entities.Vehicle) *models.GormVehicle {
	updatedAt := vehicle.UpdatedAt()
	var updatedAtPtr *time.Time
	if !updatedAt.IsZero() {
		updatedAtPtr = &updatedAt
	}

	return &models.GormVehicle{
		ID:                 vehicle.ID(),
		Brand:              vehicle.Brand(),
		Model:              vehicle.Model(),
		Color:              vehicle.Color(),
		Description:        vehicle.Description(),
		Price:              vehicle.Price().Amount(),
		BodyType:           vehicle.Specification().BodyType().Value(),
		Transmission:       vehicle.Specification().Transmission().Value(),
		FuelType:           vehicle.Specification().FuelType().Value(),
		Mileage:            vehicle.Specification().Mileage().Value(),
		Engine:             vehicle.Specification().Engine(),
		Doors:              vehicle.Specification().Doors().Value(),
		Year:               vehicle.Year().Value(),
		Status:             vehicle.Status().Value(),
		Condition:          vehicle.Condition().Value(),
		HasAirConditioning: vehicle.Features().HasAirConditioning(),
		HasAirbag:          vehicle.Features().HasAirbag(),
		HasAbsBrakes:       vehicle.Features().HasAbsBrakes(),
		HasPowerSteering:   vehicle.Features().HasPowerSteering(),
		HasPowerWindows:    vehicle.Features().HasPowerWindows(),
		HasPowerLocks:      vehicle.Features().HasPowerLocks(),
		HasMultimedia:      vehicle.Features().HasMultimedia(),
		HasAlarm:           vehicle.Features().HasAlarm(),
		HasTractionControl: vehicle.Features().HasTractionControl(),
		HasRearCamera:      vehicle.Features().HasRearCamera(),
		HasParkingSensors:  vehicle.Features().HasParkingSensors(),
		CreatedAt:          vehicle.CreatedAt(),
		UpdatedAt:          updatedAtPtr,
	}
}

// ToDomainEntity converts a GormVehicle to Vehicle domain entity
func ToDomainEntity(ormEntity *models.GormVehicle) (*entities.Vehicle, error) {
	// Create value objects
	price, err := valueobjects.NewPrice(ormEntity.Price, "BRL")
	if err != nil {
		return nil, err
	}

	bodyType, err := valueobjects.NewBodyType(ormEntity.BodyType)
	if err != nil {
		return nil, err
	}

	transmission, err := valueobjects.NewTransmission(ormEntity.Transmission)
	if err != nil {
		return nil, err
	}

	fuelType, err := valueobjects.NewFuelType(ormEntity.FuelType)
	if err != nil {
		return nil, err
	}

	mileage, err := valueobjects.NewMileage(ormEntity.Mileage)
	if err != nil {
		return nil, err
	}

	doors, err := valueobjects.NewDoors(ormEntity.Doors)
	if err != nil {
		return nil, err
	}

	year, err := valueobjects.NewYear(ormEntity.Year)
	if err != nil {
		return nil, err
	}

	status, err := valueobjects.NewVehicleStatus(ormEntity.Status)
	if err != nil {
		return nil, err
	}

	condition, err := valueobjects.NewVehicleCondition(ormEntity.Condition)
	if err != nil {
		return nil, err
	}

	features := valueobjects.NewFeatures(valueobjects.FeaturesProps{
		HasAirConditioning: ormEntity.HasAirConditioning,
		HasAirbag:          ormEntity.HasAirbag,
		HasAbsBrakes:       ormEntity.HasAbsBrakes,
		HasPowerSteering:   ormEntity.HasPowerSteering,
		HasPowerWindows:    ormEntity.HasPowerWindows,
		HasPowerLocks:      ormEntity.HasPowerLocks,
		HasMultimedia:      ormEntity.HasMultimedia,
		HasAlarm:           ormEntity.HasAlarm,
		HasTractionControl: ormEntity.HasTractionControl,
		HasRearCamera:      ormEntity.HasRearCamera,
		HasParkingSensors:  ormEntity.HasParkingSensors,
	})

	specification := valueobjects.NewSpecification(valueobjects.SpecificationProps{
		BodyType:     bodyType,
		Transmission: transmission,
		FuelType:     fuelType,
		Mileage:      mileage,
		Doors:        doors,
		Engine:       ormEntity.Engine,
	})

	var updatedAt time.Time
	if ormEntity.UpdatedAt != nil {
		updatedAt = *ormEntity.UpdatedAt
	}

	return entities.NewVehicleWithID(
		ormEntity.ID,
		ormEntity.CreatedAt,
		updatedAt,
		entities.VehicleProps{
			Brand:         ormEntity.Brand,
			Model:         ormEntity.Model,
			Color:         ormEntity.Color,
			Description:   ormEntity.Description,
			Price:         price,
			Features:      features,
			Specification: specification,
			Status:        status,
			Condition:     condition,
			Year:          year,
		},
	)
}
