package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (handler *Handler) CreateList(c *gin.Context) {
	boardID := c.PostForm("board-id")
	title := c.PostForm("title")

	list := handler.storage.NewList(boardID, title)
	log.WithFields(log.Fields{
		"id":       list.ID,
		"board-id": list.BoardID,
		"title":    list.Title,
	}).Warn("List")

	handler.storage.PutList(list)

	c.JSON(http.StatusOK, list)
}
