package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"../models"
	"../middleware"
)

func SetupAuthRoutes(router *gin.Engine) {
	router.POST("/signup", signUp)
	router.POST("/login", login)
}

func signUp(c *gin.Context) {
	var newUser models.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	// Store the new user in the slice (in-memory for simplicity)
	models.Users = append(models.Users, newUser)
	c.Status(http.StatusCreated)
}

func login(c *gin.Context) {
	var userLogin models.User
	if err := c.BindJSON(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Find the user in the slice (in-memory for simplicity)
	for _, user := range models.Users {
		if user.Username == userLogin.Username && user.Password == userLogin.Password {
			// Generate JWT token for the authenticated user
			token, err := middleware.CreateToken(user.Username)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			// Return the token to the client
			c.JSON(http.StatusOK, gin.H{"token": token})
			return
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
}
