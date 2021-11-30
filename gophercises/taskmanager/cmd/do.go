package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("do called")
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg) // ParseInt is used for more generic parsing (hex etc) - atoi = ascii to integer -> automatically assumes you are using base 10 etc
			if err != nil {
				fmt.Println("Failed to parse the arg:", arg)
			} else {
				ids = append(ids, id)
			}
		}
		fmt.Println(ids)
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
