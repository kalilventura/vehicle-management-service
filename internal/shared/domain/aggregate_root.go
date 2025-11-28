package domain

import "time"

// AggregateRoot represents the root of an aggregate
// It extends BaseEntity and manages domain events
type AggregateRoot struct {
	BaseEntity
	domainEvents []DomainEvent
}

// NewAggregateRoot creates a new AggregateRoot
func NewAggregateRoot() AggregateRoot {
	return AggregateRoot{
		BaseEntity:   NewBaseEntity(),
		domainEvents:  []DomainEvent{},
	}
}

// NewAggregateRootWithID creates a new AggregateRoot with a specific ID
func NewAggregateRootWithID(id string, createdAt, updatedAt time.Time) AggregateRoot {
	return AggregateRoot{
		BaseEntity:   NewBaseEntityWithID(id, createdAt, updatedAt),
		domainEvents:  []DomainEvent{},
	}
}

// DomainEvents returns a copy of all domain events
func (ar *AggregateRoot) DomainEvents() []DomainEvent {
	events := make([]DomainEvent, len(ar.domainEvents))
	copy(events, ar.domainEvents)
	return events
}

// AddDomainEvent adds a domain event to the aggregate
func (ar *AggregateRoot) AddDomainEvent(event DomainEvent) {
	ar.domainEvents = append(ar.domainEvents, event)
}

// ClearEvents removes all domain events from the aggregate
func (ar *AggregateRoot) ClearEvents() {
	ar.domainEvents = []DomainEvent{}
}

