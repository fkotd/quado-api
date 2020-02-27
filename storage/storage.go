package storage

import (
	"os"
	"time"

	bolt "go.etcd.io/bbolt"
)

const (
	DB_NAME      = "quado.db"
	MODE         = 0600
	BOARD_BUCKET = "boards"
	LIST_BUCKET  = "lists"
	QUADO_BUCKET = "quados"
)

type Storage struct {
	db     *bolt.DB
	config *Config
}

type Config struct {
	path string
	mode os.FileMode
}

func newStorage(config *Config) *Storage {
	return &Storage{nil, config}
}

func (s *Storage) open() error {
	var err error
	s.db, err = bolt.Open(s.config.path, s.config.mode, &bolt.Options{Timeout: 1 * time.Second})
	return err
}

func (s *Storage) close() error {
	return s.db.Close()
}

func (s *Storage) createBucket(name string) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(name))
		return err
	})
}

func bucketSize(name string, s *Storage) (size int) {
	s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(name))

		c := b.Cursor()

		for key, _ := c.First(); key != nil; key, _ = c.Next() {
			size++
		}

		return nil
	})
	return
}
