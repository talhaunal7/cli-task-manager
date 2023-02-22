package cmd

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
)

func isEmpty(stats bolt.BucketStats) bool {
	if stats.KeyN == 0 {
		return true
	}
	return false
}

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task as complete",
	Args:  cobra.ExactArgs(1),
	Long:  `task do takes exact 1 arguments for removing a task from the list`,
	Run: func(cmd *cobra.Command, args []string) {
		dirname, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}

		db, _ := bolt.Open(dirname+"/my.db", 0600, nil)
		defer db.Close()
		//Iterate over the bucket, delete the desired key
		err = db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("TaskBucket"))
			c := b.Cursor()
			goalKey, _ := strconv.Atoi(args[0])
			i := 1
			for k, _ := c.First(); k != nil; k, _ = c.Next() {
				if goalKey == i {
					b.Delete(k)
					break
				}
				i++
			}
			return nil
		})
		// Check if bucket is empty, if empty then set the sequence to zero
		err = db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("TaskBucket"))
			if stats := b.Stats(); isEmpty(stats) {
				b.SetSequence(0)
			}
			return nil
		})
		if err != nil {
			fmt.Println(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(doCmd)

}
