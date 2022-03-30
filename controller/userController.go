package controller

import (
	"ck/configs"
	"ck/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(g *gin.Context) {

	var input models.User

	var loginUser []models.User

	if err := g.ShouldBindJSON(&input); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userAdmin := models.User{
		Email:    input.Email,
		Password: input.Password,
		IsAdmin:  input.IsAdmin,
	}

	if userAdmin.Email == "" || userAdmin.Password == "" || userAdmin.IsAdmin == "" {
		g.JSON(http.StatusBadRequest, gin.H{"error": "check email  password and roles"})
		return
	}

	if userAdmin.IsAdmin != "Y" && userAdmin.IsAdmin != "N" {
		g.JSON(http.StatusBadRequest, gin.H{"error": "roles must be 'N' or 'Y'"})
		return
	}

	validToken, err := GenerateJWT(input.Email, input.IsAdmin)

	if err != nil {
		g.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	configs.DB.Table("users").Where("email = ? ", userAdmin.Email).Find(&loginUser)

	if len(loginUser) != 0 {
		g.JSON(http.StatusBadRequest, gin.H{"error": "email already exists"})
		return
	}

	_, errCreate := userAdmin.CreateUser()

	if errCreate != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusCreated, gin.H{"data": validToken})
}

func Login(g *gin.Context) {
	var loginUser []models.User
	var user models.User
	var checkAdmin string

	if err := g.ShouldBindJSON(&user); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	email := user.Email
	pass := user.Password

	fmt.Println(email)
	fmt.Println(pass)

	errDb := configs.DB.Table("users").Where("email = ? AND password = ?", email, pass).Find(&loginUser)
	configs.DB.Table("users").Select("is_admin").Where("email = ? AND password = ?", email, pass).Find(&checkAdmin)

	fmt.Println(loginUser)
	fmt.Println(checkAdmin)

	validToken, _ := GenerateJWT(email, checkAdmin)

	if len(loginUser) == 0 {
		g.JSON(http.StatusBadRequest, gin.H{"error": "email or password wrong"})
		return
	}
	if errDb.Error != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "email or password wrong"})
		return
	} else {

		g.JSON(http.StatusOK, gin.H{"data": validToken})
	}

}
