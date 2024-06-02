package util

import (
	"errors"
	"time"
)

func validateTitle(title string) error {
	if len(title) > 150 {
		return errors.New("title should be less than 150 characters")
	}
	return nil
}

func validateDescription(description string) error {
	if len(description) > 512 {
		return errors.New("description should be less than 512 characters")
	}
	return nil
}

func validateDueDate(dueDate time.Time) error {
	now := time.Now()
	if dueDate.Before(now) {
		return errors.New("due date cannot be in the past")
	}
	return nil
}

func validatePriority(priority Priority) error {
	switch priority {
	case High, Medium, Low:
		return nil
	default:
		return errors.New("invalid priority")
	}
}

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
