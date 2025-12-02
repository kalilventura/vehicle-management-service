package factories

import (
	"github.com/kalilventura/vehicle-management/internal/vehicles/application/dtos"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/value-objects"
)

// ToResponseDTO converts a Vehicle entity to dtos.VehicleResponseDTO
func ToResponseDTO(vehicle *entities.Vehicle) dtos.VehicleResponseDTO {
	return dtos.VehicleResponseDTO{
		ID:                 vehicle.ID(),
		Brand:              vehicle.Brand(),
		Model:              vehicle.Model(),
		Color:              vehicle.Color(),
		Description:        vehicle.Description(),
		Price:              vehicle.Price().Amount(),
		Currency:           vehicle.Price().Currency(),
		BodyType:           vehicle.Specification().BodyType().Value(),
		Transmission:       vehicle.Specification().Transmission().Value(),
		FuelType:           vehicle.Specification().FuelType().Value(),
		Mileage:            vehicle.Specification().Mileage().Value(),
		Doors:              vehicle.Specification().Doors().Value(),
		Engine:             vehicle.Specification().Engine(),
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
		UpdatedAt:          vehicle.UpdatedAt(),
	}
}

// ToDomainEntity converts a CreateVehicleDTO to Vehicle entity
func ToDomainEntity(dto dtos.CreateVehicleDTO) (*entities.Vehicle, error) {
	price, err := valueobjects.NewPrice(dto.Price, "BRL")
	if err != nil {
		return nil, err
	}

	bodyType, err := valueobjects.NewBodyType(dto.BodyType)
	if err != nil {
		return nil, err
	}

	transmission, err := valueobjects.NewTransmission(dto.Transmission)
	if err != nil {
		return nil, err
	}

	fuelType, err := valueobjects.NewFuelType(dto.FuelType)
	if err != nil {
		return nil, err
	}

	mileage, err := valueobjects.NewMileage(dto.Mileage)
	if err != nil {
		return nil, err
	}

	doors, err := valueobjects.NewDoors(dto.Doors)
	if err != nil {
		return nil, err
	}

	year, err := valueobjects.NewYear(dto.Year)
	if err != nil {
		return nil, err
	}

	status, err := valueobjects.NewVehicleStatus(dto.Condition)
	if err != nil {
		// Default to available if not provided
		status = valueobjects.Available()
	}

	condition, err := valueobjects.NewVehicleCondition(dto.Condition)
	if err != nil {
		return nil, err
	}

	features := valueobjects.NewFeatures(valueobjects.FeaturesProps{
		HasAirConditioning: dto.HasAirConditioning,
		HasAirbag:          dto.HasAirbag,
		HasAbsBrakes:       dto.HasAbsBrakes,
		HasPowerSteering:   dto.HasPowerSteering,
		HasPowerWindows:    dto.HasPowerWindows,
		HasPowerLocks:      dto.HasPowerLocks,
		HasMultimedia:      dto.HasMultimedia,
		HasAlarm:           dto.HasAlarm,
		HasTractionControl: dto.HasTractionControl,
		HasRearCamera:      dto.HasRearCamera,
		HasParkingSensors:  dto.HasParkingSensors,
	})

	specification := valueobjects.NewSpecification(valueobjects.SpecificationProps{
		BodyType:     bodyType,
		Transmission: transmission,
		FuelType:     fuelType,
		Mileage:      mileage,
		Doors:        doors,
		Engine:       dto.Engine,
	})

	return entities.NewVehicle(entities.VehicleProps{
		Brand:         dto.Brand,
		Model:         dto.Model,
		Color:         dto.Color,
		Description:   dto.Description,
		Price:         price,
		Features:      features,
		Specification: specification,
		Status:        status,
		Condition:     condition,
		Year:          year,
	})
}
