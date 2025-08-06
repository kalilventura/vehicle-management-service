package dtos

import (
	"fmt"
	"strings"
)

const (
	Available   Status = "available"
	Reserved    Status = "reserved"
	Sold        Status = "sold"
	Maintenance Status = "maintenance"
)

type Status string

func NewStatus(value string) (Status, error) {
	target := strings.ToLower(value)
	switch target {
	case string(Available):
		return Available, nil
	case string(Reserved):
		return Reserved, nil
	case string(Sold):
		return Sold, nil
	case string(Maintenance):
		return Maintenance, nil
	default:
		return "", fmt.Errorf("invalid status: %s", value)
	}
}

func (s Status) Value() string {
	return string(s)
}
