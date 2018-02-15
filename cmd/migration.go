package main

import (
	"log"

	"github.com/konojunya/HEW2018/service"
)

func main() {
	err := service.DeleteAllProduct()
	if err != nil {
		panic(err)
	}
	log.Println("Deleted All Products!")

	for _, product := range service.GetProductsData() {
		err = service.PostProduct(&product)
		if err != nil {
			panic(err)
		}
		log.Printf("Delete:\n\tID:%v\n\tTitle:%v", product.ID, product.Title)
	}
}
