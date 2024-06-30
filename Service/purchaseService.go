package Service

import (
	"errors"
	"strconv"
	"vendingmachine/vendingmachine/Models"
)

var Product = []Models.Product{
	{ID: "1", Title: "Pulpy 1,5L", Price: 20000},
	{ID: "2", Title: "Pulpy 800L", Price: 15000},
	{ID: "3", Title: "Pulpy 250L", Price: 5000},
}

func GetProduct(id string) (*Models.Product, error) {
	for i, b := range Product {
		if b.ID == id {
			return &Product[i], nil
		}
	}
	return nil, errors.New("Product not found")
}

func Calculate(pricePurchase int, priceProduct int) (int, error) {
	if pricePurchase <= priceProduct {
		return 0, errors.New("Uang Tidak Cukup")
	}
	return pricePurchase - priceProduct, nil
}

func CountMoney(money int) (map[string]int, error) {
	var kembalians = map[string]int{
		"50000": 0, "20000": 0, "10000": 0, "5000": 0,
	}

	var pecahan = [4]int{
		50000, 20000, 10000, 5000,
	}

	b := kembalians
	for i := 0; i < len(pecahan); i++ {
		if money-pecahan[i] >= 0 {
			b[strconv.Itoa(pecahan[i])]++
			money = money - pecahan[i]
		}
		if money-pecahan[i] >= 0 {
			i--
		}
	}
	return b, nil
}
