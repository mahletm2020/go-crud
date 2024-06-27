			package main

				import(
				//  "fmt"
				"github.com/gin-gonic/gin"
				"net/http"
        " 	github.com/dgrijalva/jwt-go"
				//"errors"
			)
			var jwtKey = []byte("your_secret_key_here")
			
			// User struct represents a user's basic information
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Slice to store registered users (in-memory )
var users []User

// JWT claims struct
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}



// Function to create a new JWT token
func createToken(username string) (string, error) {
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




// Sign-up endpoint
func signUp(c *gin.Context) {
	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	// Store the new user in the slice (in-memory for simplicity)
	users = append(users, newUser)
	c.Status(http.StatusCreated)
}


//login end point
func login(c *gin.Context) {
	var userLogin User
	if err := c.BindJSON(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Find the user in the slice (in-memory for simplicity)
	for _, user := range users {
		if user.Username == userLogin.Username && user.Password == userLogin.Password {
			// Generate JWT token for the authenticated user
			token, err := createToken(user.Username)
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


			type book struct{
				Id       string  `json:"id"`
				Title    string  `json:"title"`
				Autor    string  `json:"autor"`
				Price int    `json:"Price"` 
					}	
			var books = []book{
				{Id: "1",Title: "lalal",Autor: "blabla",Price: 2},
				{Id: "2",Title: "lawel",Autor: "wlabla",Price: 3},

			}



		//get point 
			func getbooks(c *gin.Context){
			c.IndentedJSON(http.StatusOK,books)
			}




			// getAlbumByID
		// parameter sent by the client, then returns that album as a response.
		func getbooksbyid(c *gin.Context) {
			id := c.Param("id")
			// Loop over the list of albums, looking for
			// an album whose ID value matches the parameter.
			for _, a := range books {
					if a.Id == id {
							c.IndentedJSON(http.StatusOK, a)
							return
					}
			}
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		}




		// postAlbums adds an album from JSON received in the request body.
		func postbook(c *gin.Context) {
				var newbook book

				// Call BindJSON to bind the received JSON to
				// newAlbum.
				if err := c.BindJSON(&newbook); err != nil {
						return
				}
			// Add the new album to the slice.
				books = append(books, newbook)
				c.IndentedJSON(http.StatusCreated, newbook)
		}


		func bookupdate(c *gin.Context) {
				id := c.Param("id")
				var updatedbook book
				if err := c.ShouldBindJSON(&updatedbook); err != nil {
						c.JSON(400, gin.H{"error": err.Error()})
						return
				}

				for i, b := range books {
						if b.Id == id {
								books[i] = updatedbook
								c.JSON(http.StatusOK, updatedbook)
								return
						}
				}
				c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
		}

		// Delete a book by ID
		func deletebook(c *gin.Context) {
			id := c.Param("id")
			for i, b := range books {
				if b.Id == id {
					books = append(books[:i], books[i+1:]...)
					c.JSON(http.StatusOK, gin.H{"message": "book deleted"})
					return
				}
			}
			c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
		}



			func main(){
				router:=gin.Default() //Initialize a Gin router using Default.
				//router.GET("/books/:id",bookById)
				router.GET("/books",getbooks)
				router.GET("/books/:id", getbooksbyid)
				router.POST("/books",postbook)
				router.PUT("/books/:id",bookupdate)
				router.DELETE("/books/:id",deletebook)
				


				router.Run("localhost:8088")
					}





package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("your_secret_key_here")

// User struct represents a user's basic information
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Slice to store registered users (in-memory for simplicity)
var users []User

// JWT claims struct
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Function to create a new JWT token
func createToken(username string) (string, error) {
	// Set expiration time of the token
	expirationTime := time.Now().Add(5 * time.Minute) // Token expires in 5 minutes

	// Create the JWT claims, which include username and expiration time
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "your_app_name",
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

// Sign-up endpoint
func signUp(c *gin.Context) {
	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Store the new user in the slice (in-memory for simplicity)
	users = append(users, newUser)

	c.Status(http.StatusCreated)
}

// Log-in endpoint
func login(c *gin.Context) {
	var userLogin User
	if err := c.BindJSON(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Find the user in the slice (in-memory for simplicity)
	for _, user := range users {
		if user.Username == userLogin.Username && user.Password == userLogin.Password {
			// Generate JWT token for the authenticated user
			token, err := createToken(user.Username)
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

	// Middleware function to verify JWT token
	func authMiddleware() gin.HandlerFunc {
		return func(c *gin.Context) {
			// Get token from Authorization header
			tokenString := c.GetHeader("Authorization")
			if tokenString == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
				c.Abort()
				return
			}

			// Parse and validate JWT token
			token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})

			if err != nil || !token.Valid {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				c.Abort()
				return
			}

			// Token is valid, proceed to the next middleware or handler
			c.Next()
		}
	}

type book struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Autor string `json:"autor"`
	Price int    `json:"Price"`
}

var books = []book{
	{Id: "1", Title: "lalal", Autor: "blabla", Price: 2},
	{Id: "2", Title: "lawel", Autor: "wlabla", Price: 3},
}

// Endpoint to get all books
func getbooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// Endpoint to get book by ID
func getbooksbyid(c *gin.Context) {
	id := c.Param("id")
	for _, b := range books {
		if b.Id == id {
			c.IndentedJSON(http.StatusOK, b)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

// Endpoint to add a new book
func postbook(c *gin.Context) {
	var newbook book
	if err := c.BindJSON(&newbook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Add the new book to the slice
	books = append(books, newbook)
	c.IndentedJSON(http.StatusCreated, newbook)
}

// Endpoint to update a book by ID
func bookupdate(c *gin.Context) {
	id := c.Param("id")
	var updatedbook book
	if err := c.ShouldBindJSON(&updatedbook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, b := range books {
		if b.Id == id {
			books[i] = updatedbook
			c.JSON(http.StatusOK, updatedbook)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

// Endpoint to delete a book by ID
func deletebook(c *gin.Context) {
	id := c.Param("id")
	for i, b := range books {
		if b.Id == id {
			books = append(books[:i], books[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "book deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func main() {
	router := gin.Default()

	// Sign-up and login endpoints
	router.POST("/signup", signUp)
	router.POST("/login", login)

	// Protected endpoints (require valid JWT)
	protected := router.Group("/books")
	protected.Use(authMiddleware())
	{
		protected.GET("/", getbooks)
		protected.GET("/:id", getbooksbyid)
		protected.POST("/", postbook)
		protected.PUT("/:id", bookupdate)
		protected.DELETE("/:id", deletebook)
	}

	router.Run(":8088")
}

			