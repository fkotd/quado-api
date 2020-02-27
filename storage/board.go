package storage

import (
	"encoding/json"

	uuid "github.com/satori/go.uuid"
	bolt "go.etcd.io/bbolt"
)

type Board struct {
	ID string
}

func newBoard() *Board {
	return &Board{uuid.NewV4().String()}
}

func putBoard(board *Board, s *Storage) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BOARD_BUCKET))

		json, err := json.Marshal(board)
		if err != nil {
			return err
		}

		return b.Put([]byte(board.ID), json)
	})
}

func getBoard(id string, s *Storage) (board *Board, err error) {
	err = s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BOARD_BUCKET))

		boardByte := b.Get([]byte(id))

		return json.Unmarshal(boardByte, &board)
	})
	return
}

func deleteBoard(board *Board, s *Storage) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BOARD_BUCKET))

		return b.Delete([]byte(board.ID))
	})
}
