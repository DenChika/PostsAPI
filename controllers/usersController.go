package controllers

import (
	"github.com/gin-gonic/gin"
	"postsAPI/initializers"
	"postsAPI/models"
)

func CreateUser(ctx *gin.Context) {
	var body struct {
		Username string `json:"username"`
		Name     string `json:"name"`
		Surname  string `json:"surname"`
	}
	err := ctx.Bind(&body)
	if err != nil {
		ctx.Status(400)
		return
	}
	user := models.User{Username: body.Username, Name: body.Name, Surname: body.Surname}
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"user": user,
	})
}

func AllUsers(ctx *gin.Context) {
	var users []models.User
	result := initializers.DB.Find(&users)
	if result.Error != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"users": users,
	})
}

func FindUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var user models.User
	result := initializers.DB.First(&user, id)
	if result.Error != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"user": user,
	})
}

func UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var body struct {
		Username string `json:"username"`
		Name     string `json:"name"`
		Surname  string `json:"surname"`
	}
	err := ctx.Bind(&body)
	if err != nil {
		ctx.Status(400)
		return
	}
	var user models.User
	result := initializers.DB.First(&user, id)
	if result.Error != nil {
		ctx.Status(400)
		return
	}
	initializers.DB.Model(&user).Updates(models.User{Username: body.Username, Name: body.Name, Surname: body.Surname})
	if result.Error != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"user": user,
	})
}

func DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var user models.User
	result := initializers.DB.First(&user, id)
	if result.Error != nil {
		ctx.Status(400)
		return
	}
	initializers.DB.Delete(&models.User{}, id)
	ctx.Status(200)
}
