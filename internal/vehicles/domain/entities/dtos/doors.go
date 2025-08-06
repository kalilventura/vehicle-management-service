package dtos

import "errors"

type Doors int

func NewDoors(value int) (Doors, error) {
	if value < 2 || value > 5 {
		return 0, errors.New("door count must be between 2 and 5")
	}
	return Doors(value), nil
}

func (m Doors) Value() int {
	return int(m)
}
