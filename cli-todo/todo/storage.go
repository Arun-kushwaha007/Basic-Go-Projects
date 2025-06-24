
package todo

import (
	"encoding/json"
	"os"
	// "path/filepath"
)


var tasksFile = "tasks.json"

// SaveTasks saves tasks to a JSON file.
func SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(tasksFile, data, 0644)
}


// LoadTasks loads tasks from a JSON file.
func LoadTasks() ([]Task, error) {
	var tasks []Task

	if _, err := os.Stat(tasksFile); os.IsNotExist(err) {
		return tasks, nil // return empty if file doesn't exist
	}

	data, err := os.ReadFile(tasksFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &tasks)
	return tasks, err
}


// This code provides functions to save and load tasks from a JSON file.
// The `SaveTasks` function creates or overwrites a file named "tasks.json" and
// encodes the provided tasks into JSON format.
// The `LoadTasks` function attempts to open "tasks.json" and decode its contents
// into a slice of `Task` structs. If the file does not exist, it returns an empty slice.
// If any other error occurs while opening or decoding the file, it returns an error