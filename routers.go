package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// get_groceries responds with the grocery list
func get_groceries(c *gin.Context) {
	data := read_grocery_data()
	c.IndentedJSON(http.StatusOK, data)
}

func run_routes() {
	router := gin.Default()
	router.GET("/groceries", get_groceries)

	// router.POST("/groceries", post_groceries)

	router.Run("localhost:8080")
}