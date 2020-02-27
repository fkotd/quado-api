package storage

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

func initBoard(t *testing.T) *Storage {
	storage := newStorage(createTestConfig(t))
	var err error

	if err = storage.open(); err != nil {
		t.Errorf("Opening failed")
	}

	if err = storage.createBucket(BOARD_BUCKET); err != nil {
		t.Errorf("Bucket creation failed: %s", BOARD_BUCKET)
	}

	return storage
}

func TestPutBoard(t *testing.T) {
	storage := initBoard(t)
	board := newBoard()

	if err := putBoard(board, storage); err != nil {
		t.Errorf("Board insertion in bucket failed: %s", board.ID)
	}
	log.WithField("id", board.ID).Info("New board")

	size := bucketSize(BOARD_BUCKET, storage)
	if size != 1 {
		t.Errorf("Board bucket size error: %v", size)
	}

	destroyTestStorage(storage, t)
}

func TestGetBoard(t *testing.T) {
	storage := initBoard(t)
	board := newBoard()

	if err := putBoard(board, storage); err != nil {
		t.Errorf("Board insertion in bucket failed: %s", board.ID)
	}
	log.WithField("id", board.ID).Info("New board")

	resBoard, err := getBoard(board.ID, storage)
	if err != nil {
		t.Errorf("Board retreiving failed: %s", board.ID)
	}
	if resBoard.ID != board.ID {
		t.Errorf("Board retreiving failed: %s", board.ID)
	}
	log.WithField("id", resBoard.ID).Info("Retreiving board")

	size := bucketSize(BOARD_BUCKET, storage)
	if size != 1 {
		t.Errorf("Board bucket size error: %v", size)
	}

	destroyTestStorage(storage, t)
}

func TestDeleteBoard(t *testing.T) {
	storage := initBoard(t)
	board := newBoard()

	if err := putBoard(board, storage); err != nil {
		t.Errorf("Board insertion in bucket failed: %s", board.ID)
	}
	log.WithField("id", board.ID).Info("New board")

	size := bucketSize(BOARD_BUCKET, storage)
	if size != 1 {
		t.Errorf("Board bucket size error: %v", size)
	}

	if err := deleteBoard(board, storage); err != nil {
		t.Errorf("Board deletion failed: %s", board.ID)
	}
	log.WithField("id", board.ID).Info("Deleting board")

	size = bucketSize(BOARD_BUCKET, storage)
	if size != 0 {
		t.Errorf("Board bucket size error: %v", size)
	}

	destroyTestStorage(storage, t)
}
