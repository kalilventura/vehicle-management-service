package domain

import (
	"time"

	"github.com/google/uuid"
)

// BaseEntity provides common fields and behavior for all domain entities
type BaseEntity struct {
	id        string
	createdAt time.Time
	updatedAt time.Time
}

// NewBaseEntity creates a new BaseEntity with a generated UUID
func NewBaseEntity() BaseEntity {
	now := time.Now()
	return BaseEntity{
		id:        uuid.New().String(),
		createdAt: now,
		updatedAt: now,
	}
}

// NewBaseEntityWithID creates a new BaseEntity with a specific ID (useful for reconstruction from persistence)
func NewBaseEntityWithID(id string, createdAt, updatedAt time.Time) BaseEntity {
	return BaseEntity{
		id:        id,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

// ID returns the entity's unique identifier
func (e BaseEntity) ID() string {
	return e.id
}

// CreatedAt returns when the entity was created
func (e BaseEntity) CreatedAt() time.Time {
	return e.createdAt
}

// UpdatedAt returns when the entity was last updated
func (e BaseEntity) UpdatedAt() time.Time {
	return e.updatedAt
}

// Touch updates the updatedAt timestamp
func (e *BaseEntity) Touch() {
	e.updatedAt = time.Now()
}

// Equals checks if two entities are the same based on their ID
func (e BaseEntity) Equals(other BaseEntity) bool {
	if e.id == "" || other.id == "" {
		return false
	}
	return e.id == other.id
}

