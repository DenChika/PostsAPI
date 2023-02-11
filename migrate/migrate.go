package main

import (
	"log"
	initializers "postsAPI/initializers"
	models "postsAPI/models"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	err := initializers.DB.AutoMigrate(&models.User{}, &models.Post{})
	if err != nil {
		log.Fatal("Failed to migrate the database!")
	}
}
