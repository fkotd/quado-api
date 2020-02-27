package storage

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

func initList(t *testing.T) *Storage {
	storage := newStorage(createTestConfig(t))
	var err error

	if err = storage.open(); err != nil {
		t.Errorf("Opening failed")
	}

	if err = storage.createBucket(LIST_BUCKET); err != nil {
		t.Errorf("Bucket creation failed: %s", LIST_BUCKET)
	}

	return storage
}

func TestPutList(t *testing.T) {
	storage := initList(t)
	list := newList(newBoard().ID, "test")

	if err := putList(list, storage); err != nil {
		t.Errorf("List insertion in bucket failed: %s", list.ID)
	}
	log.WithFields(log.Fields{
		"id":      list.ID,
		"boardId": list.BoardID,
		"title":   list.Title,
	}).Info("New list")

	size := bucketSize(LIST_BUCKET, storage)
	if size != 1 {
		t.Errorf("List bucket size error: %v", size)
	}

	destroyTestStorage(storage, t)
}

func TestGetList(t *testing.T) {
	storage := initList(t)
	list := newList(newBoard().ID, "test")

	if err := putList(list, storage); err != nil {
		t.Errorf("List insertion in bucket failed: %s", list.ID)
	}
	log.WithField("id", list.ID).Info("New list")

	resList, err := getList(list.ID, storage)
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
	}).Info("Retreiving list")

	size := bucketSize(LIST_BUCKET, storage)
	if size != 1 {
		t.Errorf("List bucket size error: %v", size)
	}

	destroyTestStorage(storage, t)
}

func TestUpdateList(t *testing.T) {
	storage := initList(t)
	list := newList(newBoard().ID, "test")

	if err := putList(list, storage); err != nil {
		t.Errorf("List insertion in bucket failed: %s", list.ID)
	}
	log.WithField("id", list.ID).Info("New list")

	resList, err := getList(list.ID, storage)
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
	}).Info("Retreiving list")

	size := bucketSize(LIST_BUCKET, storage)
	if size != 1 {
		t.Errorf("List bucket size error: %v", size)
	}

	list.Title = "test_modify"

	if err := putList(list, storage); err != nil {
		t.Errorf("List insertion in bucket failed: %s", list.ID)
	}
	log.WithField("id", list.ID).Info("Update list")

	resList, err = getList(list.ID, storage)
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
	}).Info("Retreiving list")

	size = bucketSize(LIST_BUCKET, storage)
	if size != 1 {
		t.Errorf("List bucket size error: %v", size)
	}

	destroyTestStorage(storage, t)
}

func TestDeleteList(t *testing.T) {
	storage := initList(t)
	list := newList(newBoard().ID, "test")

	if err := putList(list, storage); err != nil {
		t.Errorf("List insertion in bucket failed: %s", list.ID)
	}
	log.WithField("id", list.ID).Info("New list")

	size := bucketSize(LIST_BUCKET, storage)
	if size != 1 {
		t.Errorf("List bucket size error: %v", size)
	}

	if err := deleteList(list, storage); err != nil {
		t.Errorf("List deletion failed: %s", list.ID)
	}
	log.WithField("id", list.ID).Info("Deleting list")

	size = bucketSize(LIST_BUCKET, storage)
	if size != 0 {
		t.Errorf("List bucket size error: %v", size)
	}

	destroyTestStorage(storage, t)
}
