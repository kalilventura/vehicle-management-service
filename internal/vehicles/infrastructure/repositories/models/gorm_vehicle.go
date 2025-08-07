package models

import (
  "time"

  "github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
  "gorm.io/gorm"
)

type GormVehicle struct {
  ID           string  `gorm:"primaryKey;type:uuid;default:uuidv7()"`
  Price        float64 `gorm:"not null;check:price >= 0"`
  Brand        string  `gorm:"type:varchar(30);not null"`
  Model        string  `gorm:"type:varchar(30);not null"`
  BodyType     string  `gorm:"type:varchar(30);not null"`
  Transmission string  `gorm:"type:varchar(30);not null"`
  FuelType     string  `gorm:"type:varchar(30);not null"`
  Color        string  `gorm:"type:varchar(30);not null"`
  Mileage      int     `gorm:"not null;check:mileage >= 0"`
  Engine       string  `gorm:"type:text;not null"`
  Doors        int     `gorm:"not null;check:doors BETWEEN 2 AND 5"`
  Year         int     `gorm:"not null;check:year >= 1900 AND year <= EXTRACT(YEAR FROM CURRENT_DATE) + 1"`

  // Features
  HasAirConditioning bool `gorm:"not null;default:false"`
  HasAirbag          bool `gorm:"not null;default:false"`
  HasAbsBrakes       bool `gorm:"not null;default:false"`
  HasPowerSteering   bool `gorm:"not null;default:false"`
  HasPowerWindows    bool `gorm:"not null;default:false"`
  HasPowerLocks      bool `gorm:"not null;default:false"`
  HasMultimedia      bool `gorm:"not null;default:false"`
  HasAlarm           bool `gorm:"not null;default:false"`
  HasTractionControl bool `gorm:"not null;default:false"`
  HasRearCamera      bool `gorm:"not null;default:false"`
  HasParkingSensors  bool `gorm:"not null;default:false"`

  // metadata
  Status      string `gorm:"type:varchar(20);not null;default:'available'"`
  Description string `gorm:"type:text"`
  Condition   string `gorm:"type:varchar(20)"`

  CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
  UpdatedAt *time.Time
}

func (GormVehicle) TableName() string {
  return "vehicle"
}

func (gv GormVehicle) BeforeCreate(_ *gorm.DB) (err error) {
  gv.CreatedAt = time.Now()
  gv.UpdatedAt = nil
  return nil
}

func (gv GormVehicle) BeforeUpdate(_ *gorm.DB) (err error) {
  now := time.Now()
  gv.UpdatedAt = &now
  return nil
}

func FromDomain(vehicle *entities.Vehicle) GormVehicle {
  return GormVehicle{
    Price:              vehicle.GetPrice(),
    Brand:              vehicle.Brand,
    Model:              vehicle.Model,
    Color:              vehicle.Color,
    Status:             vehicle.GetStatus(),
    Description:        vehicle.Description,
    Condition:          vehicle.GetCondition(),
    BodyType:           vehicle.Specification.GetBodyType(),
    Transmission:       vehicle.Specification.GetTransmission(),
    FuelType:           vehicle.Specification.GetFuelType(),
    Mileage:            vehicle.Specification.GetMileage(),
    Engine:             vehicle.Specification.GetEngine(),
    Doors:              vehicle.Specification.GetDoors(),
    Year:               vehicle.GetYear(),
    HasAirConditioning: vehicle.Features.HasAirConditioning,
    HasAirbag:          vehicle.Features.HasAirbag,
    HasAbsBrakes:       vehicle.Features.HasAbsBrakes,
    HasPowerSteering:   vehicle.Features.HasPowerSteering,
    HasPowerWindows:    vehicle.Features.HasPowerWindows,
    HasPowerLocks:      vehicle.Features.HasPowerLocks,
    HasMultimedia:      vehicle.Features.HasMultimedia,
    HasAlarm:           vehicle.Features.HasAlarm,
    HasTractionControl: vehicle.Features.HasTractionControl,
    HasRearCamera:      vehicle.Features.HasRearCamera,
    HasParkingSensors:  vehicle.Features.HasParkingSensors,
  }
}
