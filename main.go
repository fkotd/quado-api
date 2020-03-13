package main

import (
	"github.com/finalKickOfTheDeath/quado-api/handler"
	"github.com/finalKickOfTheDeath/quado-api/storage"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}
	router.Use(cors.New(config))

	storage := storage.NewStorage(storage.NewConfig("quado.db", 0600))

	handler := handler.NewHandler(storage)

	storage.Open()
	storage.InitBuckets()

	router.POST("/boards", handler.CreateBoard)
	router.DELETE("/boards/:id", handler.RemoveBoard)

	router.POST("/boards/:id/lists", handler.CreateList)
	router.DELETE("/lists/:id", handler.RemoveList)

	router.POST("/quados", handler.CreateQuado)
	router.DELETE("/quados/:id", handler.RemoveQuado)

	return router
}

func main() {

	r := setupRouter()

	r.Run(":3333")
}
