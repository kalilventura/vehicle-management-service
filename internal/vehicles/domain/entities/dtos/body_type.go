package dtos

import (
	"fmt"
	"strings"
)

type BodyType struct {
	value string
}

var (
	Sedan       = BodyType{"sedan"}
	Hatchback   = BodyType{"hatchback"}
	SUV         = BodyType{"suv"}
	Pickup      = BodyType{"pickup"}
	Coupe       = BodyType{"coupe"}
	Convertible = BodyType{"convertible"}
	Wagon       = BodyType{"wagon"}
	Minivan     = BodyType{"minivan"}
	Fastback    = BodyType{"fastback"}
)

func NewBodyType(value string) (BodyType, error) {
	target := strings.ToLower(value)
	switch target {
	case Sedan.value, Hatchback.value, SUV.value, Pickup.value,
		Coupe.value, Convertible.value, Wagon.value, Minivan.value, Fastback.value:
		return BodyType{target}, nil
	default:
		return BodyType{}, fmt.Errorf("invalid body type: %s", value)
	}
}

func (b BodyType) Value() string {
	return b.value
}
