package model

import "time"

type TimeRange struct {
	From *time.Time `json:"from,omitempty"`
	To   *time.Time `json:"to,omitempty"`
}

type IntRange struct {
	From *int `json:"from,omitempty"`
	To   *int `json:"to,omitempty"`
}

type FloatRange struct {
	From *float64 `json:"from,omitempty"`
	To   *float64 `json:"to,omitempty"`
}

func GetIntRange(from *int, to *int) *IntRange {
	if from != nil && to != nil && *from > *to {
		return nil
	}
	return &IntRange{
		From: from,
		To:   to,
	}
}

func GetFloatRange(from *float64, to *float64) *FloatRange {
	if from != nil && to != nil && *from > *to {
		return nil
	}
	return &FloatRange{
		From: from,
		To:   to,
	}
}