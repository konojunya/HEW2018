package model

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Title string
	Body  string
	User  User
}

type User struct {
	gorm.Model
	Name     string
	Password string
	Email    string
}
