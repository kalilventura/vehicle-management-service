package models

import (
	"time"

	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities"
	"github.com/kalilventura/vehicle-management/internal/vehicles/domain/entities/dtos"
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

func (gv GormVehicle) BeforeCreate(_ *gorm.DB) error {
	gv.CreatedAt = time.Now()
	gv.UpdatedAt = nil
	return nil
}

func (gv GormVehicle) BeforeUpdate(_ *gorm.DB) error {
	now := time.Now()
	gv.UpdatedAt = &now
	return nil
}

func (gv GormVehicle) ToDomain() *entities.Vehicle {
	features := entities.Features{
		HasAirConditioning: gv.HasAirConditioning,
		HasAirbag:          gv.HasAirbag,
		HasAbsBrakes:       gv.HasAbsBrakes,
		HasPowerSteering:   gv.HasPowerSteering,
		HasPowerWindows:    gv.HasPowerWindows,
		HasPowerLocks:      gv.HasPowerLocks,
		HasMultimedia:      gv.HasMultimedia,
		HasAlarm:           gv.HasAlarm,
		HasTractionControl: gv.HasTractionControl,
		HasRearCamera:      gv.HasRearCamera,
		HasParkingSensors:  gv.HasParkingSensors,
	}
	mileage, _ := dtos.NewMileage(gv.Mileage)
	doors, _ := dtos.NewDoors(gv.Doors)
	bodyType, _ := dtos.NewBodyType(gv.BodyType)
	transmission, _ := dtos.NewTransmission(gv.Transmission)
	fuelType, _ := dtos.NewFuelType(gv.FuelType)
	specification := entities.Specification{
		BodyType:     bodyType,
		Transmission: transmission,
		FuelType:     fuelType,
		Mileage:      mileage,
		Doors:        doors,
		Engine:       gv.Engine,
	}
	price, _ := dtos.NewPrice(gv.Price)
	condition, _ := dtos.NewCondition(gv.Condition)
	status, _ := dtos.NewStatus(gv.Status)
	year, _ := dtos.NewYear(gv.Year)

	return &entities.Vehicle{
		ID:            gv.ID,
		Brand:         gv.Brand,
		Model:         gv.Model,
		Color:         gv.Color,
		Description:   gv.Description,
		Price:         price,
		Features:      features,
		Specification: specification,
		Status:        status,
		Condition:     condition,
		Year:          year,
	}
}
