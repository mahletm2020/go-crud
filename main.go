	package main

		import(
		//  "fmt"
		"github.com/gin-gonic/gin"
		"net/http"
		//"errors"
	)
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







	func main(){
		router:=gin.Default() //Initialize a Gin router using Default.
		//router.GET("/books/:id",bookById)
		router.GET("/books",getbooks)
		router.GET("/books/:id", getbooksbyid)
		router.POST("/books",postbook)
		router.PUT("/books/:id",bookupdate)
	  


		router.Run("localhost:8088")
			}






	