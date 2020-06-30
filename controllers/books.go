package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/matbot/gin-gorm-bookstore/models"
	"net/http"
)

// GET /books
// Gets all books in database.
func FindBooks(ctx *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)
	ctx.JSON(http.StatusOK, gin.H{"data": books})
}
