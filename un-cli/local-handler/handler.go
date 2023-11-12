package localhandler

import (
	"fmt"
	"time"

	bolt "go.etcd.io/bbolt"
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

  return nil
}

func (l *LocalHandler) Close() error {
  if err := l.DB.Close(); err != nil {
    return fmt.Errorf("Error closing DB connection: %w", err)
  }
  return nil
}
