package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/konojunya/HEW2018/service"
)

func GetProducts(c *gin.Context) {
	products, err := service.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, products)
}
