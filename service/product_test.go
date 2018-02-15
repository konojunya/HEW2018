package service

import (
	"fmt"
	"testing"
)

func TestGetProducts(t *testing.T) {
	products, err := GetProducts()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(products)
}
