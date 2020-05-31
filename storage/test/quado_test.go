package storage

import (
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
)

func TestPutQuado(t *testing.T) {
	storage := createTestStorage(t)
	quado := storage.NewQuado(storage.NewList(storage.NewBoard().Id, "").Id, "test", "desc", time.Now())

	if err := storage.PutQuado(quado); err != nil {
		t.Errorf("Quado insertion in bucket failed: %s", quado.Id)
	}
	log.WithFields(log.Fields{
		"id":       quado.Id,
		"ListId":   quado.ListId,
		"title":    quado.Title,
		"desc":     quado.Description,
		"deadline": quado.Deadline,
	}).Trace("New quado")

	size := storage.bucketSize(QUADO_BUCKET)
	if size != 1 {
		t.Errorf("Quado bucket size error: %v", size)
	}

	destroyTestStorage(storage, t)
}

func TestGetQuado(t *testing.T) {
	storage := createTestStorage(t)
	quado := storage.NewQuado(storage.NewList(storage.NewBoard().Id, "").Id, "test", "desc", time.Now())

	if err := storage.PutQuado(quado); err != nil {
		t.Errorf("Quado insertion in bucket failed: %s", quado.Id)
	}
	log.WithField("id", quado.Id).Trace("New quado")

	resQuado, err := storage.GetQuado(quado.Id)
	if err != nil {
		t.Errorf("Quado retreiving failed: %s", quado.Id)
	}
	if resQuado.Id != quado.Id {
		t.Errorf("Quado retreiving failed: %s", quado.Id)
	}
	log.WithFields(log.Fields{
		"id":       resQuado.Id,
		"ListId":   resQuado.ListId,
		"title":    resQuado.Title,
		"desc":     resQuado.Description,
		"deadline": resQuado.Deadline,
	}).Trace("Retreiving quado")

	size := storage.bucketSize(QUADO_BUCKET)
	if size != 1 {
		t.Errorf("Quado bucket size error: %v", size)
	}

	destroyTestStorage(storage, t)
}

func TestUpdateQuado(t *testing.T) {
	storage := createTestStorage(t)
	quado := storage.NewQuado(storage.NewList(storage.NewBoard().Id, "").Id, "test", "desc", time.Now())

	if err := storage.PutQuado(quado); err != nil {
		t.Errorf("Quado insertion in bucket failed: %s", quado.Id)
	}
	log.WithField("id", quado.Id).Trace("New quado")

	resQuado, err := storage.GetQuado(quado.Id)
	if err != nil {
		t.Errorf("Quado retreiving failed: %s", quado.Id)
	}
	if resQuado.Id != quado.Id {
		t.Errorf("Quado retreiving failed: %s", quado.Id)
	}
	log.WithFields(log.Fields{
		"id":       resQuado.Id,
		"ListId":   resQuado.ListId,
		"title":    resQuado.Title,
		"desc":     resQuado.Description,
		"deadline": resQuado.Deadline,
	}).Trace("Retreiving quado")

	size := storage.bucketSize(QUADO_BUCKET)
	if size != 1 {
		t.Errorf("Quado bucket size error: %v", size)
	}

	quado.Title = "test_modify"

	if err := storage.PutQuado(quado); err != nil {
		t.Errorf("Quado insertion in bucket failed: %s", quado.Id)
	}
	log.WithField("id", quado.Id).Trace("Update quado")

	resQuado, err = storage.GetQuado(quado.Id)
	if err != nil {
		t.Errorf("Quado retreiving failed: %s", quado.Id)
	}
	if resQuado.Id != quado.Id {
		t.Errorf("Quado retreiving failed: %s", quado.Id)
	}
	log.WithFields(log.Fields{
		"id":       resQuado.Id,
		"ListId":   resQuado.ListId,
		"title":    resQuado.Title,
		"desc":     resQuado.Description,
		"deadline": resQuado.Deadline,
	}).Trace("Retreiving quado")

	size = storage.bucketSize(QUADO_BUCKET)
	if size != 1 {
		t.Errorf("Quado bucket size error: %v", size)
	}

	destroyTestStorage(storage, t)
}

func TestDeleteQuado(t *testing.T) {
	storage := createTestStorage(t)
	quado := storage.NewQuado(storage.NewList(storage.NewBoard().Id, "").Id, "test", "desc", time.Now())

	if err := storage.PutQuado(quado); err != nil {
		t.Errorf("Quado insertion in bucket failed: %s", quado.Id)
	}
	log.WithField("id", quado.Id).Trace("New quado")

	size := storage.bucketSize(QUADO_BUCKET)
	if size != 1 {
		t.Errorf("Quado bucket size error: %v", size)
	}

	if err := storage.DeleteQuado(quado); err != nil {
		t.Errorf("Quado deletion failed: %s", quado.Id)
	}
	log.WithField("id", quado.Id).Trace("Deleting quado")

	size = storage.bucketSize(QUADO_BUCKET)
	if size != 0 {
		t.Errorf("Quado bucket size error: %v", size)
	}

	destroyTestStorage(storage, t)
}
