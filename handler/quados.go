package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type QuadoJSON struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	// TODO: add deadline and repeat
}

func (handler *Handler) CreateQuado(context *gin.Context) {
	listId := context.Param("id")
	fakeDate := time.Now()

	log.SetLevel(log.DebugLevel)

	var quadoJSON QuadoJSON
	if err := context.ShouldBindJSON(&quadoJSON); err != nil {
		context.Status(http.StatusBadRequest)
		log.Debug("Binding error")
		return
	}

	if _, err := handler.storage.GetList(listId); err != nil {
		context.Status(http.StatusBadRequest)
		log.WithFields(log.Fields{
			"listId": listId,
		}).Debug("List not found error")
		return
	}

	quado := handler.storage.NewQuado(listId, quadoJSON.Title, quadoJSON.Description, fakeDate)
	handler.storage.PutQuado(quado)

	quadoResult := handler.storage.NewQuadoResult(quado)
	context.JSON(http.StatusOK, quadoResult)

	log.WithFields(log.Fields{
		"id":     quadoResult.Id,
		"listId": quadoResult.ListId,
		"title":  quadoResult.Title,
		"desc":   quadoResult.Description,
	}).Debug("New quado result")
}

func (handler *Handler) RemoveQuado(context *gin.Context) {
	quadoId := context.Param("id")

	quado, err := handler.storage.GetQuado(quadoId)
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
