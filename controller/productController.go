package controller

import (
	"ck/configs"
	"ck/helper"
	"ck/models"
	"fmt"
	"net/http"
	"strconv"

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
	/*
		var products []models.Products

		configs.DB.Table("products").Find(&products)

		if len(products) == 0 {
			g.JSON(http.StatusNotFound, gin.H{"message": "no products found"})
			return
		}
		g.JSON(http.StatusOK, gin.H{"Products": products})
	*/
	pageIndex, pageSize := GetPaginationParameterFromRequest(g)

	items, count, _ := helper.GetProductsPage(pageIndex, pageSize)

	paginationResult := NewFromRequest(g, items, count)

	paginationResult.Data = items

	g.JSON(http.StatusOK, paginationResult)

}

func SearchProduct(g *gin.Context) {

	var products []models.Products

	var search string

	search = g.Param("search")

	configs.DB.Table("products").Where("name LIKE ?", "%"+search+"%").Find(&products)

	if len(products) == 0 {
		g.JSON(http.StatusNotFound, gin.H{"message": "no products found"})
		return
	}

	g.JSON(http.StatusOK, gin.H{"Products": products})
}

func DeleteProduct(g *gin.Context) {

	var product models.Products

	productID, err := strconv.Atoi(g.Param("id"))

	fmt.Println(productID)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if productID == 0 {
		g.JSON(http.StatusBadRequest, gin.H{"error": "check id"})
		return
	}

	configs.DB.Table("products").Where("id = ?", productID).Delete(&product)

	g.JSON(http.StatusOK, gin.H{"message": "product deleted"})
}

func UpdateProduct(g *gin.Context) {

	var product models.Products

	productID, err := strconv.Atoi(g.Param("id"))

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if productID == 0 {
		g.JSON(http.StatusBadRequest, gin.H{"error": "check id"})
		return
	}

	configs.DB.Table("products").Select("deleted_at").Where("id = ?", productID).Row().Scan(&product.DeletedAt)

	if !product.DeletedAt.Time.IsZero() {
		g.JSON(http.StatusBadRequest, gin.H{"error": "product not found"})
		return
	}

	if err := g.ShouldBindJSON(&product); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	configs.DB.Table("products").Where("id = ?", productID).Updates(&product)

	g.JSON(http.StatusOK, gin.H{"message": "product updated"})
}
