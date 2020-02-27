package storage

import (
	"io/ioutil"
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
)

func createTestConfig(t *testing.T) *Config {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		t.Errorf("Temp file creation failed")
	}
	log.Warn("Creating temp db")

	path := f.Name()
	f.Close()
	os.Remove(path)

	return &Config{path, 0600}
}

func destroyTestStorage(s *Storage, t *testing.T) {
	defer os.Remove(s.db.Path())
	err := s.db.Close()
	if err != nil {
		t.Errorf("Destroying failed")
	}
	log.Warn("Removing temp db")

}

func TestOpen(t *testing.T) {
	storage := newStorage(createTestConfig(t))

	err := storage.open()
	if err != nil {
		t.Errorf("Opening failed")
	}
	log.Info("Opening db")

	destroyTestStorage(storage, t)
}

func TestClose(t *testing.T) {
	storage := newStorage(createTestConfig(t))

	err := storage.open()
	if err != nil {
		t.Errorf("Opening failed")
	}

	err = storage.close()
	if err != nil {
		t.Errorf("Closing failed")
	}
	log.Info("Closing db")

	destroyTestStorage(storage, t)
}

func TestCreateBoardBucket(t *testing.T) {
	storage := newStorage(createTestConfig(t))
	var err error

	if err = storage.open(); err != nil {
		t.Errorf("Opening failed")
	}

	buckets := [3]string{BOARD_BUCKET, LIST_BUCKET, QUADO_BUCKET}
	for _, bucket := range buckets {
		if err = storage.createBucket(bucket); err != nil {
			t.Errorf("Bucket creation failed: %s", bucket)
		}
		log.WithField("bucket", bucket).Info("Creating bucket")
	}

	if err = storage.close(); err != nil {
		t.Errorf("Closing failed")
	}

	destroyTestStorage(storage, t)
}
