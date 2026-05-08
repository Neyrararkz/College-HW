package handlers

import (
	"lesson13-Book-Library-API/models"
	"lesson13-Book-Library-API/repositories"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	books, _ := repositories.GetAllBooks()
	c.JSON(http.StatusOK, books)
}

func GetBook(c *gin.Context) {
	id := c.Param("id")
	book, err := repositories.GetBookByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	repositories.CreateBook(&book)
	c.JSON(http.StatusCreated, book)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	repositories.UpdateBook(id, &book)
	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	repositories.DeleteBook(id)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}