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
	route.POST("/posts", controllers.CreatePost)
	route.POST("/users/:id/posts", controllers.CreatePostByUser)
	route.GET("/users/:id/posts", controllers.GetPostsByUser)
	route.GET("/posts", controllers.AllPosts)
	route.GET("/posts/:id", controllers.FindPost)
	route.PUT("/posts/:id", controllers.UpdatePost)
	route.DELETE("/posts/:id", controllers.DeletePost)

	route.POST("/users", controllers.CreateUser)
	route.GET("/users", controllers.AllUsers)
	route.GET("/users/:id", controllers.FindUser)
	route.PUT("/users/:id", controllers.UpdateUser)
	route.DELETE("/users/:id", controllers.DeleteUser)

	log.Fatal(route.Run(":8080"))
}
