package main

import "github.com/konojunya/HEW2018/service"

func main() {
	err := service.DeleteAllProduct()
	if err != nil {
		panic(err)
	}
	for _, product := range service.GetProductsData() {
		err = service.PostProduct(&product)
		if err != nil {
			panic(err)
		}
	}
}
