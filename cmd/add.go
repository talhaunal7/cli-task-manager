package cmd

import (
	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
	"strconv"
	"task/initialize"
)

func getNewKeyValue(b *bolt.Bucket) []byte {
	c := b.Cursor()
	k, _ := c.Last()
	keyIntValue, _ := strconv.Atoi(string(k))
	keyIntValue += 1
	k = []byte(strconv.Itoa(keyIntValue))
	return k
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new task",
	Args:  cobra.MinimumNArgs(1),
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		ctl := initialize.GetController()
		ctl.Add(args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

}
