package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *Handler) CreateBoard(context *gin.Context) {
	board := handler.storage.NewBoard()

	handler.storage.PutBoard(board)

	context.JSON(http.StatusOK, board)
}

func (handler *Handler) RemoveBoard(context *gin.Context) {
	boardId := context.Param("id")

	board, err := handler.storage.GetBoard(boardId)
	if err != nil {
		context.Status(http.StatusInternalServerError)
		return
	}

	err = handler.storage.DeleteBoard(board)
	if err != nil {
		context.Status(http.StatusInternalServerError)
		return
	}

	context.Status(http.StatusOK)
}
