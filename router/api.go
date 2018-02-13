package router

import (
	"github.com/gin-gonic/gin"
	"github.com/konojunya/HEW2018/controller"
)

func apiRouter(api *gin.RouterGroup) {
	api.GET("/products", controller.GetProducts)
}
