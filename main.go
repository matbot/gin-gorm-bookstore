package main

import (
	"github.com/gin-gonic/gin"
	"github.com/matbot/gin-gorm-bookstore/handlers"
	"github.com/matbot/gin-gorm-bookstore/models"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	router.GET("/books", handlers.FindBooks)
	router.GET("/books/:id", handlers.FindBook)
	router.POST("/books", handlers.CreateBook)
	router.PATCH("/books/:id", handlers.UpdateBook)
	router.DELETE("/books/:id", handlers.DeleteBook)

	_ = router.Run()
}
