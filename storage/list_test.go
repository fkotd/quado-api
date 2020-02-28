package storage

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestPutList(t *testing.T) {
	storage := createTestStorage(t)
	list := storage.NewList(storage.NewBoard().ID, "test")

	if err := storage.PutList(list); err != nil {
		t.Errorf("List insertion in bucket failed: %s", list.ID)
	}
	log.WithFields(log.Fields{
		"id":      list.ID,
		"boardId": list.BoardID,
		"title":   list.Title,
	}).Trace("New list")

	size := storage.bucketSize(LIST_BUCKET)
	if size != 1 {
		t.Errorf("List bucket size error: %v", size)
	}

	destroyTestStorage(storage, t)
}

func TestGetList(t *testing.T) {
	storage := createTestStorage(t)
	list := storage.NewList(storage.NewBoard().ID, "test")

	if err := storage.PutList(list); err != nil {
		t.Errorf("List insertion in bucket failed: %s", list.ID)
	}
	log.WithField("id", list.ID).Trace("New list")

	resList, err := storage.GetList(list.ID)
	if err != nil {
		t.Errorf("List retreiving failed: %s", list.ID)
	}
	if resList.ID != list.ID {
		t.Errorf("List retreiving failed: %s", list.ID)
	}
	log.WithFields(log.Fields{
		"id":      resList.ID,
		"boardId": resList.BoardID,
		"title":   resList.Title,
	}).Trace("Retreiving list")

	size := storage.bucketSize(LIST_BUCKET)
	if size != 1 {
		t.Errorf("List bucket size error: %v", size)
	}

	destroyTestStorage(storage, t)
}

func TestUpdateList(t *testing.T) {
	storage := createTestStorage(t)
	list := storage.NewList(storage.NewBoard().ID, "test")

	if err := storage.PutList(list); err != nil {
		t.Errorf("List insertion in bucket failed: %s", list.ID)
	}
	log.WithField("id", list.ID).Trace("New list")

	resList, err := storage.GetList(list.ID)
	if err != nil {
		t.Errorf("List retreiving failed: %s", list.ID)
	}
	if resList.ID != list.ID {
		t.Errorf("List retreiving failed: %s", list.ID)
	}
	log.WithFields(log.Fields{
		"id":      resList.ID,
		"boardId": resList.BoardID,
		"title":   resList.Title,
	}).Trace("Retreiving list")

	size := storage.bucketSize(LIST_BUCKET)
	if size != 1 {
		t.Errorf("List bucket size error: %v", size)
	}

	list.Title = "test_modify"

	if err := storage.PutList(list); err != nil {
		t.Errorf("List insertion in bucket failed: %s", list.ID)
	}
	log.WithField("id", list.ID).Trace("Update list")

	resList, err = storage.GetList(list.ID)
	if err != nil {
		t.Errorf("List retreiving failed: %s", list.ID)
	}
	if resList.ID != list.ID {
		t.Errorf("List retreiving failed: %s", list.ID)
	}
	log.WithFields(log.Fields{
		"id":      resList.ID,
		"boardId": resList.BoardID,
		"title":   resList.Title,
	}).Trace("Retreiving list")

	size = storage.bucketSize(LIST_BUCKET)
	if size != 1 {
		t.Errorf("List bucket size error: %v", size)
	}

	destroyTestStorage(storage, t)
}

func TestDeleteList(t *testing.T) {
	storage := createTestStorage(t)
	list := storage.NewList(storage.NewBoard().ID, "test")

	if err := storage.PutList(list); err != nil {
		t.Errorf("List insertion in bucket failed: %s", list.ID)
	}
	log.WithField("id", list.ID).Trace("New list")

	size := storage.bucketSize(LIST_BUCKET)
	if size != 1 {
		t.Errorf("List bucket size error: %v", size)
	}

	if err := storage.DeleteList(list); err != nil {
		t.Errorf("List deletion failed: %s", list.ID)
	}
	log.WithField("id", list.ID).Trace("Deleting list")

	size = storage.bucketSize(LIST_BUCKET)
	if size != 0 {
		t.Errorf("List bucket size error: %v", size)
	}

	destroyTestStorage(storage, t)
}
