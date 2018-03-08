package router

import (
	"github.com/gin-gonic/gin"
	"github.com/konojunya/HEW2018/controller"
)

func apiRouter(api *gin.RouterGroup) {
	api.GET("/products", controller.GetAllProducts)
	api.GET("/products/:id", controller.GetProductByID)
	api.POST("/products/:id/vote", controller.CreateVote)
	api.POST("/refresh", controller.RefreshCache)
}
