package storage

import (
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
)

func initQuado(t *testing.T) *Storage {
	storage := newStorage(createTestConfig(t))
	var err error

	if err = storage.open(); err != nil {
		t.Errorf("Opening failed")
	}

	if err = storage.createBucket(QUADO_BUCKET); err != nil {
		t.Errorf("Bucket creation failed: %s", QUADO_BUCKET)
	}

	return storage
}

func TestPutQuado(t *testing.T) {
	storage := initQuado(t)
	quado := newQuado(newList(newBoard().ID, "").ID, "test", "desc", time.Now())

	if err := putQuado(quado, storage); err != nil {
		t.Errorf("Quado insertion in bucket failed: %s", quado.ID)
	}
	log.WithFields(log.Fields{
		"id":     quado.ID,
		"ListId": quado.ListID,
		"title":  quado.Title,
		"desc":   quado.Description,
		"date":   quado.Date,
	}).Info("New quado")

	size := bucketSize(QUADO_BUCKET, storage)
	if size != 1 {
		t.Errorf("Quado bucket size error: %v", size)
	}

	destroyTestStorage(storage, t)
}

func TestGetQuado(t *testing.T) {
	storage := initQuado(t)
	quado := newQuado(newList(newBoard().ID, "").ID, "test", "desc", time.Now())

	if err := putQuado(quado, storage); err != nil {
		t.Errorf("Quado insertion in bucket failed: %s", quado.ID)
	}
	log.WithField("id", quado.ID).Info("New quado")

	resQuado, err := getQuado(quado.ID, storage)
	if err != nil {
		t.Errorf("Quado retreiving failed: %s", quado.ID)
	}
	if resQuado.ID != quado.ID {
		t.Errorf("Quado retreiving failed: %s", quado.ID)
	}
	log.WithFields(log.Fields{
		"id":     resQuado.ID,
		"ListId": resQuado.ListID,
		"title":  resQuado.Title,
		"desc":   resQuado.Description,
		"date":   resQuado.Date,
	}).Info("Retreiving quado")

	size := bucketSize(QUADO_BUCKET, storage)
	if size != 1 {
		t.Errorf("Quado bucket size error: %v", size)
	}

	destroyTestStorage(storage, t)
}

func TestUpdateQuado(t *testing.T) {
	storage := initQuado(t)
	quado := newQuado(newList(newBoard().ID, "").ID, "test", "desc", time.Now())

	if err := putQuado(quado, storage); err != nil {
		t.Errorf("Quado insertion in bucket failed: %s", quado.ID)
	}
	log.WithField("id", quado.ID).Info("New quado")

	resQuado, err := getQuado(quado.ID, storage)
	if err != nil {
		t.Errorf("Quado retreiving failed: %s", quado.ID)
	}
	if resQuado.ID != quado.ID {
		t.Errorf("Quado retreiving failed: %s", quado.ID)
	}
	log.WithFields(log.Fields{
		"id":     resQuado.ID,
		"ListId": resQuado.ListID,
		"title":  resQuado.Title,
		"desc":   resQuado.Description,
		"date":   resQuado.Date,
	}).Info("Retreiving quado")

	size := bucketSize(QUADO_BUCKET, storage)
	if size != 1 {
		t.Errorf("Quado bucket size error: %v", size)
	}

	quado.Title = "test_modify"

	if err := putQuado(quado, storage); err != nil {
		t.Errorf("Quado insertion in bucket failed: %s", quado.ID)
	}
	log.WithField("id", quado.ID).Info("Update quado")

	resQuado, err = getQuado(quado.ID, storage)
	if err != nil {
		t.Errorf("Quado retreiving failed: %s", quado.ID)
	}
	if resQuado.ID != quado.ID {
		t.Errorf("Quado retreiving failed: %s", quado.ID)
	}
	log.WithFields(log.Fields{
		"id":     resQuado.ID,
		"ListId": resQuado.ListID,
		"title":  resQuado.Title,
		"desc":   resQuado.Description,
		"date":   resQuado.Date,
	}).Info("Retreiving quado")

	size = bucketSize(QUADO_BUCKET, storage)
	if size != 1 {
		t.Errorf("Quado bucket size error: %v", size)
	}

	destroyTestStorage(storage, t)
}

func TestDeleteQuado(t *testing.T) {
	storage := initQuado(t)
	quado := newQuado(newList(newBoard().ID, "").ID, "test", "desc", time.Now())

	if err := putQuado(quado, storage); err != nil {
		t.Errorf("Quado insertion in bucket failed: %s", quado.ID)
	}
	log.WithField("id", quado.ID).Info("New quado")

	size := bucketSize(QUADO_BUCKET, storage)
	if size != 1 {
		t.Errorf("Quado bucket size error: %v", size)
	}

	if err := deleteQuado(quado, storage); err != nil {
		t.Errorf("Quado deletion failed: %s", quado.ID)
	}
	log.WithField("id", quado.ID).Info("Deleting quado")

	size = bucketSize(QUADO_BUCKET, storage)
	if size != 0 {
		t.Errorf("Quado bucket size error: %v", size)
	}

	destroyTestStorage(storage, t)
}
