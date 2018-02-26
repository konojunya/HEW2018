package main

import (
	"log"
	"sync"

	"github.com/konojunya/HEW2018/model"
	"github.com/konojunya/HEW2018/service"
)

func main() {
	err := service.DeleteAllProduct()
	if err != nil {
		panic(err)
	}
	log.Println("Deleted All Products!")

	var wg sync.WaitGroup
	for _, product := range service.GetProductsData() {
		wg.Add(1)
		go func(product model.Product) {
			if err := service.PostProduct(&product); err != nil {
				panic(err)
			}
			log.Printf("\nCreate:\n\tID: %v\n\tTitle: %v", product.ID, product.Title)
			wg.Done()
		}(product)
	}
	wg.Wait()
}
