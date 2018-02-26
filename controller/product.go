package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/konojunya/HEW2018/service"
)

// GetProducts プロダクト一覧を返す
func GetProducts(c *gin.Context) {
	products, err := service.GetProducts()
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, products)
}

// IncrementVote 投票する
func IncrementVote(c *gin.Context) {
	id := c.Param("id")
	err := service.IncrementVote(id)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.AbortWithStatus(http.StatusNoContent)
}
