package service

import (
	"log"
	"testing"

	"github.com/konojunya/HEW2018/model"
)

func TestPostProduct(t *testing.T) {
	err := PostProduct(&model.Product{
		ID:        1,
		Title:     "hoge",
		Thumbnail: "/asdf/asdf",
		Author:    "hoge",
		Votes:     0,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetProducts(t *testing.T) {
	products, err := GetProducts()
	if err != nil {
		t.Fatal(err)
	}
	log.Println(products)
}

func TestDeleteProducts(t *testing.T) {
	err := DeleteAllProduct()
	if err != nil {
		t.Fatal(err)
	}
}
