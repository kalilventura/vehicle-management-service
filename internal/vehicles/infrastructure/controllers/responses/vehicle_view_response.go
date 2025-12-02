package responses

import (
	"github.com/kalilventura/vehicle-management/internal/vehicles/application/dtos"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
)

// VehicleViewResponse
// @Description basic Vehicle information
type VehicleViewResponse struct {
	ID      string  `json:"id,omitempty"`
	Brand   string  `json:"brand,omitempty"`
	Price   float64 `json:"price,omitempty"`
	Model   string  `json:"model,omitempty"`
	Mileage int     `json:"mileage,omitempty"`
	Year    int     `json:"year,omitempty"`
} // @name VehicleViewResponse

func NewVehicleViewResponse(vehicle entities.Vehicle) VehicleViewResponse {
	return VehicleViewResponse{
		ID:      vehicle.ID,
		Brand:   vehicle.Brand,
		Price:   vehicle.GetPrice(),
		Model:   vehicle.Model,
		Mileage: vehicle.Specification.GetMileage(),
		Year:    vehicle.GetYear(),
	}
}

func NewVehicleViewResponseFromDTO(dto dtos.VehicleResponseDTO) VehicleViewResponse {
	return VehicleViewResponse{
		ID:      dto.ID,
		Brand:   dto.Brand,
		Price:   dto.Price,
		Model:   dto.Model,
		Mileage: dto.Mileage,
		Year:    dto.Year,
	}
}
