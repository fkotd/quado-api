package storage

import (
	"encoding/json"
	"time"

	uuid "github.com/satori/go.uuid"
	bolt "go.etcd.io/bbolt"
)

type Quado struct {
	ID          string
	ListID      string
	Title       string
	Description string
	Date        time.Time
}

func newQuado(listID string, title string, description string, date time.Time) *Quado {
	return &Quado{uuid.NewV4().String(), listID, title, description, date}
}

func putQuado(quado *Quado, s *Storage) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(QUADO_BUCKET))

		json, err := json.Marshal(quado)
		if err != nil {
			return err
		}

		return b.Put([]byte(quado.ID), json)
	})
}

func getQuado(id string, s *Storage) (quado *Quado, err error) {
	err = s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(QUADO_BUCKET))

		quadoByte := b.Get([]byte(id))

		return json.Unmarshal(quadoByte, &quado)
	})
	return
}

func deleteQuado(quado *Quado, s *Storage) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(QUADO_BUCKET))

		return b.Delete([]byte(quado.ID))
	})
}
