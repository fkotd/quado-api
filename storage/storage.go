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

func NewStorage(config *Config) *Storage {
	return &Storage{nil, config}
}

func NewConfig(path string, mode os.FileMode) *Config {
	return &Config{path, mode}
}

func (storage *Storage) Open() error {
	var err error
	storage.db, err = bolt.Open(storage.config.path, storage.config.mode, &bolt.Options{Timeout: 1 * time.Second})
	return err
}

func (storage *Storage) Close() error {
	return storage.db.Close()
}

func (storage *Storage) InitBuckets() error {
	buckets := [3]string{BOARD_BUCKET, LIST_BUCKET, QUADO_BUCKET}
	var err error
	for _, bucket := range buckets {
		err = storage.createBucket(bucket)
	}
	return err
}

func (storage *Storage) createBucket(name string) error {
	return storage.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(name))
		return err
	})
}
