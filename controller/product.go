package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/konojunya/HEW2018/service"
)

// GetAllProducts プロダクト一覧を返す
func GetAllProducts(c *gin.Context) {
	products, err := service.GetAll()
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
