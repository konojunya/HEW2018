package service

import (
	"fmt"
	"testing"
)

func TestRanking(t *testing.T) {
	products, _ := GetAll()
	fmt.Println(len(*products.RankingSort().FilterZero().Cut(5)))
}
