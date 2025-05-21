package tools

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv(filename string) {
	err := godotenv.Load(filename)
	if err != nil {
		log.Fatal("Error loading env file")
	}

}
