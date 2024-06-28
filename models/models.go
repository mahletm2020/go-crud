package models

// User struct represents a user's basic information
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Book struct represents a book's information
type Book struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Price  int    `json:"price"`
}

// Slice to store registered users (in-memory)
var Users []User

// Slice to store books (in-memory)
var Books = []Book{
	{Id: "1", Title: "lalal", Author: "blabla", Price: 2},
	{Id: "2", Title: "lawel", Author: "wlabla", Price: 3},
}
