package main

import (
	"coffeeshop/handler"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	// r.GET("/home", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.html", nil)
	// })
	r.GET("/home", handler.HTMLCoffeeDetails)
	r.GET("/ping", handler.Ping)
	r.GET("/coffee", handler.CoffeeDetails)
	fmt.Println("Starting the server at port 8081")
	r.Run(":8081")

	// http.HandleFunc("/ping", handler.Ping)
	// fmt.Println("Starting the server at port 8081")
	// err := http.ListenAndServe(":8081", nil)
	// if err != nil {
	// 	fmt.Println("Error while starting the server ", err)
	// }
}
