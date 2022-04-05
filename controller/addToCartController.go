package controller

import (
	"ck/configs"
	"ck/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddProductToCard(g *gin.Context) {

	var cartProduct models.Cart
	var user_id uint
	var product_id uint
	var product_user_id uint
	var stock uint
	var quantity uint
	var cart_id uint

	if err := g.ShouldBindJSON(&cartProduct); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	configs.DB.Table("products").Select("stock").Where("id = ? ", cartProduct.ProductID).Find(&stock)

	configs.DB.Table("carts").Select("quantity").Where("product_id = ? ", cartProduct.ProductID).Find(&quantity)

	userEmail := models.GetEmail(g)

	configs.DB.Table("users").Select("id").Where("email = ? ", userEmail).Find(&user_id)

	configs.DB.Table("products").Select("id").Where("id = ?", cartProduct.ProductID).Find(&product_id)

	configs.DB.Table("carts").Select("user_id").Where("product_id = ? ", product_id).Find(&product_user_id)

	configs.DB.Table("carts").Select("product_İd").Where("product_id = ? ", product_id).Find(&cart_id)

	fmt.Println(user_id)

	if product_id == cart_id && product_user_id == user_id {
		var quantity uint

		configs.DB.Table("carts").Select("quantity").Where("product_id = ? ", cartProduct.ProductID).Find(&quantity)

		configs.DB.Table("carts").Where("product_id = ? ", cartProduct.ProductID).Update("quantity", cartProduct.Quantity+quantity)

		configs.DB.Table("products").Where("id = ? ", cartProduct.ProductID).Update("stock", stock-cartProduct.Quantity)

		g.JSON(http.StatusOK, gin.H{"message": "product updated"})
		return
	} else {

		cartProduct.UserID = user_id
		_, err := cartProduct.CreateCart()

		if err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	if user_id == 0 || cartProduct.ProductID != product_id || stock <= 0 {
		g.JSON(http.StatusBadRequest, gin.H{"error": "check productID and stock"})
		return
	}

	if quantity > stock {
		g.JSON(http.StatusBadRequest, gin.H{"error": "out of stock"})
		return
	}

	if cartProduct.Quantity > stock {
		g.JSON(http.StatusBadRequest, gin.H{"error": "check quantity"})
		return
	}

	g.JSON(http.StatusOK, gin.H{"message": "product added to card"})

}

//! get all data from product

func GetCarts(g *gin.Context) {

	var products []models.Products

	var product_id []int

	var user_id uint

	userEmail := models.GetEmail(g)

	configs.DB.Table("users").Select("id").Where("email = ? ", userEmail).Find(&user_id)

	rows, err := configs.DB.Table("carts").Select("product_id").Where("user_id = ? ", user_id).Rows()

	defer rows.Close()

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for rows.Next() {
		configs.DB.ScanRows(rows, &product_id)
	}

	fmt.Println(product_id)

	if user_id == 0 {
		g.JSON(http.StatusBadRequest, gin.H{"error": "check productID and stock"})
		return
	}

	for _, id := range product_id {
		fmt.Println(id)
		configs.DB.Table("products").Where("id = ? ", id).Find(&products)

	}

	g.JSON(http.StatusOK, gin.H{"data": products})

}

func DeleteCartsItem(g *gin.Context) {

	var carts models.Cart
	var user_id uint
	var cart_user_id []int

	CartID, err := strconv.Atoi(g.Param("id"))

	userEmail := models.GetEmail(g)

	configs.DB.Table("users").Select("id").Where("email = ? ", userEmail).Find(&user_id)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, _ := configs.DB.Table("carts").Select("id").Where("user_id = ? ", user_id).Rows()

	defer result.Close()

	for result.Next() {
		configs.DB.ScanRows(result, &cart_user_id)
	}

	if user_id == 0 || CartID == 0 {
		g.JSON(http.StatusBadRequest, gin.H{"error": "check productID and stock"})
		return
	}

	for _, id := range cart_user_id {

		configs.DB.Table("carts").Select("deleted_at").Where("id = ? ", id).Row().Scan(&carts.DeletedAt)

		if !carts.DeletedAt.Time.IsZero() {
			g.JSON(http.StatusBadRequest, gin.H{"error": "Item not found"})
			return
		}
		if id == CartID {
			configs.DB.Table("carts").Where("id = ? ", CartID).Delete(&carts)
			g.JSON(http.StatusOK, gin.H{"message": "product deleted"})
			return
		}
	}

	g.JSON(http.StatusBadRequest, gin.H{"error": "you can't delete this product"})

}

//! not tested

func UpdateCartsItem(g *gin.Context) {

	var carts models.Cart
	var user_id uint
	var cart_user_id []int
	var product_id uint

	var stock uint

	CartID, err := strconv.Atoi(g.Param("id"))

	userEmail := models.GetEmail(g)

	configs.DB.Table("users").Select("id").Where("email = ? ", userEmail).Find(&user_id)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := g.ShouldBindJSON(&carts); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, _ := configs.DB.Table("carts").Select("id").Where("user_id = ? ", user_id).Rows()

	defer result.Close()

	for result.Next() {
		configs.DB.ScanRows(result, &cart_user_id)
	}

	configs.DB.Table("carts").Select("product_id").Where("id = ? ", CartID).Find(&product_id)
	configs.DB.Table("products").Select("stock").Where("id = ? ", product_id).Find(&stock)

	fmt.Println(stock)
	fmt.Println(product_id)

	for _, id := range cart_user_id {

		if id == CartID {
			if carts.Quantity <= 0 {
				g.JSON(http.StatusBadRequest, gin.H{"error": "check quantity"})
				return
			}
			if carts.Quantity > stock {
				g.JSON(http.StatusBadRequest, gin.H{"error": "out of stock"})
				return
			}
			configs.DB.Table("carts").Where("id = ? AND user_id=? ", CartID, user_id).Update("quantity", carts.Quantity)
			configs.DB.Table("products").Where("id = ? ", product_id).Update("stock", stock-carts.Quantity)

			g.JSON(http.StatusOK, gin.H{"message": "Item updated"})
			return
		}
	}

	g.JSON(http.StatusBadRequest, gin.H{"error": "you can't update this Item"})

}
