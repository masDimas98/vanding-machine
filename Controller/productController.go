package Controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vendingmachine/vendingmachine/Models"
	"vendingmachine/vendingmachine/Service"
)

func Purchase(c *gin.Context) {
	var newPurchase Models.Purchase

	if err := c.BindJSON(&newPurchase); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"mesage": "Product Tidak Terdaftar"})
		return
	}
	prod, err := Service.GetProduct(newPurchase.ID)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Product Tidak Terdaftar"})
		return
	}

	sisaUang, err := Service.Calculate(newPurchase.Price, prod.Price)
	if err != nil {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Uang Kurang"})
		return
	}

	kembalian, err := Service.CountMoney(sisaUang)
	if err != nil {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Money less that 5000"})
		return
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"kembalian": kembalian})
		return
	}
}
