package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
		task := strings.Join(args, " ")
		fmt.Printf("Added '%s' to your task list", task)
	},
}

// init is a function that runs before the main function - all init functions will get run before main func in main package
func init() {
	RootCmd.AddCommand(addCmd)
}
