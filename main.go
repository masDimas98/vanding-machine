package main

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"vendingmachine/vendingmachine/Controller"
	"vendingmachine/vendingmachine/Models"
	"vendingmachine/vendingmachine/database"
)

func main() {
	route := gin.Default()
	route.GET("/testConnection", func(context *gin.Context) {
		var test []Models.Test
		find, _ := database.Db.Collection("test").Find(context, bson.D{})
		err := find.All(context, &test)
		if err != nil {
			context.JSON(404, gin.H{})
		}
		context.JSON(200, test)
	})
	route.POST("/purchase", Controller.Purchase)
	route.Run("localhost:8080")
}
