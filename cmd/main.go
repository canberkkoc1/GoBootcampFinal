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

	err := db.AutoMigrate(&models.User{}, &models.Category{}, &models.Products{}, &models.Cart{})

	if err != nil {
		panic(err.Error())
	}

	var router = gin.Default()

	router.POST("/product/:search", controller.SearchProduct)
	router.POST("/register", controller.Register)
	router.POST("/login", controller.Login)

	//* User Middleware
	router.Use(middlewares.AuthLogin())
	router.POST("/addCategory", controller.CreateCategory)
	router.GET("/categories", controller.GetCategories)
	router.GET("/products", controller.GetProducts)
	router.GET("/carts", controller.GetCarts)
	router.POST("/addToCart", controller.AddProductToCard)
	router.POST("/updateCart/:id", controller.UpdateCartsItem)
	router.POST("/deleteCart/:id", controller.DeleteCartsItem)

	//* Admin Middleware
	router.Use(middlewares.AuthJWTAdmin())
	router.POST("/addproduct", controller.CreateProduct)
	router.POST("/deleteProduct/:id", controller.DeleteProduct)
	router.POST("/updateProduct/:id", controller.UpdateProduct)
	router.POST("/uploadFile", controller.CreateFileCategory)

	router.Run()

}
