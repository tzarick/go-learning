package cmd

import "github.com/spf13/cobra"

// RootCmd will stitch all the commands together
var RootCmd = &cobra.Command{
	Use:   "taskmanager",
	Short: "Taskmanager is a CLI task manager",
}
