package dtos

import (
	"fmt"
	"strings"
)

type Transmission struct {
	value string
}

var (
	Manual    = Transmission{"manual"}
	Automatic = Transmission{"automatic"}
	CVT       = Transmission{"cvt"}
)

func NewTransmission(value string) (Transmission, error) {
	target := strings.ToLower(value)
	switch target {
	case Manual.value, Automatic.value, CVT.value:
		return Transmission{target}, nil
	default:
		return Transmission{}, fmt.Errorf("invalid transmission: %s", value)
	}
}

func (t Transmission) Value() string {
	return t.value
}
