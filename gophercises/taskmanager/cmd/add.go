package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}

// init is a function that runs before the main function - all init functions will get run before main func in main package
func init() {
	RootCmd.AddCommand(addCmd)
}
