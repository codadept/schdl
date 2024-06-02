package util

import (
	"time"
)

type Priority string

const (
	High   Priority = "high"
	Medium Priority = "medium"
	Low    Priority = "low"
)

type Task struct {
	ID          int
	Title       string
	Description string
	DueDate     time.Time
	Priority    Priority
}
