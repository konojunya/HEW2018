package main

import "github.com/konojunya/HEW2018/model"

func main() {
	db := model.GetDBConn()

	db.DropTableIfExists(&model.Product{})
	db.DropTableIfExists(&model.Vote{})

	db.AutoMigrate(&model.Product{})
	db.AutoMigrate(&model.Vote{})
}
