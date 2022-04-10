package main

import (
	"ck/configs"
	"ck/controller"
	"ck/middlewares"
	"ck/models"

	"github.com/gin-gonic/gin"
)

func main() {

	// init database

	db := configs.NewMySQLDB("<rootname>:<password>@tcp(127.0.0.1:<port>)/<dbName>?parseTime=True&loc=Local")

	err := db.AutoMigrate(&models.User{}, &models.Category{}, &models.Products{}, &models.Cart{}, &models.Orders{})

	if err != nil {
		panic(err.Error())
	}

	// init router

	var router = gin.Default()

	// init middlewares and routes

	router.POST("/product/:search", controller.SearchProduct)
	router.POST("/register", controller.Register)
	router.POST("/login", controller.Login)

	//* User Middleware
	router.Use(middlewares.AuthLogin())
	router.POST("/addCategory", controller.CreateCategory)
	router.GET("/categories", controller.GetCategories)
	router.GET("/products", controller.GetProducts)
	router.GET("/carts", controller.ListCart)
	router.POST("/addToCart", controller.AddProductToCard)
	router.PATCH("/updateCart/:id", controller.UpdateCartsItem)
	router.DELETE("/deleteCart/:id", controller.DeleteCartsItem)
	router.POST("/compeleteOrder", controller.CompleteOrder)
	router.GET("/listOrder", controller.ListOrder)
	router.POST("/cancelOrder/:id", controller.CancelOrder)

	//* Admin Middleware

	router.Use(middlewares.AuthJWTAdmin())
	router.POST("/addproduct", controller.CreateProduct)
	router.DELETE("/deleteProduct/:id", controller.DeleteProduct)
	router.PATCH("/updateProduct/:id", controller.UpdateProduct)
	router.POST("/uploadFile", controller.CreateFileCategory)

	router.Run()

}
