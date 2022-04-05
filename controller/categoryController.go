package controller

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"ck/configs"
	"ck/helper"
	"ck/models"

	"github.com/gin-gonic/gin"
)

func CreateFileCategory(g *gin.Context) {

	file, err := g.FormFile("uploadFile")

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	filepath := filepath.Base(file.Filename)

	g.SaveUploadedFile(file, "../helper/"+filepath)

	time.Sleep(time.Second * 1)

	result, err := helper.ReadFile("../helper/" + filepath)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	fmt.Println(len(result))

	var category models.Category

	for _, value := range result {

		categories := models.Category{
			Name:        value[0],
			Description: value[1],
		}
		configs.DB.Table("categories").Select("name").Where("name = ? ", categories.Name).Find(&category.Name)

		if categories.Name == category.Name {
			g.JSON(http.StatusBadRequest, gin.H{"error": "category already exists"})
			return
		}

		_, err = categories.CreateCategory()

		if err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		g.JSON(http.StatusOK, gin.H{"message": categories})
	}

}

func CreateCategory(g *gin.Context) {

	var category models.Category

	var categoryName string

	if err := g.ShouldBindJSON(&category); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	categories := models.Category{
		Name:        category.Name,
		Description: category.Description,
		Products:    category.Products,
	}

	if categories.Name == "" || categories.Description == "" {
		g.JSON(http.StatusBadRequest, gin.H{"error": "check name  description and products"})
		return
	}

	configs.DB.Table("categories").Select("name").Where("name = ? ", categories.Name).Find(&categoryName)
	if categoryName != "" {
		g.JSON(http.StatusBadRequest, gin.H{"error": "category already exists"})
		return
	}
	_, err := categories.CreateCategory()

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"message": categories})
}

func GetCategories(g *gin.Context) {

	var categories []string

	configs.DB.Table("categories").Select("name").Find(&categories)

	if len(categories) == 0 {
		g.JSON(http.StatusNotFound, gin.H{"message": "no categories found"})
		return
	}

	g.JSON(http.StatusOK, gin.H{"Categoies": categories})

}
