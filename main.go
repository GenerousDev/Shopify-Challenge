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
	fmt.Println("Database connected")

//routes
		routes.temRoute(router)

router.Run("localhost:3030")
	fmt.Println("Server running n port 3030")
}
