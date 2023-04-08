package cmd

import (
	"github.com/spf13/cobra"
	"task/initialize"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task as complete",
	Args:  cobra.ExactArgs(1),
	Long:  `task do takes exact 1 arguments for removing a task from the list`,
	Run: func(cmd *cobra.Command, args []string) {
		ctl := initialize.GetController()
		ctl.Do(args[0])

	},
}

func init() {
	rootCmd.AddCommand(doCmd)

}
