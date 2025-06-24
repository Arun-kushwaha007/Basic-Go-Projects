// main.go
package main

import (
	"fmt"
	"os"
	"strconv"
	"github.com/Arun-kushwaha007/Basic-Go-Projects/tree/main/cli-todo/todo"
)

func main() {
	args := os.Args[1:]  //Gets all command-line arguments except the program name and stores them in args.

	if len(args) < 1 {
		fmt.Println("Usage: add <task> | list | done <num> | delete <num>")
		return
	}
// Checks if any arguments were provided.
// If not, prints usage instructions and exits.


	command := args[0]  //Stores the first argument (the command: add, list, done, or delete) in the variable command.

	switch command {
	case "add":
		if len(args) < 2 {
			fmt.Println("Usage: add <task description>")
			return
		}
		desc := args[1]
		tasks, _ := todo.LoadTasks()
		tasks = append(tasks, todo.Task{Description: desc, Done: false})
		todo.SaveTasks(tasks)
		fmt.Println("Task added.")
		// If the command is "add":
		// Checks if a task description is provided; if not, prints usage and exits.
		// Stores the description in desc.
		// Loads existing tasks from file.
		// Appends a new Task (with the provided description and Done set to false) to the tasks slice.
		// Saves the updated tasks slice back to file.
		// Prints confirmation.




	case "list":
		tasks, _ := todo.LoadTasks()
		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}
		for i, t := range tasks {
			status := " "
			if t.Done {
				status = "x"
			}
			fmt.Printf("%d. [%s] %s\n", i+1, status, t.Description)
		}
		// If the command is "list":
		// Loads tasks from file.
		// If no tasks exist, prints "No tasks found." and exits.
		// Otherwise, loops through each task:
		// If the task is done, status is "x"; otherwise, it's a blank space.
		// Prints each task with its number, completion status, and description.




		case "done":
			if len(args) < 2 {
				fmt.Println("Usage: done <task number>")
				return
			}
			index, err := strconv.Atoi(args[1])
			if err != nil || index < 1 {
				fmt.Println("Invalid task number")
				return
			}
			tasks, _ := todo.LoadTasks()
			if index > len(tasks) {
				fmt.Println("Task number out of range")
				return
			}
			tasks[index-1].Done = true
			todo.SaveTasks(tasks)
			fmt.Printf("Marked task %d as done.\n", index)
		


			
			case "delete":
				if len(args) < 2 {
					fmt.Println("Usage: delete <task number>")
					return
				}
				index, err := strconv.Atoi(args[1])
				if err != nil || index < 1 {
					fmt.Println("Invalid task number")
					return
				}
				tasks, _ := todo.LoadTasks()
				if index > len(tasks) {
					fmt.Println("Task number out of range")
					return
				}
				tasks = append(tasks[:index-1], tasks[index:]...)
				todo.SaveTasks(tasks)
				fmt.Printf("Deleted task %d.\n", index)
			



	default:
		fmt.Println("Unknown command. Use: add | list | done | delete")
	}

// ✅ Test it!
// go run main.go add "Sample Task"
// go run main.go list

// You should see:
// 1. [ ] Sample Task

}

