package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *Handler) CreateBoard(c *gin.Context) {
	board := handler.storage.NewBoard()

	handler.storage.PutBoard(board)

	c.JSON(http.StatusOK, board)
}

func (handler *Handler) DeleteBoard(c *gin.Context) {
	boardID := c.Param("id")

	board, err := handler.storage.GetBoard(boardID)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	err = handler.storage.DeleteBoard(board)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
