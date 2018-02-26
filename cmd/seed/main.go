package main

import (
	"fmt"

	"github.com/konojunya/HEW2018/model"
	"github.com/konojunya/HEW2018/service"
)

func main() {
	create(model.Product{
		Title:     "dinner match",
		Author:    "西川かずき",
		Votes:     0,
		Thumbnail: "/image/presen-board.png",
	})
}

func create(p model.Product) {
	product, err := service.Create(p)
	if err != nil {
		panic(err)
	}
	fmt.Printf("product created: %v\n", product.CreatedAt)
}
