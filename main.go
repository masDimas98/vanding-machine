package main

import (
	"github.com/gin-gonic/gin"
	"vendingmachine/vendingmachine/Controller"
)

func main() {
	route := gin.Default()
	route.POST("/purchase", Controller.Purchase)
	route.Run("localhost:8080")
}
