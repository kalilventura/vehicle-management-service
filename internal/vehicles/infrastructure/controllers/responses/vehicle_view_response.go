package responses

import "github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"

// VehicleViewResponse
// @Description basic Vehicle information
type VehicleViewResponse struct {
	ID      string  `json:"id"`
	Brand   string  `json:"brand"`
	Price   float64 `json:"price"`
	Model   string  `json:"model"`
	Mileage int     `json:"mileage"`
	Year    int     `json:"year"`
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
