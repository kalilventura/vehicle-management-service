package valueobjects

import (
	"fmt"
	"strings"

	"github.com/kalilventura/vehicle-management/internal/shared/domain"
)

type VehicleStatusProps struct {
	Status string
}

const (
	StatusAvailable   = "available"
	StatusReserved    = "reserved"
	StatusSold        = "sold"
	StatusMaintenance = "maintenance"
)

// VehicleStatus represents the status of a vehicle
type VehicleStatus struct {
	domain.BaseValueObject[VehicleStatusProps]
}

// NewVehicleStatus creates a new VehicleStatus value object
func NewVehicleStatus(status string) (VehicleStatus, error) {
	target := strings.ToLower(status)
	switch target {
	case StatusAvailable, StatusReserved, StatusSold, StatusMaintenance:
		return VehicleStatus{
			BaseValueObject: domain.NewBaseValueObject(VehicleStatusProps{Status: target}),
		}, nil
	default:
		return VehicleStatus{}, fmt.Errorf("invalid vehicle status: %s", status)
	}
}

// Available creates a VehicleStatus with available status
func Available() VehicleStatus {
	status, _ := NewVehicleStatus(StatusAvailable)
	return status
}

// Reserved creates a VehicleStatus with reserved status
func Reserved() VehicleStatus {
	status, _ := NewVehicleStatus(StatusReserved)
	return status
}

// Sold creates a VehicleStatus with sold status
func Sold() VehicleStatus {
	status, _ := NewVehicleStatus(StatusSold)
	return status
}

// Maintenance creates a VehicleStatus with maintenance status
func Maintenance() VehicleStatus {
	status, _ := NewVehicleStatus(StatusMaintenance)
	return status
}

// Value returns the status value
func (vs VehicleStatus) Value() string {
	return vs.Props().Status
}

// IsAvailable checks if the status is available
func (vs VehicleStatus) IsAvailable() bool {
	return vs.Value() == StatusAvailable
}

// IsReserved checks if the status is reserved
func (vs VehicleStatus) IsReserved() bool {
	return vs.Value() == StatusReserved
}

// IsSold checks if the status is sold
func (vs VehicleStatus) IsSold() bool {
	return vs.Value() == StatusSold
}

// IsMaintenance checks if the status is maintenance
func (vs VehicleStatus) IsMaintenance() bool {
	return vs.Value() == StatusMaintenance
}

// CanBeSold checks if the vehicle can be sold (must be available or reserved)
func (vs VehicleStatus) CanBeSold() bool {
	return vs.IsAvailable() || vs.IsReserved()
}

// CanBeReserved checks if the vehicle can be reserved (must be available)
func (vs VehicleStatus) CanBeReserved() bool {
	return vs.IsAvailable()
}

