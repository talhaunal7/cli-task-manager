package service

import (
	"fmt"
	"github.com/boltdb/bolt"
	"strconv"
	"strings"
)

type TaskServiceImp struct {
	db *bolt.DB
}

func NewTaskService(db *bolt.DB) TaskService {
	return &TaskServiceImp{
		db: db,
	}
}

func (tsk *TaskServiceImp) Add(args []string) error {
	err := tsk.db.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucketIfNotExists([]byte("TaskBucket"))
		task := strings.Join(args, " ")
		key, _ := b.NextSequence()
		err := b.Put([]byte(fmt.Sprintf("%04d", key)), []byte(task))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (tsk *TaskServiceImp) List() error {
	err := tsk.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("TaskBucket"))
		c := b.Cursor()
		i := 0
		for k, v := c.First(); k != nil; k, v = c.Next() {
			i++
			fmt.Printf("%d- %s\n", i, v)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func isEmpty(stats bolt.BucketStats) bool {
	if stats.KeyN == 0 {
		return true
	}
	return false
}

func (tsk *TaskServiceImp) Do(key string) error {
	err := tsk.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("TaskBucket"))
		c := b.Cursor()
		goalKey, _ := strconv.Atoi(key)
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
	err = tsk.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("TaskBucket"))
		if stats := b.Stats(); isEmpty(stats) {
			b.SetSequence(0)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
