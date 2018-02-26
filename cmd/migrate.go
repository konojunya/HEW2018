package main

import "github.com/konojunya/HEW2018/model"

func main() {
	db := model.GetDBConn()

	db.DropTableIfExists(&model.Product{})
	db.AutoMigrate(&model.Product{})
}
