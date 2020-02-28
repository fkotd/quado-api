package storage

import (
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
)

func TestPutQuado(t *testing.T) {
	storage := createTestStorage(t)
	quado := storage.NewQuado(storage.NewList(storage.NewBoard().ID, "").ID, "test", "desc", time.Now())

	if err := storage.PutQuado(quado); err != nil {
		t.Errorf("Quado insertion in bucket failed: %s", quado.ID)
	}
	log.WithFields(log.Fields{
		"id":     quado.ID,
		"ListId": quado.ListID,
		"title":  quado.Title,
		"desc":   quado.Description,
		"date":   quado.Date,
	}).Trace("New quado")

	size := storage.bucketSize(QUADO_BUCKET)
	if size != 1 {
		t.Errorf("Quado bucket size error: %v", size)
	}

	destroyTestStorage(storage, t)
}

func TestGetQuado(t *testing.T) {
	storage := createTestStorage(t)
	quado := storage.NewQuado(storage.NewList(storage.NewBoard().ID, "").ID, "test", "desc", time.Now())

	if err := storage.PutQuado(quado); err != nil {
		t.Errorf("Quado insertion in bucket failed: %s", quado.ID)
	}
	log.WithField("id", quado.ID).Trace("New quado")

	resQuado, err := storage.GetQuado(quado.ID)
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
	}).Trace("Retreiving quado")

	size := storage.bucketSize(QUADO_BUCKET)
	if size != 1 {
		t.Errorf("Quado bucket size error: %v", size)
	}

	destroyTestStorage(storage, t)
}

func TestUpdateQuado(t *testing.T) {
	storage := createTestStorage(t)
	quado := storage.NewQuado(storage.NewList(storage.NewBoard().ID, "").ID, "test", "desc", time.Now())

	if err := storage.PutQuado(quado); err != nil {
		t.Errorf("Quado insertion in bucket failed: %s", quado.ID)
	}
	log.WithField("id", quado.ID).Trace("New quado")

	resQuado, err := storage.GetQuado(quado.ID)
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
	}).Trace("Retreiving quado")

	size := storage.bucketSize(QUADO_BUCKET)
	if size != 1 {
		t.Errorf("Quado bucket size error: %v", size)
	}

	quado.Title = "test_modify"

	if err := storage.PutQuado(quado); err != nil {
		t.Errorf("Quado insertion in bucket failed: %s", quado.ID)
	}
	log.WithField("id", quado.ID).Trace("Update quado")

	resQuado, err = storage.GetQuado(quado.ID)
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
	}).Trace("Retreiving quado")

	size = storage.bucketSize(QUADO_BUCKET)
	if size != 1 {
		t.Errorf("Quado bucket size error: %v", size)
	}

	destroyTestStorage(storage, t)
}

func TestDeleteQuado(t *testing.T) {
	storage := createTestStorage(t)
	quado := storage.NewQuado(storage.NewList(storage.NewBoard().ID, "").ID, "test", "desc", time.Now())

	if err := storage.PutQuado(quado); err != nil {
		t.Errorf("Quado insertion in bucket failed: %s", quado.ID)
	}
	log.WithField("id", quado.ID).Trace("New quado")

	size := storage.bucketSize(QUADO_BUCKET)
	if size != 1 {
		t.Errorf("Quado bucket size error: %v", size)
	}

	if err := storage.DeleteQuado(quado); err != nil {
		t.Errorf("Quado deletion failed: %s", quado.ID)
	}
	log.WithField("id", quado.ID).Trace("Deleting quado")

	size = storage.bucketSize(QUADO_BUCKET)
	if size != 0 {
		t.Errorf("Quado bucket size error: %v", size)
	}

	destroyTestStorage(storage, t)
}
