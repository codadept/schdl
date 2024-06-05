// Package util provides utility functions for the schdl application.
package util

import (
	"errors"
	"time"
)

// validateTitle checks if the title length is within limits.
func validateTitle(title string) error {
	if len(title) > 150 {
		return errors.New("title should be less than 150 characters")
	}
	return nil
}

// validateDescription checks if the description length is within limits.
func validateDescription(description string) error {
	if len(description) > 512 {
		return errors.New("description should be less than 512 characters")
	}
	return nil
}

// validateDueDate checks if the due date is in the future.
func validateDueDate(dueDate time.Time) error {
	now := time.Now()
	if dueDate.Before(now) {
		return errors.New("due date cannot be in the past")
	}
	return nil
}

// validatePriority checks if the priority is valid.
func validatePriority(priority Priority) error {
	switch priority {
	case High, Medium, Low:
		return nil
	default:
		return errors.New("invalid priority")
	}
}

// ValidateTask validates a task by checking its title, description, due date, and priority.
func ValidateTask(task *Task) error {
	if err := validateTitle(task.Title); err != nil {
		return err
	}
	if err := validateDescription(task.Description); err != nil {
		return err
	}
	if err := validateDueDate(task.DueDate); err != nil {
		return err
	}
	if err := validatePriority(task.Priority); err != nil {
		return err
	}
	return nil
}
