package intializers

import (
	"log"

	"github.com/joho/godotenv"
)
func LoadEnvfile() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
}