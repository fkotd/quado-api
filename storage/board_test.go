package storage

import (
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
)

func TestPutBoard(t *testing.T) {
	storage := createTestStorage(t)
	board := storage.NewBoard()

	if err := storage.PutBoard(board); err != nil {
		t.Errorf("Board insertion in bucket failed: %s", board.ID)
	}
	log.WithField("id", board.ID).Trace("New board")

	size := storage.bucketSize(BOARD_BUCKET)
	if size != 1 {
		t.Errorf("Board bucket size error: %v", size)
	}

	destroyTestStorage(storage, t)
}

func TestGetBoard(t *testing.T) {
	storage := createTestStorage(t)
	board := storage.NewBoard()

	if err := storage.PutBoard(board); err != nil {
		t.Errorf("Board insertion in bucket failed: %s", board.ID)
	}
	log.WithField("id", board.ID).Trace("New board")

	resBoard, err := storage.GetBoard(board.ID)
	if err != nil {
		t.Errorf("Board retreiving failed: %s", board.ID)
	}
	if resBoard.ID != board.ID {
		t.Errorf("Board retreiving failed: %s", board.ID)
	}
	log.WithField("id", resBoard.ID).Trace("Retreiving board")

	size := storage.bucketSize(BOARD_BUCKET)
	if size != 1 {
		t.Errorf("Board bucket size error: %v", size)
	}

	destroyTestStorage(storage, t)
}

func TestDeleteBoard(t *testing.T) {
	storage := createTestStorage(t)
	board := storage.NewBoard()

	if err := storage.PutBoard(board); err != nil {
		t.Errorf("Board insertion in bucket failed: %s", board.ID)
	}
	log.WithField("id", board.ID).Trace("New board")

	size := storage.bucketSize(BOARD_BUCKET)
	if size != 1 {
		t.Errorf("Board bucket size error: %v", size)
	}

	if err := storage.DeleteBoard(board); err != nil {
		t.Errorf("Board deletion failed: %s", board.ID)
	}
	log.WithField("id", board.ID).Trace("Deleting board")

	size = storage.bucketSize(BOARD_BUCKET)
	if size != 0 {
		t.Errorf("Board bucket size error: %v", size)
	}

	destroyTestStorage(storage, t)
}

func TestDeleteBoardCascade(t *testing.T) {
	storage := createTestStorage(t)

	board := storage.NewBoard()
	list1 := storage.NewList(board.ID, "list1")
	list2 := storage.NewList(board.ID, "list2")
	quado11 := storage.NewQuado(list1.ID, "quado11", "desc11", time.Now())
	quado12 := storage.NewQuado(list1.ID, "quado12", "desc12", time.Now())

	if err := storage.PutBoard(board); err != nil {
		t.Errorf("Board insertion in bucket failed: %s", board.ID)
	}
	log.WithField("id", board.ID).Trace("New board")

	size := storage.bucketSize(BOARD_BUCKET)
	if size != 1 {
		t.Errorf("Board bucket size error: %v", size)
	}

	if err := storage.PutList(list1); err != nil {
		t.Errorf("List insertion in bucket failed: %s", list1.ID)
	}
	log.WithField("id", list1.ID).Trace("New list")

	size = storage.bucketSize(LIST_BUCKET)
	if size != 1 {
		t.Errorf("List bucket size error: %v", size)
	}

	if err := storage.PutList(list2); err != nil {
		t.Errorf("List insertion in bucket failed: %s", list2.ID)
	}
	log.WithField("id", list2.ID).Trace("New list")

	size = storage.bucketSize(LIST_BUCKET)
	if size != 2 {
		t.Errorf("List bucket size error: %v", size)
	}

	if err := storage.PutQuado(quado11); err != nil {
		t.Errorf("Quado insertion in bucket failed: %s", quado11.ID)
	}
	log.WithField("id", quado11.ID).Trace("New quado")

	size = storage.bucketSize(QUADO_BUCKET)
	if size != 1 {
		t.Errorf("Quado bucket size error: %v", size)
	}

	if err := storage.PutQuado(quado12); err != nil {
		t.Errorf("Quado insertion in bucket failed: %s", quado12.ID)
	}
	log.WithField("id", quado12.ID).Trace("New quado")

	size = storage.bucketSize(QUADO_BUCKET)
	if size != 2 {
		t.Errorf("Quado bucket size error: %v", size)
	}

	if err := storage.DeleteBoard(board); err != nil {
		t.Errorf("Board deletion failed: %s", board.ID)
	}
	log.WithField("id", board.ID).Trace("Deleting board")

	size = storage.bucketSize(BOARD_BUCKET)
	if size != 0 {
		t.Errorf("Board bucket size error: %v", size)
	}

	size = storage.bucketSize(LIST_BUCKET)
	if size != 0 {
		t.Errorf("List bucket size error: %v", size)
	}

	size = storage.bucketSize(QUADO_BUCKET)
	if size != 0 {
		t.Errorf("Quado bucket size error: %v", size)
	}

	destroyTestStorage(storage, t)
}
