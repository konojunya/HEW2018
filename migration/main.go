package main

import "github.com/konojunya/HEW2018/model"

func main() {
	db := model.GetDBConn()

	// migrate
	db.DropTableIfExists(&model.Product{})
	db.DropTableIfExists(&model.Vote{})
	db.CreateTable(&model.Product{})
	db.CreateTable(&model.Vote{})

}
