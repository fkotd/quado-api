package storage

import (
	"encoding/json"
	"time"

	uuid "github.com/satori/go.uuid"
	bolt "go.etcd.io/bbolt"
)

type Quado struct {
	Id          string    `json:"id"`
	ListId      string    `json:"listId" binding:"required"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
}

type QuadoResult struct {
	Id          string    `json:"id"`
	ListId      string    `json:"listId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
}

func (storage *Storage) NewQuado(listId string, title string, description string, deadline time.Time) *Quado {
	return &Quado{uuid.NewV4().String(), listId, title, description, deadline}
}

func (storage *Storage) NewQuadoResult(quado *Quado) *QuadoResult {
	return &QuadoResult{quado.Id, quado.ListId, quado.Title, quado.Description, quado.Deadline}
}

func (storage *Storage) PutQuado(quado *Quado) error {
	return storage.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(QUADO_BUCKET))

		json, err := json.Marshal(quado)
		if err != nil {
			return err
		}

		return bucket.Put([]byte(quado.Id), json)
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

func (storage *Storage) getQuados(listId string, tx *bolt.Tx) (quados []QuadoResult, err error) {
	quadoBucket := tx.Bucket([]byte(QUADO_BUCKET))

	err = quadoBucket.ForEach(func(key, value []byte) error {
		var quado Quado

		if err := json.Unmarshal(value, &quado); err != nil {
			return err
		}

		if quado.ListId == listId {
			quadoResult := QuadoResult{quado.Id, listId, quado.Title, quado.Description, quado.Deadline}
			quados = append(quados, quadoResult)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return quados, nil
}

func (storage *Storage) DeleteQuado(quado *Quado) error {
	return storage.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(QUADO_BUCKET))

		return bucket.Delete([]byte(quado.Id))
	})
}

func (storage *Storage) deleteQuado(quado *Quado, tx *bolt.Tx) error {
	bucket := tx.Bucket([]byte(QUADO_BUCKET))

	return bucket.Delete([]byte(quado.Id))
}
