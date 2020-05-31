package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TitleJSON struct {
	Title string `json:"title"`
}

func (handler *Handler) CreateList(context *gin.Context) {
	boardId := context.Param("id")

	var titleJSON TitleJSON
	if err := context.ShouldBindJSON(&titleJSON); err != nil {
		context.Status(http.StatusBadRequest)
		return
	}

	if _, err := handler.storage.GetBoard(boardId); err != nil {
		context.Status(http.StatusBadRequest)
		return
	}

	list := handler.storage.NewList(boardId, titleJSON.Title)
	handler.storage.PutList(list)

	listResult := handler.storage.NewListResult(list)
	context.JSON(http.StatusOK, listResult)
}

func (handler *Handler) GetLists(context *gin.Context) {
	boardId := context.Param("id")

	lists, err := handler.storage.GetLists(boardId)
	if err != nil {
		context.Status(http.StatusInternalServerError)
		return
	}

	context.JSON(http.StatusOK, gin.H{"lists": lists})
}

func (handler *Handler) RemoveList(context *gin.Context) {
	listId := context.Param("id")

	list, err := handler.storage.GetList(listId)
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
