package localhandler

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
  tasksResp = un.Tasks { Items: []un.Task{} }
  err := l.DB.View(func(tx *bolt.Tx) error {
    b := tx.Bucket([]byte(BucketTasks))
    c := b.Cursor()

    for k,v := c.First(); k != nil; k,v = c.Next() {
      var id int
      var task un.Task
      var err error
      idBuffer := bytes.NewBuffer(k)
      dec := gob.NewDecoder(idBuffer)
      if err = dec.Decode(&id); err != nil {
        return fmt.Errorf("Failed to decode to int value: %w", err)
      }
      taskBuffer := bytes.NewBuffer(v)
      dec = gob.NewDecoder(taskBuffer)
      if err = dec.Decode(&task); err != nil {
        return fmt.Errorf("Failed to decode to Task value: %w", err)
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

func (l *LocalHandler) PostTasks(tasksData un.Tasks) error {
  err := l.DB.Update(func(tx *bolt.Tx) error {
    b := tx.Bucket([]byte(BucketTasks))
    for _,v := range tasksData.Items {
      tasksBytes, err := atob(v)
      if err != nil {
        return err
      }
      idBytes, err := atob(v.ID)
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

func (l *LocalHandler) Close() error {
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

func btoTask(b []byte) (un.Task, error) {
  buffer := bytes.NewBuffer(b)
  var result un.Task
  dec := gob.NewDecoder(buffer)
  if err := dec.Decode(&result); err != nil {
    return result, fmt.Errorf("Failed to decode value: %w", err)
  }
  return result, nil
}
