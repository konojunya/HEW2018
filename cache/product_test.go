package cache

import (
	"fmt"
	"testing"
)

func TestCache(t *testing.T) {
	products := Product.GetAll()
	fmt.Println(products)
}
