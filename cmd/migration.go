package main

import "github.com/konojunya/HEW2018/service"

func main() {
	err := service.DeleteAllProduct()
	if err != nil {
		panic(err)
	}
	err = service.PostProduct()
	if err != nil {
		panic(err)
	}
}
