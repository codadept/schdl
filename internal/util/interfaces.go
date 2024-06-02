// Package util provides utility functions for the schdl application.
package util

import (
	"time"
)

// Priority represents the priority of a task.
type Priority string

const (
	// High priority task.
	High Priority = "high"
	// Medium priority task.
	Medium Priority = "medium"
	// Low priority task.
	Low Priority = "low"
)

// Task represents a task with its details.
type Task struct {
	// ID is the unique identifier of the task.
	ID int
	// Title is the title or name of the task.
	Title string
	// Description is the description or details of the task.
	Description string
	// DueDate is the deadline or due date of the task.
	DueDate time.Time
	// Priority is the priority level of the task.
	Priority Priority
}
