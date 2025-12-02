package mappers

import (
	"github.com/kalilventura/vehicle-management/internal/vehicles/application/dtos"
	"github.com/kalilventura/vehicle-management/internal/vehicles/presentation/controllers/responses"
)

// VehicleResponseMapper maps application DTOs to presentation responses
type VehicleResponseMapper struct{}

// NewVehicleResponseMapper creates a new VehicleResponseMapper
func NewVehicleResponseMapper() *VehicleResponseMapper {
	return &VehicleResponseMapper{}
}

// ToResponse converts a VehicleResponseDTO to VehicleResponse
func (m *VehicleResponseMapper) ToResponse(dto dtos.VehicleResponseDTO) responses.VehicleResponse {
	return responses.VehicleResponse{
		ID:            dto.ID,
		Brand:         dto.Brand,
		Model:         dto.Model,
		Color:         dto.Color,
		Description:   dto.Description,
		Price:         dto.Price,
		BodyType:      dto.BodyType,
		Transmission:  dto.Transmission,
		FuelType:      dto.FuelType,
		Mileage:       dto.Mileage,
		Doors:         dto.Doors,
		Engine:        dto.Engine,
		Year:          dto.Year,
		Status:        dto.Status,
		Condition:     dto.Condition,
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
		CreatedAt:     dto.CreatedAt,
		UpdatedAt:     dto.UpdatedAt,
	}
}

