package storage

import (
	"encoding/json"

	uuid "github.com/satori/go.uuid"
	bolt "go.etcd.io/bbolt"
)

type Board struct {
	Id string `json:"id"`
}

func (storage *Storage) NewBoard() *Board {
	return &Board{uuid.NewV4().String()}
}

func (storage *Storage) PutBoard(board *Board) error {
	return storage.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BOARD_BUCKET))

		json, err := json.Marshal(board)
		if err != nil {
			return err
		}

		return bucket.Put([]byte(board.Id), json)
	})
}

func (storage *Storage) GetBoard(id string) (board *Board, err error) {
	err = storage.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BOARD_BUCKET))

		boardByte := bucket.Get([]byte(id))

		return json.Unmarshal(boardByte, &board)
	})
	return
}

func (storage *Storage) DeleteBoard(board *Board) error {
	tx, err := storage.db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	listBucket := tx.Bucket([]byte(LIST_BUCKET))

	err = listBucket.ForEach(func(key, value []byte) error {
		var list List

		if err := json.Unmarshal(value, &list); err != nil {
			return err
		}

		if list.BoardId == board.Id {
			if err := storage.deleteList(&list, tx); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	boardBucket := tx.Bucket([]byte(BOARD_BUCKET))
	boardBucket.Delete([]byte(board.Id))

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
