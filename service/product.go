package service

import "github.com/konojunya/HEW2018/model"

// GetProducts プロダクト一覧とエラーを返す
func GetProducts() ([]model.Product, error) {
	products := make([]model.Product, 0)
	return products, nil
}

// IncrementVote 投票する
func IncrementVote(id string) error {
	return nil
}
