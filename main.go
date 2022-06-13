package main

import (
	"Shopify-Challenge/configs"
	"Shopify-Challenge/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("Entry Poin")

	router := gin.Default()
	configs.ConnectDB()

	//routes
	routes.ItemRoute(router)

	router.Run("0.0.0.0:3000")
}
