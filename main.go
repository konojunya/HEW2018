package main

import (
	"github.com/konojunya/HEW2018/router"
)

func main() {
	r := router.GetRouter()
	r.Run(":8000")
}
