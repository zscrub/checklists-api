package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// get_groceries responds with the grocery list
func get_groceries(c *gin.Context) {
	data, err := go_get_groceries()
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, err)
	}
	c.IndentedJSON(http.StatusOK, data)
}

func run_routes() {
	router := gin.Default()
	router.GET("/groceries", get_groceries)

	// router.POST("/groceries", post_groceries)

	router.Run("localhost:8080")
}
