package responses

type VehicleViewResponse struct {
	ID      string  `json:"id"`
	Brand   string  `json:"brand"`
	Price   float64 `json:"price"`
	Model   string  `json:"model"`
	Mileage string  `json:"mileage"`
	Year    int     `json:"year"`
}
