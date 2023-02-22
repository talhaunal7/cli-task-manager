package cmd

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the tasks",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		dirname, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}

		db, err := bolt.Open(dirname+"/my.db", 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("TaskBucket"))
			c := b.Cursor()
			i := 0
			for k, v := c.First(); k != nil; k, v = c.Next() {
				i++
				fmt.Printf("%d- %s\n", i, v)
			}
			return nil
		})
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

}
