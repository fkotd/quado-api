package storage

import (
	"encoding/json"
	"time"

	uuid "github.com/satori/go.uuid"
	bolt "go.etcd.io/bbolt"
)

type Quado struct {
	ID          string    `json:"id"`
	ListID      string    `json:"idList" binding:"required"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

func (storage *Storage) NewQuado(listID string, title string, description string, date time.Time) *Quado {
	return &Quado{uuid.NewV4().String(), listID, title, description, date}
}

func (storage *Storage) PutQuado(quado *Quado) error {
	return storage.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(QUADO_BUCKET))

		json, err := json.Marshal(quado)
		if err != nil {
			return err
		}

		return bucket.Put([]byte(quado.ID), json)
	})
}

func (storage *Storage) GetQuado(id string) (quado *Quado, err error) {
	err = storage.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(QUADO_BUCKET))

		quadoByte := bucket.Get([]byte(id))

		return json.Unmarshal(quadoByte, &quado)
	})
	return
}

func (storage *Storage) DeleteQuado(quado *Quado) error {
	return storage.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(QUADO_BUCKET))

		return bucket.Delete([]byte(quado.ID))
	})
}

func (storage *Storage) deleteQuado(quado *Quado, tx *bolt.Tx) error {
	bucket := tx.Bucket([]byte(QUADO_BUCKET))

	return bucket.Delete([]byte(quado.ID))
}
