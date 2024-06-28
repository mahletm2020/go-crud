package router

import (
	"github.com/gin-gonic/gin"
	"project/models"
)

// User struct represents a user's basic information
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SetupUserRoutes(router *gin.Engine) {
	// Define routes related to users if needed
}
