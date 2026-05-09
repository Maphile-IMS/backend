package intializers

import (
	"log"

	"example.com/maphile/models"
)

func Migrate() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("error happens in migration")
	}
	log.Println("Migration success")
}