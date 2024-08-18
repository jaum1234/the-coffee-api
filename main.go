package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type brewingMethod struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var brewingMethods = []brewingMethod{
	{ID: 1, Name: "French Press"},
	{ID: 2, Name: "Moka Pot"},
	{ID: 3, Name: "Pour Over"},
}

func getBrewingMethods(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, brewingMethods)
}

func storeBrewingMethod(c *gin.Context) {
	var brewingMethod brewingMethod

	err := c.BindJSON(&brewingMethod)

	if err != nil {
		return
	}

	brewingMethods = append(brewingMethods, brewingMethod)

	c.IndentedJSON(http.StatusCreated, brewingMethod)
}

func getBrewingMethodById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return
	}

	for _, method := range brewingMethods {
		if method.ID == id {
			c.IndentedJSON(http.StatusOK, method)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Brewing method not found."})
}

func main() {
	app := gin.Default()
	app.GET("/brewing-methods", getBrewingMethods)
	app.POST("/brewing-methods", storeBrewingMethod)
	app.GET("/brewing-methods/:id", getBrewingMethodById)

	app.Run("localhost:8080")
}
