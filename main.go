package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: 1, Title: "The Silent Patient", Author: "Alex Michaelides", Quantity: 5},
	{ID: 2, Title: "Educated", Author: "Tara Westover", Quantity: 6},
	{ID: 3, Title: "Atomic Habits", Author: "James Clear", Quantity: 7},
	{ID: 4, Title: "The Midnight Library", Author: "Matt Haig", Quantity: 8},
}

func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)

}

func CreateBooks(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()
	router.GET("/books", GetBooks)
	router.POST("/books", CreateBooks)
	router.Run("localhost:8081")

}
