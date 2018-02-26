package model

import (
	"time"
)

// Model 標準のmodel
type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

// Product プロダクト
type Product struct {
	Model
	Thumbnail string `json:"thumbnail"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Votes     int    `json:"votes" gorm:"-"`
}

type Vote struct {
	Model
	ProductID uint
}

func (p *Product) SetVote() {
	var votes []Vote
	db.Where("product_id = ?", p.ID).Find(&votes)
	p.Votes = len(votes)
}
