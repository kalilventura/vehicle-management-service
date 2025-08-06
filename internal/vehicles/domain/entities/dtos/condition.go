package dtos

import (
	"fmt"
	"strings"
)

const (
	New           Condition = "new"
	Used          Condition = "used"
	Demonstration Condition = "demonstration"
)

type Condition string

func NewCondition(value string) (Condition, error) {
	target := strings.ToLower(value)
	switch target {
	case string(New):
		return New, nil
	case string(Used):
		return Used, nil
	case string(Demonstration):
		return Demonstration, nil
	default:
		return "", fmt.Errorf("invalid condition: %s", value)
	}
}

func (c Condition) Value() string {
	return string(c)
}
