package domain

import (
	"encoding/json"
)

// BaseValueObject provides common behavior for all value objects
// Value objects are immutable and compared by their attributes, not identity
type BaseValueObject[T any] struct {
	props T
}

// NewBaseValueObject creates a new BaseValueObject with the given properties
func NewBaseValueObject[T any](props T) BaseValueObject[T] {
	return BaseValueObject[T]{props: props}
}

// Props returns the value object's properties (read-only)
func (vo BaseValueObject[T]) Props() T {
	return vo.props
}

// Equals checks if two value objects are equal based on their properties
func (vo BaseValueObject[T]) Equals(other BaseValueObject[T]) bool {
	// Use JSON marshaling for deep comparison
	voJSON, err1 := json.Marshal(vo.props)
	otherJSON, err2 := json.Marshal(other.props)
	
	if err1 != nil || err2 != nil {
		return false
	}
	
	return string(voJSON) == string(otherJSON)
}

