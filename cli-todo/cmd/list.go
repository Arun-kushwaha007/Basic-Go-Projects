package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/Arun-kushwaha007/Basic-Go-Projects/tree/main/cli-todo/todo"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, _ := todo.LoadTasks()
		if len(tasks) == 0 {
			fmt.Println("No tasks.")
			return
		}
		for i, t := range tasks {
			status := " "
			if t.Done {
				status = "x"
			}
			fmt.Printf("%d. [%s] %s\n", i+1, status, t.Description)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
