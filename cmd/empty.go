package cmd

import (
	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var emptyCmd = &cobra.Command{
	Use:   "empty",
	Short: "Deletes all tasks.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		dirname, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		db, _ := bolt.Open(dirname+"/my.db", 0600, nil)
		defer db.Close()
		db.Update(func(tx *bolt.Tx) error {
			err := tx.DeleteBucket([]byte("TaskBucket"))
			if err != nil {
				return err
			}
			_, err = tx.CreateBucketIfNotExists([]byte("TaskBucket"))
			if err != nil {
				return err
			}
			return nil
		})
	},
}

func init() {
	rootCmd.AddCommand(emptyCmd)

}
