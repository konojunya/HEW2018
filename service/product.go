package service

import "github.com/konojunya/HEW2018/model"

// GetAll プロダクト一覧とエラーを返す
func GetAll() ([]model.Product, error) {
	products := make([]model.Product, 0)
	return products, nil
}

// IncrementVote 投票する
func IncrementVote(id string) error {
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
