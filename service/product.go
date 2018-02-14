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
