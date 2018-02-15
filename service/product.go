package service

import (
	"io/ioutil"

	"github.com/JustinTulloss/firebase"
	"github.com/konojunya/HEW2018/model"
	yaml "gopkg.in/yaml.v2"
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
	buf, err := ioutil.ReadFile("config.yml")
	if err != nil {
		return "", err
	}

	var config model.Config
	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		return "", err
	}

	return config.Token, nil
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
