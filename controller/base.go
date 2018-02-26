package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BatRequest(err string, c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"err": err,
	})
}

func Json(obj interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, obj)
}
