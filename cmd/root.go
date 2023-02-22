package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of CLI Task Manager",
	Long:  `All software has versions. This is Task's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CLI Task Manager by Talha UNAL v1.0")
	},
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "Manage your tasks simply from your terminal using task command.",
	Long: `Task is a CLI tool that helps you add, delete and list tasks easily from your terminal.

	-task add (add a new task)
	-task list (list all incomplete tasks)
	-task do (mark a task as complete)
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
