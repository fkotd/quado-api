package handler

import "github.com/finalKickOfTheDeath/quado-api/storage"

type Handler struct {
	storage *storage.Storage
}

func NewHandler(storage *storage.Storage) *Handler {
	return &Handler{storage}
}
