package controller

import (
	"ck/configs"
	"ck/helper"
	"ck/models"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CompleteOrder(g *gin.Context) {

	var user_id uint

	var product_id uint

	var quantity uint

	var price uint

	var carts []models.Cart

	userEmail := models.GetEmail(g)

	configs.DB.Table("users").Select("id").Where("email = ? ", userEmail).Find(&user_id)

	configs.DB.Table("carts").Select("product_id").Where("user_id = ? ", user_id).Find(&product_id)

	configs.DB.Table("carts").Select("quantity").Where("user_id = ? ", user_id).Find(&quantity)

	configs.DB.Table("carts").Select("price").Where("user_id = ?", user_id).Find(&price)

	result, err := configs.DB.Table("carts").Where("user_id = ?", user_id).Rows()

	defer result.Close()

	if err != nil {
		fmt.Println(err)

	}

	for result.Next() {
		var cart models.Cart
		configs.DB.ScanRows(result, &cart)
		carts = append(carts, cart)

	}

	fmt.Println(carts)

	for _, cart := range carts {

		if cart.DeletedAt.Time.IsZero() {

			orders := models.Orders{
				UserID:    user_id,
				ProductID: cart.ProductID,
				Quantity:  cart.Quantity,
				Price:     cart.Price,
				Status:    "completed",
			}

			configs.DB.Create(&orders)

			configs.DB.Table("carts").Where("user_id = ?", cart.UserID).Delete(&models.Cart{})
		}

	}

	g.JSON(200, gin.H{"message": "order completed"})

}

func ListOrder(g *gin.Context) {

	var orders []models.Orders

	var user_id uint

	userEmail := models.GetEmail(g)

	configs.DB.Table("users").Select("id").Where("email = ? ", userEmail).Find(&user_id)

	configs.DB.Table("orders").Select("*").Where("user_id = ?", user_id).Find(&orders)

	g.JSON(200, gin.H{"orders": orders})

}

func CancelOrder(g *gin.Context) {

	var order models.Orders

	var user_id uint

	var date time.Time

	var order_user_id []int

	userEmail := models.GetEmail(g)

	order_id, err := strconv.Atoi(g.Param("id"))

	if err != nil {

		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	configs.DB.Table("users").Select("id").Where("email = ? ", userEmail).Find(&user_id)

	configs.DB.Table("orders").Select("*").Where("user_id = ? AND id = ?", user_id, order_id).Find(&order)

	configs.DB.Table("orders").Select("created_at").Where("id= ?", order_id).Find(&date)

	result, _ := configs.DB.Table("orders").Select("user_id").Where("id = ?", order_id).Rows()

	defer result.Close()

	for result.Next() {

		configs.DB.ScanRows(result, &order_user_id)

	}

	isUserExist := helper.CheckSlice(order_user_id, int(user_id))

	if !isUserExist {

		g.JSON(http.StatusBadRequest, gin.H{"error": "you are not allowed to cancel this order"})

		return

	}

	if !order.DeletedAt.Time.IsZero() {

		g.JSON(http.StatusBadRequest, gin.H{"error": "order already Deleted"})

		return

	}

	if order.ID == uint(order_id) {

		t1 := time.Now()

		if t1.Sub(date).Hours() > 336 {

			g.JSON(200, gin.H{"message": "order can not be canceled after 14 hours"})
			return

		} else {

			configs.DB.Table("orders").Where("id = ?", order_id).Update("status", "canceled")

			configs.DB.Table("orders").Where("id = ?", order_id).Delete(&models.Orders{})

			g.JSON(200, gin.H{"message": "order canceled"})

			return

		}

	} else {

		g.JSON(http.StatusBadRequest, gin.H{"error": "order not found"})

		return

	}

}
