package models

import (
	"ck/configs"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID    uint   `gorm:"primary_key;auto_increment" json:"user_id"`
	ProductID uint   `json:"product_id"`
	Quantity  uint   `json:"quantity"`
	Price     uint   `json:"price"`
	Status    string `json:"status"`
}

func (c *Cart) CreateCart() (*Cart, error) {

	err := configs.DB.Create(c).Error

	if err != nil {
		return nil, err
	}

	return c, nil

}
