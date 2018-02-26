package main

import (
	"github.com/konojunya/HEW2018/model"
	"github.com/konojunya/HEW2018/service"
)

func main() {
	service.Create(model.Product{
		Title:     "dinner match",
		Author:    "西川かずき",
		Votes:     0,
		Thumbnail: "/image/presen-board.png",
	})
}
