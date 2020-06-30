package handlers

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

// GET /books/:id
// Get a specific book.
func FindBook(ctx *gin.Context) {
	var book models.Book
	err := models.DB.Where("id = ?", ctx.Param("id")).First(&book).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found."})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": book})
}

// POST /books
// Create a new book and add to database.
// Database access object schema
type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

func CreateBook(ctx *gin.Context) {
	// Validate incoming payload
	var input CreateBookInput
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Add book to DB.
	book := models.Book{
		Title:  input.Title,
		Author: input.Author,
	}
	models.DB.Create(&book)

	ctx.JSON(http.StatusOK, gin.H{"data": book})
}

// PATCH /books/:id
// Update a specific book's information.
// Could use a single access object schema if opt for PUT, requiring idempotent payload.
type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func UpdateBook(ctx *gin.Context) {
	var book models.Book
	err := models.DB.Where("id = ?", ctx.Param("id")).First(&book).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found."})
		return
	}
	var input UpdateBookInput
	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Model(&book).Updates(input)
	ctx.JSON(http.StatusOK, gin.H{"data": book})
}

// DELETE /books/:id
// Delete a book's record.
func DeleteBook(ctx *gin.Context) {
	var book models.Book
	err := models.DB.Where("id = ?", ctx.Param("id")).First(&book).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found."})
		return
	}
	models.DB.Delete(&book)
	ctx.JSON(http.StatusOK, gin.H{"data": true})
}
