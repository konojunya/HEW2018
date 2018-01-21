package main

import (
	"github.com/konojunya/HEW2018/router"
	"os"
)

func main() {
	r := router.GetRouter()
	r.Run(":" + os.Getenv("PORT"))
}
