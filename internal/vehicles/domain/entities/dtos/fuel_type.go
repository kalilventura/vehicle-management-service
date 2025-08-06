package dtos

import (
	"fmt"
	"strings"
)

type FuelType struct {
	value string
}

var (
	Gasoline = FuelType{"gasoline"}
	Ethanol  = FuelType{"ethanol"}
	Diesel   = FuelType{"diesel"}
	Flex     = FuelType{"flex"}
	Hybrid   = FuelType{"hybrid"}
	Electric = FuelType{"electric"}
	GNV      = FuelType{"gnv"}
)

func NewFuelType(value string) (FuelType, error) {
	target := strings.ToLower(value)
	switch target {
	case Gasoline.value, Ethanol.value, Diesel.value, Flex.value,
		Hybrid.value, Electric.value, GNV.value:
		return FuelType{target}, nil
	default:
		return FuelType{}, fmt.Errorf("invalid fuel type: %s", value)
	}
}

func (f FuelType) Value() string {
	return f.value
}
