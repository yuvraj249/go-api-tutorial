package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "The Silent Patient", Author: "Alex Michaelides", Quantity: 5},
	{ID: "2", Title: "Educated", Author: "Tara Westover", Quantity: 6},
	{ID: "3", Title: "Atomic Habits", Author: "James Clear", Quantity: 7},
	{ID: "4", Title: "The Midnight Library", Author: "Matt Haig", Quantity: 8},
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

func GetBookID(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}

	}
	return nil, errors.New("error while getting book")
}

func BookID(c *gin.Context) {
	id := c.Param("id")
	book, err := GetBookID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func CheckoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id in query parameter"})
		return
	}
	book, err := GetBookID(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id in books"})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id in books"})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}
func ReturnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id in query parameter"})
		return
	}
	book, err := GetBookID(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id in books"})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id in books"})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}

func main() {
	router := gin.Default()
	router.GET("/books", GetBooks)
	router.GET("/books/:id", BookID)
	router.PATCH("/checkout", CheckoutBook)
	router.PATCH("/return", ReturnBook)
	router.POST("/books", CreateBooks)

	router.Run("localhost:8082")
}
