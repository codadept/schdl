package ui

import (
	"fmt"
	"strconv"
	"time"

	"github.com/codadept/schdl/internal/storage"
	"github.com/codadept/schdl/internal/task"
	"github.com/codadept/schdl/internal/util"
)

// CreateTask handles the creation of tasks from user input
func CreateTask(storage *storage.Storage) {
	// Take input of the title
	title, err := util.PromptInput("Title")
	if err != nil {
		fmt.Println("Input failed:", err)
		return
	}

	// Take input of the description
	description, err := util.PromptInput("Description")
	if err != nil {
		fmt.Println("Input failed:", err)
		return
	}

	// Take input of the due date which is of YYYY-MM-DD HH:MM format
	dueDate, err := util.PromptInput("Due Date (YYYY-MM-DD HH:MM)")
	if err != nil {
		fmt.Println("Input failed:", err)
		return
	}

	// Parsing the date to match the format
	parsedDate, err := time.Parse("2006-01-02 15:04", dueDate)
	if err != nil {
		fmt.Println("Invalid date format:", err)
		return
	}

	// Take input of the priority
	priority, err := util.PromptSelect("Priority", []string{"high", "medium", "low"})
	if err != nil {
		fmt.Println("Selection failed:", err)
		return
	}

	newTask := &util.Task{
		Title:       title,
		Description: description,
		DueDate:     parsedDate,
		Priority:    util.Priority(priority),
	}

	_, err = task.CreateTask(storage, newTask)
	if err != nil {
		fmt.Println("Failed to create task:", err)
	} else {
		fmt.Println("Task created successfully")
	}
}

// ListTasks handles the listing of all the tasks from the database
func ListTasks(storage *storage.Storage) {
	tasks, err := storage.ListTasks()
	if err != nil {
		fmt.Println("Failed to list tasks:", err)
		return
	}

	// Printing all the tasks in proper format
	for _, task := range tasks {
		fmt.Printf("ID: %d, Title: %s, Description: %s, Due Date: %s, Priority: %s\n",
			task.ID, task.Title, task.Description, task.DueDate.Format("2006-01-02 15:04"), task.Priority)
	}
}

// UpdateTask handles the updation of the tasks given the task id
func UpdateTask(storage *storage.Storage) {
	// Take input of the ID of the task to update
	id, err := util.PromptInput("Task ID to update")
	if err != nil {
		fmt.Println("Input failed:", err)
		return
	}

	taskID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Invalid ID:", err)
		return
	}

	// Get the task for the given ID from database
	task, err := storage.GetTask(taskID)
	if err != nil {
		fmt.Println("Failed to get task:", err)
		return
	}
	if task == nil {
		fmt.Println("Task not found")
		return
	}

	// Take input of the Title, default is the old Title
	title, err := util.PromptInputWithDefault("Title", task.Title)
	if err != nil {
		fmt.Println("Input failed:", err)
		return
	}

	// Take input of the Description, default is the old Description
	description, err := util.PromptInputWithDefault("Description", task.Description)
	if err != nil {
		fmt.Println("Input failed:", err)
		return
	}

	// Take input of the due date, default is the old due date
	dueDate, err := util.PromptInputWithDefault("Due Date (YYYY-MM-DD HH:MM)", task.DueDate.Format("2006-01-02 15:04"))
	if err != nil {
		fmt.Println("Input failed:", err)
		return
	}

	// Parsing the due date to match the format
	parsedDate, err := time.Parse("2006-01-02 15:04", dueDate)
	if err != nil {
		fmt.Println("Invalid date format:", err)
		return
	}

	// Take input of the Priority, default is the old Priority
	priority, err := util.PromptSelectWithDefault("Priority", []string{"high", "medium", "low"}, string(task.Priority))
	if err != nil {
		fmt.Println("Selection failed:", err)
		return
	}

	task.Title = title
	task.Description = description
	task.DueDate = parsedDate
	task.Priority = util.Priority(priority)

	err = storage.UpdateTask(task)
	if err != nil {
		fmt.Println("Failed to update task:", err)
	} else {
		fmt.Println("Task updated successfully")
	}
}

// DeleteTask handles the deleting of the task with the given Id from the database
func DeleteTask(storage *storage.Storage) {
	// Take input of the task ID to delete
	id, err := util.PromptInput("Task ID to delete")
	if err != nil {
		fmt.Println("Input failed:", err)
		return
	}

	taskID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Invalid ID:", err)
		return
	}

	err = storage.DeleteTask(taskID)
	if err != nil {
		fmt.Println("Failed to delete task:", err)
	} else {
		fmt.Println("Task deleted successfully")
	}
}
