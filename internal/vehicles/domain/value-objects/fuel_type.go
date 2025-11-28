package valueobjects

import (
	"fmt"
	"strings"

	"github.com/kalilventura/vehicle-management/internal/shared/domain"
)

type FuelTypeProps struct {
	Value string
}

const (
	FuelTypeGasoline = "gasoline"
	FuelTypeEthanol  = "ethanol"
	FuelTypeDiesel   = "diesel"
	FuelTypeFlex     = "flex"
	FuelTypeHybrid   = "hybrid"
	FuelTypeElectric = "electric"
	FuelTypeGNV      = "gnv"
)

// FuelType represents a vehicle fuel type
type FuelType struct {
	domain.BaseValueObject[FuelTypeProps]
}

// NewFuelType creates a new FuelType value object
func NewFuelType(value string) (FuelType, error) {
	target := strings.ToLower(value)
	validTypes := []string{
		FuelTypeGasoline, FuelTypeEthanol, FuelTypeDiesel, FuelTypeFlex,
		FuelTypeHybrid, FuelTypeElectric, FuelTypeGNV,
	}

	for _, validType := range validTypes {
		if target == validType {
			return FuelType{
				BaseValueObject: domain.NewBaseValueObject(FuelTypeProps{Value: target}),
			}, nil
		}
	}

	return FuelType{}, fmt.Errorf("invalid fuel type: %s", value)
}

// Value returns the fuel type value
func (ft FuelType) Value() string {
	return ft.Props().Value
}

