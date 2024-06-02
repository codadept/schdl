// Package task provides functionality for Creating, Updating, Deleting and Listing all the task from the storage source
package task

import (
	"github.com/codadept/schdl/internal/storage"
	"github.com/codadept/schdl/internal/util"
)

// CreateTask creates a new task using the provided storage and task details.
// It validates the new task and adds it to the storage.
func CreateTask(storage *storage.Storage, newTask *util.Task) (*util.Task, error) {
	// Validate the new task
	if err := util.ValidateTask(newTask); err != nil {
		return nil, err
	}

	// Add the task to the storage
	taskID, err := storage.AddTask(newTask)
	if err != nil {
		return nil, err
	}

	newTask.ID = taskID
	return newTask, nil
}

// DeleteTask deletes a task with the specified ID from the storage.
func DeleteTask(storage *storage.Storage, id int) error {
	return storage.DeleteTask(id)
}

// UpdateTask updates an existing task with the provided updated task details in the storage.
func UpdateTask(storage *storage.Storage, updatedTask *util.Task) error {
	// Validate the updated task
	if err := util.ValidateTask(updatedTask); err != nil {
		return err
	}

	return storage.UpdateTask(updatedTask)
}

// GetListOfTasks retrieves a list of all tasks stored in the storage.
func GetListOfTasks(storage *storage.Storage) ([]util.Task, error) {
	return storage.ListTasks()
}
