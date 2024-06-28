package middleware

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"../models"
	"fmt"
)

var jwtKey = []byte("your_secret_key_here")

// JWT claims struct
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Function to create a new JWT token
func CreateToken(username string) (string, error) {
	// Set expiration time of the token
	expirationTime := time.Now().Add(5 * time.Minute) // Token expires in 5 minutes

	// Create the JWT claims, which include username and expiration time
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	// Create JWT token using claims and the HMAC SHA256 signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key and get the complete signed token as a string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", fmt.Errorf("error generating token: %w", err)
	}

	return tokenString, nil
}

// Middleware function to verify JWT token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get token from Authorization header
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			// If the token is not provided, return an unauthorized error
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort() // Stop further processing
			return
		}

		// Parse and validate JWT token
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			// Return the secret key used to sign the token
			return jwtKey, nil
		})

		// Check if there was an error parsing the token or if the token is invalid
		if err != nil || !token.Valid {
			// If the token is invalid or there was an error, return an unauthorized error
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort() // Stop further processing
			return
		}

		// Token is valid, proceed to the next middleware or handler
		c.Next()
	}
}
