package handler

import (
	"coffeeshop/coffee"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the Coffeeshop!",
	})
}

func HTMLCoffeeDetails(c *gin.Context) {
	coffee, _ := coffee.GetCoffees()
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"list": coffee.List,
		},
	)
}
func CoffeeDetails(c *gin.Context) {
	coffeelist, _ := coffee.GetCoffees()
	c.String(http.StatusOK, " %s", coffeelist)
}
