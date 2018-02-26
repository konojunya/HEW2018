package service

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"github.com/JustinTulloss/firebase"
	"github.com/konojunya/HEW2018/model"
)

var (
	endpoint string
	auth     string
	path     string
	client   firebase.Client
)

func init() {
	endpoint = "https://hew2018-9ab24.firebaseio.com"
	auth, err := getToken()
	if err != nil {
		panic(err)
	} else if auth == "" {
		panic("token can not load.")
	}
	path = "/products"
	client = firebase.NewClient(endpoint+path, auth, nil)
}

func getToken() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}

	return os.Getenv("FIREBASE_TOKEN"), nil
}

func productAlloc() interface{} {
	return &model.Product{}
}

// GetProducts プロダクト一覧を取得する
func GetProducts() ([]model.Product, error) {
	var products []model.Product

	for n := range client.Iterator(productAlloc) {
		products = append(products, *n.Value.(*model.Product))
	}

	fmt.Println(products)

	return products, nil
}

// IncrementVote 投票数をインクリメントする
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

// DeleteAllProduct プロダクトをすべて削除する
func DeleteAllProduct() error {
	for n := range client.Iterator(productAlloc) {
		err := client.Remove(n.Key, nil)
		if err != nil {
			return err
		}
	}
	return nil
}

// PostProduct プロダクトを追加する
func PostProduct(product *model.Product) error {
	_, err := client.Push(product, nil)
	if err != nil {
		return err
	}
	return nil
}
