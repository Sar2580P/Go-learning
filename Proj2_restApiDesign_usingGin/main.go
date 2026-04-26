package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)

// this `json:...` allows apis to directly consume the struct
type book struct{
	ID string   `json:"id"`
	Title string   `json:"title"`
	Author string    `json:"author"`
	Quantity int   `json:"quantity"`
}


// dummy database --> slice of book struct
var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func getBooks(c *gin.Context){   // Context --> stores all info related to headers, payload etc...
	c.IndentedJSON(http.StatusOK, books)   // returning book json as response
}

func getBookById(id string) (*book, error){
	for i, b := range books{
		if b.ID==id{
			return &books[i], nil
		}
	}
	return nil, errors.New("No book with this id found.")
}

func bookById(c *gin.Context){
	id:= c.Param("id")  // /book/{id}
	book, err := getBookById((id))

	if err!=nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return 
	}
	c.IndentedJSON(http.StatusOK, book)
}

func createBook(c *gin.Context){
	var newBook book
	if err:= c.BindJSON(&newBook); err!=nil{   // return on getting err
		return 
	}
	books=append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func checkoutBook(c *gin.Context){
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query param"})
		return 
	}

	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return 
	}

	if book.Quantity<=0 {
		c.IndentedJSON(http.StatusBadRequest , gin.H{"message": "Book not found."})
		return 
	}
	book.Quantity -=1 
	c.IndentedJSON(http.StatusOK, book)
}

func returnBook(c *gin.Context){
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query param"})
		return 
	}

	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return 
	}


	book.Quantity +=1 
	c.IndentedJSON(http.StatusOK, book)

}

func main(){
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.GET("/books/:id", bookById)
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/return", returnBook)
	router.Run("localhost:8080")
}