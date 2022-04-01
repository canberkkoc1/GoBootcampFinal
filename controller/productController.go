package controller

import (
	"ck/configs"
	"ck/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProduct(g *gin.Context) {

	var product models.Products

	if err := g.ShouldBindJSON(&product); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if product.Name == "" || product.Price <= 0 || product.Description == "" {
		g.JSON(http.StatusBadRequest, gin.H{"error": "check name  description and products"})
		return
	}

	_, err := product.CreateProduct()

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"message": product})
}

func GetProducts(g *gin.Context) {

	var products []models.Products

	configs.DB.Table("products").Find(&products)

	if len(products) == 0 {
		g.JSON(http.StatusNotFound, gin.H{"message": "no products found"})
		return
	}

	g.JSON(http.StatusOK, gin.H{"Products": products})
}
