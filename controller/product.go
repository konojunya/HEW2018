package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/konojunya/HEW2018/model"
)

func GetProducts(c *gin.Context) {
	var products []model.Product
	for i := 1; i < 10; i++ {
		products = append(products, model.Product{
			ID:        i,
			Thumbnail: "/image/presen-board.png",
			Title:     "Dinner Match",
		})
	}
	c.JSON(http.StatusOK, products)
}
