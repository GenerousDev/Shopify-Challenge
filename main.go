package main

import (
	"Shopify-Challenge/configs"
	"Shopify-Challenge/routes"
	"fmt"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("Entry Poin")

	router := gin.Default()
	configs.ConnectDB()

	//routes
	routes.ItemRoute(router)
	// CORS or https://foo.com and https://github.com origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Cedentials share
	// -Preflight requests cached for 12 hours
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders :     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	port := os.Getenv("PORT")

	router.Run(":" + port)
}
