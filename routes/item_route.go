package routes

import (
	"Shopify-Challenge/controllers"

	"github.com/gin-gonic/gin"
)

func ItemRoute(router *gin.Engine) {
	// Rendering templates routes
	router.GET("/items", controllers.GetAllItems())
	// router.GET("item/edit/:itemId", controllers.GetUpdatePage())

	// Handling payloads
	router.POST("item/create", controllers.CreateItem())
	router.PUT("item/edit/:itemId", controllers.EditAItem())
	router.DELETE("item/delete/:itemId", controllers.DeleteAItem())
}
