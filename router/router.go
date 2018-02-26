package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetRouter ルーターを設定してgin.Engineを返す
func GetRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/js", "./public/js")
	r.Static("/image", "./public/image")
	r.Static("/css", "./public/css")

	r.LoadHTMLGlob("view/*")
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	api := r.Group("/api")
	apiRouter(api)

	return r
}
