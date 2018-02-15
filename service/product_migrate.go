package service

import "github.com/konojunya/HEW2018/model"

func GetProductsData() []model.Product {
	var products []model.Product

	products = append(products, model.Product{
		ID:        1,
		Thumbnail: "/image/presen-board.png",
		Title:     "Dinner Match",
		Author:    "1西川 かずき",
		Votes:     10,
	})

	products = append(products, model.Product{
		ID:        2,
		Thumbnail: "/image/presen-board.png",
		Title:     "Dinner Match",
		Author:    "6西川 かずき2",
		Votes:     1,
	})

	products = append(products, model.Product{
		ID:        3,
		Thumbnail: "/image/presen-board.png",
		Title:     "Dinner Match",
		Author:    "5西川 かずき3",
		Votes:     5,
	})

	products = append(products, model.Product{
		ID:        4,
		Thumbnail: "/image/presen-board.png",
		Title:     "Dinner Match",
		Author:    "4西川 かずき4",
		Votes:     7,
	})

	products = append(products, model.Product{
		ID:        5,
		Thumbnail: "/image/presen-board.png",
		Title:     "Dinner Match",
		Author:    "3西川 かずき5",
		Votes:     8,
	})

	products = append(products, model.Product{
		ID:        6,
		Thumbnail: "/image/presen-board.png",
		Title:     "Dinner Match",
		Author:    "7西川 かずき6",
		Votes:     0,
	})

	products = append(products, model.Product{
		ID:        7,
		Thumbnail: "/image/presen-board.png",
		Title:     "Dinner Match",
		Author:    "8西川 かずき7",
		Votes:     0,
	})

	products = append(products, model.Product{
		ID:        8,
		Thumbnail: "/image/presen-board.png",
		Title:     "Dinner Match",
		Author:    "2西川 かずき8",
		Votes:     9,
	})

	return products
}
