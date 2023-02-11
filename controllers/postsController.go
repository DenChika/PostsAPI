package controllers

import (
	"github.com/gin-gonic/gin"
	"postsAPI/initializers"
	"postsAPI/models"
)

func CreatePost(ctx *gin.Context) {
	var body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}
	err := ctx.Bind(&body)
	if err != nil {
		ctx.Status(400)
		return
	}

	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"post": post,
	})
}

func CreatePostByUser(ctx *gin.Context) {
	var body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}
	err := ctx.Bind(&body)
	if err != nil {
		ctx.Status(400)
		return
	}

	post := models.Post{Title: body.Title, Body: body.Body, UserId: ctx.Param("id")}
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"post": post,
	})
}

func GetPostsByUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var posts []models.Post
	result := initializers.DB.Where("user_id = ?", id).Find(&posts)
	if result.Error != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"posts": posts,
	})
}

func AllPosts(ctx *gin.Context) {
	var posts []models.Post
	result := initializers.DB.Find(&posts)
	if result.Error != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"posts": posts,
	})
}

func FindPost(ctx *gin.Context) {
	id := ctx.Param("id")
	var post models.Post
	result := initializers.DB.First(&post, id)
	if result.Error != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"post": post,
	})
}

func UpdatePost(ctx *gin.Context) {
	id := ctx.Param("id")
	var body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}
	err := ctx.Bind(&body)
	if err != nil {
		ctx.Status(400)
		return
	}

	var post models.Post
	result := initializers.DB.First(&post, id)
	if result.Error != nil {
		ctx.Status(400)
		return
	}

	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	ctx.JSON(200, gin.H{
		"post": post,
	})
}

func DeletePost(ctx *gin.Context) {
	id := ctx.Param("id")
	var post models.Post
	result := initializers.DB.First(&post, id)
	if result.Error != nil {
		ctx.Status(400)
		return
	}
	initializers.DB.Delete(&models.Post{}, id)
	ctx.Status(200)
}
