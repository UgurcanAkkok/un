package embeddedhandler

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"time"

	bolt "go.etcd.io/bbolt"
	un "uakkok.dev/un/common"
)

type BucketNames string

const (
	BucketTasks BucketNames = "tasks"
)

type EmbeddedHandler struct {
	LocalDBFile string
	DB          *bolt.DB
}

func (l *EmbeddedHandler) Init() error {
	db, err := bolt.Open(l.LocalDBFile, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return fmt.Errorf("Failed to open db: %w", err)
	}
	l.DB = db

	err = l.DB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(BucketTasks))
		if err != nil {
			return fmt.Errorf("Failed to initalize %v bucket: %w", BucketTasks, err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("Failed to initalize buckets: %w", err)
	}
	return nil
}

func (l *EmbeddedHandler) GetTasks() (un.Tasks, error) {
	var tasksResp un.Tasks
	tasksResp = un.Tasks{Items: []un.Task{}}
	err := l.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketTasks))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			var task un.Task
			if err := gob.NewDecoder(bytes.NewBuffer(v)).Decode(&task); err != nil {
				return fmt.Errorf("Failed to decode task from byte to Task: %w", err)
			}
			tasksResp.Items = append(tasksResp.Items, task)
		}
		return nil
	})
	if err != nil {
		return tasksResp, fmt.Errorf("Failed to get tasks: %w", err)
	}

	return tasksResp, nil
}

func (l *EmbeddedHandler) PostTasks(tasksData un.Tasks) error {
	err := l.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketTasks))
		for _, v := range tasksData.Items {
			// This returns an error only if the Tx is closed or not writeable.
			// That can't happen in an Update() call so I ignore the error check.
			id, _ := b.NextSequence()
			idBytes, err := atob(int(id))
			if err != nil {
				return err
			}
			v.ID = int(id)
			tasksBytes, err := atob(v)
			if err != nil {
				return err
			}
			b.Put(idBytes, tasksBytes)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("Failed to put tasks: %w", err)
	}
	return nil
}

func (l *EmbeddedHandler) Close() error {
	if err := l.DB.Close(); err != nil {
		return fmt.Errorf("Error closing DB connection: %w", err)
	}
	return nil
}

func atob(value any) ([]byte, error) {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	if err := enc.Encode(value); err != nil {
		return nil, fmt.Errorf("Failed to encode value to byte: %w", err)
	}
	return buffer.Bytes(), nil
}

