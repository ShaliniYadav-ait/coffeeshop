package main

import (
	"coffeeshop/handler"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.GET("/home", handler.HTMLCoffeeDetails)
	r.GET("/ping", handler.Ping)
	r.GET("/coffee", handler.CoffeeDetails)
	fmt.Println("Starting the server at port 8081")
	r.Run(":8081")
}
