package intializers

import (
	"log"

	"example.com/maphile/models"
)

func Migrate() {
	DB.Exec("CREATE TYPE user_role AS ENUM ('admin', 'sales', 'shop_owner', 'shop_keeper');")
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("error happens in migration")
	}
	log.Println("Migration success")
}