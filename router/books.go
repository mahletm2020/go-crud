package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"../models"
	"../middleware"
)

func SetupBookRoutes(router *gin.Engine) {
	protected := router.Group("/books")
	protected.Use(middleware.AuthMiddleware()) // Apply JWT middleware
	{
		protected.GET("/", getBooks)
		protected.GET("/:id", getBookByID)
		protected.POST("/", postBook)
		protected.PUT("/:id", updateBook)
		protected.DELETE("/:id", deleteBook)
	}
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Books)
}

func getBookByID(c *gin.Context) {
	id := c.Param("id")
	for _, b := range models.Books {
		if b.Id == id {
			c.IndentedJSON(http.StatusOK, b)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func postBook(c *gin.Context) {
	var newBook models.Book
	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.Books = append(models.Books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func updateBook(c *gin.Context) {
	id := c.Param("id")
	var updatedBook models.Book
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, b := range models.Books {
		if b.Id == id {
			models.Books[i] = updatedBook
			c.JSON(http.StatusOK, updatedBook)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func deleteBook(c *gin.Context) {
	id := c.Param("id")
	for i, b := range models.Books {
		if b.Id == id {
			models.Books = append(models.Books[:i], models.Books[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "book deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}
