package controller

import (
	"net/http"
	"strconv"

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

func IncrementVote(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	err = service.IncrementVote(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.AbortWithStatus(http.StatusNoContent)

}
