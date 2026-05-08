package intializers

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB 

func Database_connection() {
	var err error 

	dsn := os.Getenv("DB_URL") 
	DB , err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect the Database")
		}
    fmt.Println("Database connection successfully")
}