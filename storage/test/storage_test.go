package storage

import (
	"io/ioutil"
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
	bolt "go.etcd.io/bbolt"
)

func createTestConfig(t *testing.T) *Config {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		t.Errorf("Temp file creation failed")
	}
	log.Trace("Creating temp db")

	path := f.Name()
	f.Close()
	os.Remove(path)

	return NewConfig(path, 0600)
}

func createTestStorage(t *testing.T) *Storage {
	storage := NewStorage(createTestConfig(t))
	var err error

	if err = storage.Open(); err != nil {
		t.Errorf("Opening failed")
	}

	if err = storage.createBucket(BOARD_BUCKET); err != nil {
		t.Errorf("Bucket creation failed: %s", BOARD_BUCKET)
	}

	if err = storage.createBucket(LIST_BUCKET); err != nil {
		t.Errorf("Bucket creation failed: %s", LIST_BUCKET)
	}

	if err = storage.createBucket(QUADO_BUCKET); err != nil {
		t.Errorf("Bucket creation failed: %s", QUADO_BUCKET)
	}

	return storage
}

func destroyTestStorage(s *Storage, t *testing.T) {
	defer os.Remove(s.db.Path())

	err := s.db.Close()
	if err != nil {
		t.Errorf("Destroying failed")
	}

	log.Trace("Removing temp db")
}

func (storage *Storage) bucketSize(name string) (size int) {
	storage.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(name))

		cursor := bucket.Cursor()

		for key, _ := cursor.First(); key != nil; key, _ = cursor.Next() {
			size++
		}

		return nil
	})
	return
}

func TestOpen(t *testing.T) {
	storage := NewStorage(createTestConfig(t))

	err := storage.Open()
	if err != nil {
		t.Errorf("Opening failed")
	}
	log.Trace("Opening db")

	destroyTestStorage(storage, t)
}

func TestClose(t *testing.T) {
	storage := NewStorage(createTestConfig(t))

	err := storage.Open()
	if err != nil {
		t.Errorf("Opening failed")
	}

	err = storage.Close()
	if err != nil {
		t.Errorf("Closing failed")
	}
	log.Trace("Closing db")

	destroyTestStorage(storage, t)
}

func TestCreateBoardBucket(t *testing.T) {
	storage := NewStorage(createTestConfig(t))
	var err error

	if err = storage.Open(); err != nil {
		t.Errorf("Opening failed")
	}

	if err = storage.InitBuckets(); err != nil {
		t.Errorf("Bucket creation failed")
	}
	log.Trace("Creating buckets")

	if err = storage.Close(); err != nil {
		t.Errorf("Closing failed")
	}

	destroyTestStorage(storage, t)
}
