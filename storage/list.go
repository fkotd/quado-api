package storage

import (
	"encoding/json"

	uuid "github.com/satori/go.uuid"
	bolt "go.etcd.io/bbolt"
)

type List struct {
	ID      string
	BoardID string
	Title   string
}

func newList(boardID string, title string) *List {
	return &List{uuid.NewV4().String(), boardID, title}
}

func putList(list *List, s *Storage) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(LIST_BUCKET))

		json, err := json.Marshal(list)
		if err != nil {
			return err
		}

		return b.Put([]byte(list.ID), json)
	})
}

func getList(id string, s *Storage) (list *List, err error) {
	err = s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(LIST_BUCKET))

		listByte := b.Get([]byte(id))

		return json.Unmarshal(listByte, &list)
	})
	return
}

func deleteList(list *List, s *Storage) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(LIST_BUCKET))

		return b.Delete([]byte(list.ID))
	})
}
