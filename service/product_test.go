package service

import (
	"fmt"
	"testing"

	"github.com/konojunya/HEW2018/model"
)

func TestToken(t *testing.T) {
	token, err := getToken()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(token)
}
func TestPost(t *testing.T) {
	product := &model.Product{
		ID:        1,
		Thumbnail: "/image/presen-board.png",
		Title:     "Dinner Match",
	}

	_, err := client.Push(product, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetProducts(t *testing.T) {
	products, err := GetProducts()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(products)
}
