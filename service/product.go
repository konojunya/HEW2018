package service

import (
	"os"

	"github.com/JustinTulloss/firebase"
	"github.com/konojunya/HEW2018/model"
)

var (
	endpoint string
	auth     string
	client   firebase.Client
)

func init() {
	endpoint = ""
	auth = os.Getenv("FIREBASE_AUTH_TOKEN")
	client = firebase.NewClient(endpoint+"/products", auth, nil)
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
