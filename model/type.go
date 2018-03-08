package model

import (
	"sort"
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

type Products []Product

// Vote 投票数
type Vote struct {
	Model
	ProductID uint
}

// SetVote 投票をカウントして入れる
func (p *Product) SetVote() {
	var votes []Vote
	db.Where("product_id = ?", p.ID).Find(&votes)
	p.Votes = len(votes)
}

// FilterZero 0を排除する
func (p *Products) FilterZero() *Products {
	var products Products
	for _, product := range *p {
		if product.Votes != 0 {
			products = append(products, product)
		}
	}

	return &products
}

// RankingSort ランキングしたProductsを返す
func (p *Products) RankingSort() *Products {
	sort.Slice(*p, func(i, j int) bool {
		return (*p)[i].Votes > (*p)[j].Votes
	})
	return p
}

// Cut 指定の数でプロダクトを減らす
func (p *Products) Cut(count int) *Products {
	var products Products
	for key, product := range *p {
		if key == count {
			break
		}
		products = append(products, product)
	}
	return &products
}
