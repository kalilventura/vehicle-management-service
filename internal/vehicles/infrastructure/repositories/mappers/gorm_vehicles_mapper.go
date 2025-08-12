package mappers

import (
	"reflect"

	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/vehicles/infrastructure/repositories/models"
)

func MapToDomain(vehicle *entities.Vehicle) models.GormVehicle {
	return models.GormVehicle{
		Price:              vehicle.GetPrice(),
		Brand:              vehicle.Brand,
		Model:              vehicle.Model,
		Color:              vehicle.Color,
		Status:             vehicle.GetStatus(),
		Description:        vehicle.Description,
		Condition:          vehicle.GetCondition(),
		BodyType:           vehicle.Specification.GetBodyType(),
		Transmission:       vehicle.Specification.GetTransmission(),
		FuelType:           vehicle.Specification.GetFuelType(),
		Mileage:            vehicle.Specification.GetMileage(),
		Engine:             vehicle.Specification.GetEngine(),
		Doors:              vehicle.Specification.GetDoors(),
		Year:               vehicle.GetYear(),
		HasAirConditioning: vehicle.Features.HasAirConditioning,
		HasAirbag:          vehicle.Features.HasAirbag,
		HasAbsBrakes:       vehicle.Features.HasAbsBrakes,
		HasPowerSteering:   vehicle.Features.HasPowerSteering,
		HasPowerWindows:    vehicle.Features.HasPowerWindows,
		HasPowerLocks:      vehicle.Features.HasPowerLocks,
		HasMultimedia:      vehicle.Features.HasMultimedia,
		HasAlarm:           vehicle.Features.HasAlarm,
		HasTractionControl: vehicle.Features.HasTractionControl,
		HasRearCamera:      vehicle.Features.HasRearCamera,
		HasParkingSensors:  vehicle.Features.HasParkingSensors,
	}
}

func MapToDomainList(gormEntities []models.GormVehicle) []entities.Vehicle {
	list := make([]entities.Vehicle, len(gormEntities))
	for i, gormEntity := range gormEntities {
		mapped := gormEntity.ToDomain()
		list[i] = *mapped
	}
	return list
}

func MapToUpdate(input *entities.UpdateVehicleInput) models.GormVehicle {
	gormVehicle := &models.GormVehicle{
		ID: input.ID,
	}

	if input.Color != nil {
		gormVehicle.Color = *input.Color
	}
	if input.Description != nil {
		gormVehicle.Description = *input.Description
	}
	if input.Price != nil {
		gormVehicle.Price = input.Price.Value()
	}
	if input.Status != nil {
		gormVehicle.Status = input.Status.Value()
	}
	if input.Condition != nil {
		gormVehicle.Condition = input.Condition.Value()
	}
	mapFeatures(input, gormVehicle)
	return *gormVehicle
}

func mapFeatures(input *entities.UpdateVehicleInput, gormVehicle *models.GormVehicle) {
	if input.Features == nil {
		return
	}

	inVal := reflect.ValueOf(input.Features).Elem()
	outVal := reflect.ValueOf(gormVehicle).Elem()

	for i := 0; i < inVal.NumField(); i++ {
		field := inVal.Field(i)
		fieldName := inVal.Type().Field(i).Name

		if field.Kind() == reflect.Ptr && !field.IsNil() {
			outField := outVal.FieldByName(fieldName)
			if outField.IsValid() && outField.CanSet() {
				outField.SetBool(field.Elem().Bool())
			}
		}
	}
}
