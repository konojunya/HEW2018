package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/js", "./public/js")
	r.Static("/css", "./public/css")
	r.Static("/image", "./public/image")

	r.LoadHTMLGlob("view/*")

	// web routing
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	api := r.Group("/api")
	apiRouter(api)
	return r

}
