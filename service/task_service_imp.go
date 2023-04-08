package service

import (
	"fmt"
	"github.com/boltdb/bolt"
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

func (tsk *TaskServiceImp) Add() error {
	err := tsk.db.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucketIfNotExists([]byte("TaskBucket"))
		task := strings.Join([]string{"asd ", "dea"}, " ")
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
