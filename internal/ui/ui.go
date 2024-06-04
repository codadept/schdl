package ui

import (
	"fmt"

	"github.com/codadept/schdl/internal/storage"
	"github.com/manifoldco/promptui"
)

// Run the Terminal UI Prompt
func Run(storage *storage.Storage) {
	for {
		prompt := promptui.Select{
			Label: "Select Action",
			Items: []string{"Create Task", "List Tasks", "Update Task", "Delete Task", "Exit"},
		}

		_, action, err := prompt.Run()
		if err != nil {
			fmt.Println("Prompt failed:", err)
			return
		}

		switch action {
		case "Create Task":
			CreateTask(storage)
		case "List Tasks":
			ListTasks(storage)
		case "Update Task":
			UpdateTask(storage)
		case "Delete Task":
			DeleteTask(storage)
		case "Exit":
			return
		}
	}
}
