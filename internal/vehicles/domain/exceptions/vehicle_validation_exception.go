package exceptions

import "fmt"

// VehicleValidationException is raised when vehicle validation fails
type VehicleValidationException struct {
	Field   string
	Message string
}

func (e VehicleValidationException) Error() string {
	if e.Field != "" {
		return fmt.Sprintf("validation error on field %s: %s", e.Field, e.Message)
	}
	return e.Message
}

// NewVehicleValidationException creates a new VehicleValidationException
func NewVehicleValidationException(field, message string) VehicleValidationException {
	return VehicleValidationException{
		Field:   field,
		Message: message,
	}
}

