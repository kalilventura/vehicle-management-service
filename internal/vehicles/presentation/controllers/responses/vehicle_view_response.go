package responses

import (
  "github.com/kalilventura/vehicle-management/internal/vehicles/application/dtos"
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
