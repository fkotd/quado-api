package storage

import (
	"encoding/json"

	uuid "github.com/satori/go.uuid"
	bolt "go.etcd.io/bbolt"
)

type List struct {
	ID      string `json:"id"`
	BoardID string `json:"board-id" binding:"required"`
	Title   string `json:"title"`
}

func (storage *Storage) NewList(boardID string, title string) *List {
	return &List{uuid.NewV4().String(), boardID, title}
}

func (storage *Storage) PutList(list *List) error {
	return storage.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(LIST_BUCKET))

		json, err := json.Marshal(list)
		if err != nil {
			return err
		}

		return bucket.Put([]byte(list.ID), json)
	})
}

func (storage *Storage) GetList(id string) (list *List, err error) {
	err = storage.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(LIST_BUCKET))

		listByte := bucket.Get([]byte(id))

		return json.Unmarshal(listByte, &list)
	})
	return
}

func (storage *Storage) DeleteList(list *List) error {
	return storage.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(LIST_BUCKET))

		return bucket.Delete([]byte(list.ID))
	})
}

func (storage *Storage) deleteList(list *List, tx *bolt.Tx) error {
	quadoBucket := tx.Bucket([]byte(QUADO_BUCKET))

	err := quadoBucket.ForEach(func(key, value []byte) error {
		var quado Quado
		if err := json.Unmarshal(value, &quado); err != nil {
			return err
		}
		if quado.ListID == list.ID {
			if err := storage.deleteQuado(&quado, tx); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	bucketList := tx.Bucket([]byte(LIST_BUCKET))
	bucketList.Delete([]byte(list.ID))

	return nil
}
