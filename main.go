package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/konojunya/HEW2018/router"
)

func envLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	envLoad()
	r := router.GetRouter()
	r.Run(":" + os.Getenv("PORT"))
}
