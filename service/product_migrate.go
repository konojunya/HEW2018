package service

import "github.com/konojunya/HEW2018/model"

func getProductsData() []model.Product {
	var products []model.Product

	products = append(products, model.Product{
		ID:        1,
		Thumbnail: "/image/presen-board.png",
		Title:     "Dinner Match",
		Votes:     0,
	})

	products = append(products, model.Product{
		ID:        2,
		Thumbnail: "/image/presen-board.png",
		Title:     "Dinner Match",
		Votes:     0,
	})

	return products
}
