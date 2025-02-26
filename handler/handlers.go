package handler

import (
	"coffeeshop/coffee"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//	func Ping(w http.ResponseWriter, r *http.Request) {
//		fmt.Println("Got a ping")
//		io.WriteString(w, "Welcome to the Coffeeshop!\n")
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
	fmt.Println("Printing the list of coffees available in the CoffeeShop ")
	// coffees, err := coffee.GetCoffees()
	// if err != nil {
	// 	fmt.Println("Error while getting coffeelist ", err)
	// 	return
	// }

	// for _, element := range coffees.List {
	// 	result := fmt.Sprintf("%s for $%v", element.Name, element.Price)
	// 	fmt.Println(result)
	// }
	// fmt.Println("Is Latte Available? ", coffee.IsCoffeeAvailable("Latte"))
}
