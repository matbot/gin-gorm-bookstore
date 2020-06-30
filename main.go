package main

import (
	"github.com/gin-gonic/gin"
	"github.com/matbot/gin-gorm-bookstore/controllers"
	"github.com/matbot/gin-gorm-bookstore/models"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	router.GET("/books", controllers.FindBooks)

	_ = router.Run()
}
