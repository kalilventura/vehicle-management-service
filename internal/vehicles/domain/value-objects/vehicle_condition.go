package valueobjects

import (
	"fmt"
	"strings"

	"github.com/kalilventura/vehicle-management/internal/shared/domain"
)

type VehicleConditionProps struct {
	Condition string
}

const (
	ConditionNew           = "new"
	ConditionUsed          = "used"
	ConditionDemonstration = "demonstration"
)

// VehicleCondition represents the condition of a vehicle
type VehicleCondition struct {
	domain.BaseValueObject[VehicleConditionProps]
}

// NewVehicleCondition creates a new VehicleCondition value object
func NewVehicleCondition(condition string) (VehicleCondition, error) {
	target := strings.ToLower(condition)
	switch target {
	case ConditionNew, ConditionUsed, ConditionDemonstration:
		return VehicleCondition{
			BaseValueObject: domain.NewBaseValueObject(VehicleConditionProps{Condition: target}),
		}, nil
	default:
		return VehicleCondition{}, fmt.Errorf("invalid vehicle condition: %s", condition)
	}
}

// New creates a VehicleCondition with new condition
func NewCondition() VehicleCondition {
	condition, _ := NewVehicleCondition(ConditionNew)
	return condition
}

// Used creates a VehicleCondition with used condition
func UsedCondition() VehicleCondition {
	condition, _ := NewVehicleCondition(ConditionUsed)
	return condition
}

// Demonstration creates a VehicleCondition with demonstration condition
func DemonstrationCondition() VehicleCondition {
	condition, _ := NewVehicleCondition(ConditionDemonstration)
	return condition
}

// Value returns the condition value
func (vc VehicleCondition) Value() string {
	return vc.Props().Condition
}

// IsNew checks if the condition is new
func (vc VehicleCondition) IsNew() bool {
	return vc.Value() == ConditionNew
}

// IsUsed checks if the condition is used
func (vc VehicleCondition) IsUsed() bool {
	return vc.Value() == ConditionUsed
}

// IsDemonstration checks if the condition is demonstration
func (vc VehicleCondition) IsDemonstration() bool {
	return vc.Value() == ConditionDemonstration
}

