package service

import (
	"github.com/konojunya/HEW2018/cache"
	"github.com/konojunya/HEW2018/model"
)

// GetAll プロダクト一覧とエラーを返す
func GetAll() ([]model.Product, error) {
	return cache.Product.GetAll()
}

// GetByID idを元にproductを返す
func GetByID(id uint) (*model.Product, error) {
	product := &model.Product{}
	err := db.First(&product, id).Error
	return product, err
}

// CreateVote 投票する
func CreateVote(id uint) error {
	err := db.Create(&model.Vote{
		ProductID: id,
	}).Error
	if err != nil {
		return err
	}

	go cache.Product.Reload()

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
