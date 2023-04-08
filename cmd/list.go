package cmd

import (
	"github.com/spf13/cobra"
	"task/initialize"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the tasks",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		ctl := initialize.GetController()
		ctl.List()

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

}
