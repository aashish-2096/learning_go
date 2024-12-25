package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Year   string  `json:"year"`
	Price  float64 `json:"price"`
}

type ApiResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var books = []Book{
	{ID: "1", Title: "1984", Author: "George Orwell", Price: 9.99},
	{ID: "2", Title: "Brave New World", Author: "Aldous Huxley", Price: 12.50},
}

func getBooks(c *gin.Context) {
	if len(books) > 0 {
		response := ApiResponse{
			Status:  "success",
			Message: "Fetch is complete",
			Data:    books,
		}
		c.JSON(http.StatusOK, response)
	} else {
		response := ApiResponse{
			Status:  "failure",
			Message: "Fetch is complete",
			Data:    nil,
		}
		c.JSON(http.StatusNotFound, response)
	}
	return
}

func main() {
	router := gin.Default()
	router.GET("/getBooks", getBooks)
	router.Run("localhost:8080")
}
