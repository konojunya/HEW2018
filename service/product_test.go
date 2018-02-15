package service

import (
	"fmt"
	"testing"
)

func TestToken(t *testing.T) {
	token, err := getToken()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(token)
}
func TestPost(t *testing.T) {
	for _, product := range getProductsData() {
		_, err := client.Push(product, nil)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestGetProducts(t *testing.T) {
	products, err := GetProducts()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(products)
}
