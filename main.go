package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"postsAPI/controllers"
	initializers "postsAPI/initializers"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	route := gin.Default()
	route.POST("/posts", controllers.PostCreate)
	route.GET("/posts", controllers.AllPosts)
	route.GET("/posts/:id", controllers.FindPost)
	route.PUT("/posts/:id", controllers.PostUpdate)
	route.DELETE("/posts/:id", controllers.PostDelete)
	log.Fatal(route.Run(":8080"))
}
