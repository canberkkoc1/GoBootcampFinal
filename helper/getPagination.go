package helper

import (
	"ck/configs"
	"ck/models"
)

func GetCategory(pageIndex, pageSize int) ([]models.Category, int, error) {

	var categories []models.Category

	var count int64

	err := configs.DB.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&categories).Count(&count).Error

	if err != nil {
		return nil, int(count), err
	}

	return categories, int(count), nil

}

func GetProductsPage(pageIndex, pageSize int) ([]models.Products, int, error) {

	var products []models.Products

	var count int64

	err := configs.DB.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&products).Count(&count).Error

	if err != nil {
		return nil, int(count), err
	}

	return products, int(count), nil

}

func GetCartsPagination(pageIndex, pageSize int) ([]models.Cart, int, error) {

	var carts []models.Cart

	var count int64

	err := configs.DB.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&carts).Count(&count).Error

	if err != nil {
		return nil, int(count), err
	}

	return carts, int(count), nil

}

func GetOrderPagination(pageIndex, pageSize int) ([]models.Orders, int, error) {

	var orders []models.Orders

	var count int64

	err := configs.DB.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&orders).Count(&count).Error

	if err != nil {
		return nil, int(count), err
	}

	return orders, int(count), nil

}
