package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := "host=salt.db.elephantsql.com user=zgenpbbc password=EuRwgrEsl2UggpSP__Xg8e4wmvvhIzvC dbname=zgenpbbc port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!")
	}
}
