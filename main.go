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

	func getbooks(c *gin.Context){
	c.IndentedJSON(http.StatusOK,books)
	}









	// postAlbums adds an album from JSON received in the request body.
		func postbooks(c *gin.Context) {
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













// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	//post
	for _, a := range books {
			if a.Id == id {
					c.IndentedJSON(http.StatusOK, a)
					return
			}
	}
	c.IndentedJSON(http.StatusNotFound,
		 gin.H{"message": "album not found"})
}








	// func bookById(c *gin.Context){
	// 	id := c.Param("/id")
	// 	book,err := getBookById(id)

	// 	if err !=nil{
	// 		return
	// 	}
	// 	c.IndentedJSON(http.StatusOK,book)
	// }


	// func getbook(id string)(*book ,error){
	//    for i, b:=range books{
	// 		if b.ID ==id{
	// 			return &books[i],nil
	// 		  }
	// 	 }
	// 	 return nil,errors.New("book not found")
	// }

	func createbook(c *gin.Context){//context carries req nd response detailes in go
		var newBook book //book here is type of vaiabl aka newbook which is alredy defind array
		if err := c.BindJSON(&newBook);err != nil{
			return // When used without any value, it simply stops further execution of the function.
		}
		books = append(books,newBook)
		c.IndentedJSON(http.StatusCreated , newBook)
	}

	func main(){
	router:=gin.Default() //Initialize a Gin router using Default.
	//router.GET("/books/:id",bookById)
	router.GET("/books",getbooks)
	router.POST("/books",createbook)

	router.Run("localhost:8080")
		}