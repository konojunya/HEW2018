package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/konojunya/HEW2018/cache"
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

// GetRankedProducts ランキングを返す
func GetRankedProducts(c *gin.Context) {
	products, err := service.GetAll()
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, products.SortByVote().FilterZero().Cut(5))
}

// GetRankingForMiyauchi 宮内先生に見せるためのデータをとってくるAPI
func GetRankingForMiyauchi(c *gin.Context) {
	products, err := service.GetAll()
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, products.SortByVote())
}

// CreateVote 投票する
func CreateVote(c *gin.Context) {
	id, err := GetUint(c, "id")
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	err = service.CreateVote(id)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.AbortWithStatus(http.StatusNoContent)
}

// RefreshCache キャッシュをリフレッシュ
func RefreshCache(c *gin.Context) {
	cache.RefreshAll()
}
