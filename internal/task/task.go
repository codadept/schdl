package task

import (
	"errors"
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
	Description string
	DueDate     time.Time
	Priority    Priority
}

func ValidatePriority(priority Priority) error {
	switch priority {
	case High, Medium, Low:
		return nil
	default:
		return errors.New("invalid priority")
	}
}

func createTask() {}

func deleteTask() {}

func updateTask() {}

func getListOfTasks() []Task {
	var task []Task = []Task{{1, "demo", time.Now(), "high"}}
	return task
}
