package service

import (
	"github.com/konojunya/HEW2018/model"
)

// GetAll プロダクト一覧とエラーを返す
func GetAll() ([]model.Product, error) {
	products := make([]model.Product, 0)
	err := db.Find(&products).Error
	for key, product := range products {
		product.SetVote()
		products[key] = product
	}

	if err != nil {
		return nil, err
	}

	return products, nil
}

// CreateVote 投票する
func CreateVote(id uint) error {
	err := db.Create(&model.Vote{
		ProductID: id,
	}).Error
	if err != nil {
		return err
	}

	return nil
}

// Create プロダクトを作成する
func Create(product model.Product) (model.Product, error) {
	err := db.Create(&product).Error
	if err != nil {
		return model.Product{}, err
	}
	return product, nil
}
