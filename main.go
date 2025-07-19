package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{
	{
		ID: "1",
		Title: "Metamorphosis",
		Author: "Franz Kafka",
	},
	{
		ID: "2",
		Title: "Mathematics for Class X",
		Author: "NCERT",
	},
}

func main(){
	r := gin.Default()
	r.GET("/books", getBooks)
	r.GET("/books/:id", getBookbyID)
	r.POST("/books", createBook)
	r.PUT("/books/:id", updateBook)
	r.DELETE("/books/:id", deleteBook)

	r.Run() // listens on : 8000
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func getBookbyID(c *gin.Context){
	id := c.Param("id")
	for _, b := range books {
		if b.ID == id {
			c.IndentedJSON(http.StatusOK, b)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message":"book not found"})
}

func createBook(c *gin.Context) {
	var newBook Book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func updateBook(c *gin.Context) {
	id := c.Param("id")
	var updatedBook Book
	if err := c.BindJSON(&updatedBook); err != nil{
		return
	}

	for i, b := range books {
		if b.ID == id {
			books[i] = updatedBook
			c.IndentedJSON(http.StatusOK, updatedBook)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "book not found",})
}

func deleteBook(c  *gin.Context) {
	id := c.Param("id")
	for i, b := range books {
		if b.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "book not found"})
}