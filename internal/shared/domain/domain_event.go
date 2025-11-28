package domain

import (
	"time"
)

// DomainEvent represents something that happened in the domain and is important for the business
type DomainEvent interface {
	OccurredOn() time.Time
	EventName() string
	AggregateID() string
}

// BaseDomainEvent provides common fields for all domain events
type BaseDomainEvent struct {
	occurredOn  time.Time
	eventName   string
	aggregateID string
}

// NewBaseDomainEvent creates a new BaseDomainEvent
func NewBaseDomainEvent(aggregateID string, eventName string) BaseDomainEvent {
	return BaseDomainEvent{
		occurredOn:  time.Now(),
		eventName:   eventName,
		aggregateID: aggregateID,
	}
}

// OccurredOn returns when the event occurred
func (e BaseDomainEvent) OccurredOn() time.Time {
	return e.occurredOn
}

// EventName returns the name of the event
func (e BaseDomainEvent) EventName() string {
	return e.eventName
}

// AggregateID returns the ID of the aggregate that generated the event
func (e BaseDomainEvent) AggregateID() string {
	return e.aggregateID
}

