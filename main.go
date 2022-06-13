package main

import (
	"Shopify-Challenge/configs"
	"Shopify-Challenge/routes"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	fmt.Println("Entry Poin")

	router := gin.Default()
	configs.ConnectDB()

	//routes
	routes.ItemRoute(router)
	port := os.Getenv("PORT")

	router.Run(":"+ port )
}
