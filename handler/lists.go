package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type TitleJSON struct {
	Title string `json:"title"`
}

func (handler *Handler) CreateList(context *gin.Context) {
	boardID := context.Param("id")

	var titleJSON TitleJSON
	if err := context.ShouldBindJSON(&titleJSON); err != nil {
		context.Status(http.StatusBadRequest)
		return
	}

	log.SetLevel(log.DebugLevel)
	log.WithFields(log.Fields{
		"idBoard": boardID,
		"title":   titleJSON.Title,
	}).Debug("A walrus appears")

	if _, err := handler.storage.GetBoard(boardID); err != nil {
		context.Status(http.StatusBadRequest)
		return
	}

	list := handler.storage.NewList(boardID, titleJSON.Title)

	handler.storage.PutList(list)

	context.JSON(http.StatusOK, list)
}

func (handler *Handler) RemoveList(context *gin.Context) {
	listID := context.Param("id")

	list, err := handler.storage.GetList(listID)
	if err != nil {
		context.Status(http.StatusInternalServerError)
		return
	}

	err = handler.storage.DeleteList(list)
	if err != nil {
		context.Status(http.StatusInternalServerError)
		return
	}

	context.Status(http.StatusOK)
}
