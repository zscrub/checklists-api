package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"zscrub/checklists-api/structure"

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

func post_groceries(c *gin.Context) {
	byte_arr, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, err)
	}

	var data []structure.GroceryChecklist
	err = json.Unmarshal(byte_arr, &data)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(data)

	err = go_post_groceries(data)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, err)
	}

	c.IndentedJSON(http.StatusCreated, data)
}

func run_routes() {
	router := gin.Default()
	router.GET("/groceries", get_groceries)

	router.POST("/groceries", post_groceries)
	router.Run("localhost:8080")
}
