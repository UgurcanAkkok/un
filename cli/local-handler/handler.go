package localhandler

import (
	"fmt"
	"time"

	bolt "go.etcd.io/bbolt"
	un "uakkok.dev/un/common"
)

type BucketNames string
const (
  BucketTasks BucketNames = "tasks"
)

type LocalHandler struct {
  LocalDBFile string
  DB *bolt.DB
}

func (l *LocalHandler) Init() error {
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

func (l *LocalHandler) GetTasks() (un.Tasks, error) {
  var tasksResp un.Tasks
  err := l.DB.View(func(tx *bolt.Tx) error {
    b := tx.Bucket([]byte(BucketTasks))
    c := b.Cursor()

    for k,v := c.First(); k != nil; k,v = c.Next() {
      fmt.Printf("key=%s, value=%s\n", k, v)
    }
    return nil
  })
  if err != nil {
    return tasksResp, fmt.Errorf("Failed to get tasks: %w", err)
  }

  return tasksResp, nil
}

func (l *LocalHandler) Close() error {
  if err := l.DB.Close(); err != nil {
    return fmt.Errorf("Error closing DB connection: %w", err)
  }
  return nil
}
