package main

import (
	"fmt"
	"session8/config"
	"session8/models"
)

func main() {
	config.Connect()
}

func Demo1() {
	db, err := config.Connect()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.productModel{
			Db: db,
		}
		products, err2 := productModel.FindAll()
		if err2 != nil {
			fmt.Println(err2)
		}else {
			for _, product := range products {
				fmt.Println(product.ToString())
				fmt.Println("-----------------")
			}
	}
}

func Demo3() {
	db, err := config.Connect()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.productModel{
			Db: db,
		}
		product := entities.Product{
			Name: "abc"
			Price: 111,
			Quantity: 22,
		}
		err2 := productModel.Create(&product)
		if err2 != nil{
			fmt.Println(err2)
		}else{
			fmt.Println(product.ToString())
		}
}