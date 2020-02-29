package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *Handler) CreateList(context *gin.Context) {
	boardID := context.PostForm("idBoard")
	title := context.PostForm("title")

	list := handler.storage.NewList(boardID, title)

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
