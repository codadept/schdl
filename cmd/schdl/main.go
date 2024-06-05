package main

import (
	"fmt"

	"github.com/codadept/schdl/internal/storage"
	"github.com/codadept/schdl/internal/ui"
)

func main() {
	storage, err := storage.NewStorage("tasks.db")

	if err != nil {
		fmt.Println("Failed to initialize storage:", err)
		return
	}
	defer storage.Close()

	ui.Run(storage)
}
