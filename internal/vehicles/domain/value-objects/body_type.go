package valueobjects

import (
	"fmt"
	"strings"

	"github.com/kalilventura/vehicle-management/internal/shared/domain"
)

type BodyTypeProps struct {
	Value string
}

const (
	BodyTypeSedan       = "sedan"
	BodyTypeHatchback   = "hatchback"
	BodyTypeSUV         = "suv"
	BodyTypePickup      = "pickup"
	BodyTypeCoupe       = "coupe"
	BodyTypeConvertible = "convertible"
	BodyTypeWagon       = "wagon"
	BodyTypeMinivan     = "minivan"
	BodyTypeFastback    = "fastback"
)

// BodyType represents a vehicle body type
type BodyType struct {
	domain.BaseValueObject[BodyTypeProps]
}

// NewBodyType creates a new BodyType value object
func NewBodyType(value string) (BodyType, error) {
	target := strings.ToLower(value)
	validTypes := []string{
		BodyTypeSedan, BodyTypeHatchback, BodyTypeSUV, BodyTypePickup,
		BodyTypeCoupe, BodyTypeConvertible, BodyTypeWagon, BodyTypeMinivan, BodyTypeFastback,
	}

	for _, validType := range validTypes {
		if target == validType {
			return BodyType{
				BaseValueObject: domain.NewBaseValueObject(BodyTypeProps{Value: target}),
			}, nil
		}
	}

	return BodyType{}, fmt.Errorf("invalid body type: %s", value)
}

// Value returns the body type value
func (bt BodyType) Value() string {
	return bt.Props().Value
}

