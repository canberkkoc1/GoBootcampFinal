package main

import (
	"ck/configs"
	"ck/controller"
	"ck/middlewares"
	"ck/models"

	"github.com/gin-gonic/gin"
)

func main() {

	db := configs.NewMySQLDB("root:123456yhN@tcp(127.0.0.1:3306)/picus?parseTime=True&loc=Local")

	err := db.AutoMigrate(&models.User{}, &models.Category{}, &models.Products{})

	if err != nil {
		panic(err.Error())
	}

	var router = gin.Default()

	router.POST("/createUser", controller.CreateUser)
	router.POST("/login", controller.Login)
	router.POST("/addCategory", controller.CreateCategory)
	router.GET("/categories", controller.GetCategories)

	//* Admin Middleware
	router.Use(middlewares.AuthJWTAdmin())
	router.POST("/addproduct", controller.CreateProduct)

	router.Run()

}
