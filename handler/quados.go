package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (handler *Handler) CreateQuado(context *gin.Context) {
	listID := context.PostForm("idList")
	title := context.PostForm("title")
	description := context.PostForm("description")
	// TODO: require some changes when we know more
	// date := context.PostForm("date")
	fakeDate := time.Now()

	quado := handler.storage.NewQuado(listID, title, description, fakeDate)

	handler.storage.PutQuado(quado)

	context.JSON(http.StatusOK, quado)
}

func (handler *Handler) RemoveQuado(context *gin.Context) {
	quadoID := context.Param("id")

	quado, err := handler.storage.GetQuado(quadoID)
	if err != nil {
		context.Status(http.StatusInternalServerError)
		return
	}

	err = handler.storage.DeleteQuado(quado)
	if err != nil {
		context.Status(http.StatusInternalServerError)
		return
	}

	context.Status(http.StatusOK)
}
