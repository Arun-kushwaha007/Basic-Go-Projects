package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/Arun-kushwaha007/Basic-Go-Projects/tree/main/cli-todo/todo"
)

var doneCmd = &cobra.Command{
	Use:   "done [task number]",
	Short: "Mark a task as completed",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		index, err := strconv.Atoi(args[0])
		if err != nil || index < 1 {
			fmt.Println("Invalid task number.")
			return
		}

		tasks, err := todo.LoadTasks()
		if err != nil {
			fmt.Println("Failed to load tasks:", err)
			return
		}

		if index > len(tasks) {
			fmt.Println("Task number out of range.")
			return
		}

		tasks[index-1].Done = true
		todo.SaveTasks(tasks)
		fmt.Printf("Task %d marked as done.\n", index)
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
