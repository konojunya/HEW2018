package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func apiRouter(api *gin.RouterGroup) {

	api.GET("/some", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "hoge",
		})
	})

}
