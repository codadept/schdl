package task

import (
	"github.com/codadept/schdl/internal/storage"
	"github.com/codadept/schdl/internal/util"
)

func CreateTask(storage *storage.Storage, newTask *util.Task) (*util.Task, error) {

	if err := util.ValidateTask(newTask); err != nil {
		return nil, err
	}

	taskID, err := storage.AddTask(newTask)
	if err != nil {
		return nil, err
	}

	newTask.ID = taskID

	return newTask, nil
}

func DeleteTask(storage *storage.Storage, id int) error {
	return storage.DeleteTask(id)
}

func UpdateTask(storage *storage.Storage, updatedTask *util.Task) error {
	if err := util.ValidateTask(updatedTask); err != nil {
		return err
	}

	return storage.UpdateTask(updatedTask)
}

func GetListOfTasks(storage *storage.Storage) ([]util.Task, error) {
	return storage.ListTasks()
}
