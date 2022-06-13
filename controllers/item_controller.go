package controllers

import (
	"Shopify-Challenge/configs"
	"Shopify-Challenge/models"
	"context"
	"net/http"
	"time"

	"Shopify-Challenge/responses"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var itemCollection *mongo.Collection = configs.GetCollection(configs.DB, "items")
var validate = validator.New()

func CreateItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var item models.Item
		defer cancel()

		//validate the request body
		if err := c.ShouldBind(&item); err != nil {
			c.JSON(http.StatusBadRequest, responses.ItemResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&item); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.ItemResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		// item has been binded with the request body
		newItem := models.Item{
			Id:           primitive.NewObjectID(),
			ItemName:     item.ItemName,
			Location:     item.Location,
			ItemPrice:    item.ItemPrice,
			ItemBrand:    item.ItemBrand,
			ItemCategory: item.ItemCategory,
		}

		//insert the item into the database and return the item
		result, err := itemCollection.InsertOne(ctx, newItem)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ItemResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})

			return
		}
		c.JSON(http.StatusCreated, responses.ItemResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func DeleteAItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		itemId := c.Param("itemId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(itemId)

		result, err := itemCollection.DeleteOne(ctx, bson.M{"id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ItemResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound, responses.ItemResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Item with specified ID not found!"}})
			return
		}
		c.JSON(http.StatusOK, responses.ItemResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Item successfully deleted!"}})

	}
}

func EditAItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		itemId := c.Param("itemId")
		var item models.Item
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(itemId)

		//validate the request body --- save coresponding c to item
		if err := c.ShouldBind(&item); err != nil {
			c.JSON(http.StatusBadRequest, responses.ItemResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&item); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.ItemResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		update := bson.M{
			"itemname":     item.ItemName,
			"location":     item.Location,
			"itemprice":    item.ItemPrice,
			"itembrand":    item.ItemBrand,
			"itemcategory": item.ItemCategory,
		}
		result, err := itemCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ItemResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//get updated item details
		var updatedItem models.Item
		if result.MatchedCount == 1 {
			err := itemCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedItem)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.ItemResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.ItemResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedItem}})
	}
}

func GetAllItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var items []models.Item
		defer cancel()

		results, err := itemCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ItemResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleItem models.Item
			if err = results.Decode(&singleItem); err != nil {
				c.JSON(http.StatusInternalServerError, responses.ItemResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}

			items = append(items, singleItem)
		}

		c.JSON(http.StatusOK, responses.ItemResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": items}})

	}
}
