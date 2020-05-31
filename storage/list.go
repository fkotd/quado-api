package storage

import (
	"encoding/json"

	uuid "github.com/satori/go.uuid"
	bolt "go.etcd.io/bbolt"
)

type List struct {
	Id      string `json:"id"`
	BoardId string `json:"boardId" binding:"required"`
	Title   string `json:"title"`
}

type ListResult struct {
	Id      string        `json:"id"`
	BoardId string        `json:"boardId"`
	Title   string        `json:"title"`
	Quados  []QuadoResult `json:"quados"`
}

func (storage *Storage) NewList(boardId string, title string) *List {
	return &List{uuid.NewV4().String(), boardId, title}
}

func (storage *Storage) NewListResult(list *List) *ListResult {
	var quados []QuadoResult
	return &ListResult{list.Id, list.BoardId, list.Title, quados}
}

func (storage *Storage) PutList(list *List) error {
	return storage.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(LIST_BUCKET))

		json, err := json.Marshal(list)
		if err != nil {
			return err
		}

		return bucket.Put([]byte(list.Id), json)
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

func (storage *Storage) GetLists(boardId string) (lists []ListResult, err error) {
	tx, err := storage.db.Begin(true)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	listBucket := tx.Bucket([]byte(LIST_BUCKET))

	err = listBucket.ForEach(func(key, value []byte) error {
		var list List

		if err := json.Unmarshal(value, &list); err != nil {
			return err
		}

		if list.BoardId == boardId {
			quados, err := storage.getQuados(list.Id, tx)
			if err != nil {
				return err
			}
			listResult := ListResult{list.Id, boardId, list.Title, quados}
			lists = append(lists, listResult)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return lists, nil
}

func (storage *Storage) DeleteList(list *List) error {
	return storage.db.Update(func(tx *bolt.Tx) error {
		if err := storage.deleteList(list, tx); err != nil {
			return err
		}
		return nil
	})
}

func (storage *Storage) deleteList(list *List, tx *bolt.Tx) error {
	quadoBucket := tx.Bucket([]byte(QUADO_BUCKET))

	err := quadoBucket.ForEach(func(key, value []byte) error {
		var quado Quado

		if err := json.Unmarshal(value, &quado); err != nil {
			return err
		}

		if quado.ListId == list.Id {
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
	bucketList.Delete([]byte(list.Id))

	return nil
}
