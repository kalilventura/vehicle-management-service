package models

import (
	"gorm.io/gorm"
	"time"

	"github.com/google/uuid"
)

type GormVehicle struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey"`
	Brand     string     `gorm:"type:text;not null"`
	Model     string     `gorm:"type:text;not null"`
	Year      int        `gorm:"not null;check:year >= 1900 AND year <= EXTRACT(YEAR FROM CURRENT_DATE) + 1"`
	Color     string     `gorm:"type:varchar(30);not null"`
	Price     float64    `gorm:"type:decimal;not null;check:price >= 0"`
	CreatedAt time.Time  `gorm:"type:timestamp;not null;default:current_timestamp"`
	UpdatedAt *time.Time `gorm:"type:timestamp"`
}

func (GormVehicle) TableName() string {
	return "vehicle"
}

func (gv GormVehicle) BeforeCreate(_ *gorm.DB) (err error) {
	gv.ID = uuid.New()
	gv.UpdatedAt = nil
	return
}

func (gv GormVehicle) BeforeUpdate(_ *gorm.DB) (err error) {
	now := time.Now()
	gv.UpdatedAt = &now
	return
}
