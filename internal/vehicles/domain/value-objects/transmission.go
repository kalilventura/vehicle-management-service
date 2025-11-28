package valueobjects

import (
	"fmt"
	"strings"

	"github.com/kalilventura/vehicle-management/internal/shared/domain"
)

type TransmissionProps struct {
	Value string
}

const (
	TransmissionManual    = "manual"
	TransmissionAutomatic = "automatic"
	TransmissionCVT       = "cvt"
)

// Transmission represents a vehicle transmission type
type Transmission struct {
	domain.BaseValueObject[TransmissionProps]
}

// NewTransmission creates a new Transmission value object
func NewTransmission(value string) (Transmission, error) {
	target := strings.ToLower(value)
	validTypes := []string{TransmissionManual, TransmissionAutomatic, TransmissionCVT}

	for _, validType := range validTypes {
		if target == validType {
			return Transmission{
				BaseValueObject: domain.NewBaseValueObject(TransmissionProps{Value: target}),
			}, nil
		}
	}

	return Transmission{}, fmt.Errorf("invalid transmission: %s", value)
}

// Value returns the transmission value
func (t Transmission) Value() string {
	return t.Props().Value
}

