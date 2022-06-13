package main

import (
	"Shopify-Challenge/configs"

	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("Entry Poin")

	router := gin.Default()

	configs.ConnectDB()
	fmt.Println("Database connected")

//routes
	routes.ItemRoute(router)

	router.Run("localhost:3030")
	fmt.Println("Server running on port 3030")
}
