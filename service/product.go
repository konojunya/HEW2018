package service

import (
	"fmt"
	"os"

	"github.com/JustinTulloss/firebase"
	"github.com/joho/godotenv"
	"github.com/konojunya/HEW2018/model"
)

var (
	endpoint string
	auth     string
	client   firebase.Client
)

func init() {
	endpoint = "https://hew2018-9ab24.firebaseio.com"
	auth, err := getToken()
	if err != nil {
		panic(err)
	}
	client = firebase.NewClient(endpoint+"/products", auth, nil)
}

func getToken() (string, error) {
	mode := os.Getenv("MODE")
	var token string

	if mode != "production" {
		if err := godotenv.Load(); err != nil {
			panic(err)
		}
	}

	if token = os.Getenv("FIREBASE_TOKEN"); len(token) == 0 {
		return "", fmt.Errorf("Error: %s", "token length zero.")
	}
	return token, nil
}

func productAlloc() interface{} {
	return &model.Product{}
}

func GetProducts() ([]model.Product, error) {
	var products []model.Product

	for n := range client.Iterator(productAlloc) {
		products = append(products, *n.Value.(*model.Product))
	}

	return products, nil
}

func IncrementVote(id int) error {
	for n := range client.Iterator(productAlloc) {
		product := n.Value.(*model.Product)
		if product.ID == id {
			err := client.Update(n.Key, &model.Product{
				ID:        product.ID,
				Thumbnail: product.Thumbnail,
				Title:     product.Title,
				Author:    product.Author,
				Votes:     product.Votes + 1,
			}, nil)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func DeleteAllProduct() error {
	for n := range client.Iterator(productAlloc) {
		err := client.Remove(n.Key, nil)
		if err != nil {
			return err
		}
	}
	return nil
}

func PostProduct() error {
	for _, product := range getProductsData() {
		_, err := client.Push(product, nil)
		if err != nil {
			return err
		}
	}
	return nil
}
