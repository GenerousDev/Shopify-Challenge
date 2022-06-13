package main

import (
	"Shopify-Challenge/configs"
	"Shopify-Challenge/routes"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}


func main() {

	fmt.Println("Entry Poin")

	router := gin.Default()
	configs.ConnectDB()
	router.Use(CORSMiddleware())
	//routes
	routes.ItemRoute(router)
	
	port := os.Getenv("PORT")

	router.Run(":" + port)
}
