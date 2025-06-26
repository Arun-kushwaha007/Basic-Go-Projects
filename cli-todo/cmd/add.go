package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/Arun-kushwaha007/Basic-Go-Projects/tree/main/cli-todo/todo"
)

var addCmd = &cobra.Command{
	Use:   "add [task description]",
	Short: "Add a new task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		desc := strings.Join(args, " ")
		tasks, _ := todo.LoadTasks()
		tasks = append(tasks, todo.Task{Description: desc, Done: false})
		todo.SaveTasks(tasks)
		fmt.Println("Task added.")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
