package helper

import (
	"ck/configs"
	"ck/models"
)

func GetCategory(pageIndex, pageSize int) ([]models.Category, int, error) {

	var categories []models.Category

	var count int64

	err := configs.DB.Offset(pageIndex - 1).Limit(pageSize).Find(&categories).Count(&count).Error

	if err != nil {
		return nil, int(count), err
	}

	return categories, int(count), nil

}
