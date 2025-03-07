package main

import (
	"coffeeshop/handler"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	portNumber := os.Getenv("APP_PORT")
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.GET("/home", handler.HTMLCoffeeDetails)
	r.GET("/ping", handler.Ping)
	r.GET("/coffee", handler.CoffeeDetails)
	fmt.Println("Starting the server at the port number ", portNumber)
	r.Run(fmt.Sprintf(":%s", portNumber))        
}
