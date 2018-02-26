package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/makki0205/config"
)

var db = NewDBConn()

func NewDBConn() *gorm.DB {
	db, err := gorm.Open(GetDBConfig())
	if err != nil {
		panic(err)
	}
	return db
}

func GetDBConn() *gorm.DB {
	return db
}

func GetDBConfig() (string, string) {
	return config.Env("dialect"), config.Env("datasource")
}
