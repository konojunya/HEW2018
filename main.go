package main

import (
	"os"

	"github.com/konojunya/HEW2018/router"
)

func main() {
	r := router.GetRouter()
	r.Run(":" + os.Getenv("PORT"))
}
