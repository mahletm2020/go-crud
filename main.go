package main

  import(
	//  "fmt"
   "github.com/gin-gonic/gin"
   "net/http"
	//  "errors"
)
type book struct{
	Id       string    `json:"id"`
	Title    string  `json:"title"`
  Autor    string  `json:"autor"`
	Quantity int `"json:quantity"`
}
 var books = []book{
	{Id: "1",Title: "lalal",Autor: "blabla",Quantity: 2},
	{Id: "2",Title: "lawel",Autor: "wlabla",Quantity: 3},

 }
func getbooks(c *gin.Context){
 c.IndentedJSON(http.StatusOK,books)
}
func createbook(c *gin.Context){
	var newBook book
	if err := c.BindJSON(&newBook);err != nil{
		return
	}
	books = append(books,newBook)
	c.IndentedJSON(http.StatusCreated , newBook)
}

func main(){
router:=gin.Default()
router.GET("/books",getbooks)
router.POST("/books",createbook)

router.Run("localhost:8000")
	}