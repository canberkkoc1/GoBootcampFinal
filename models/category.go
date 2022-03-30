package models

import (
	"ck/configs"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Products    []Products `json:"products" gorm:"type:text" gorm:"many2many:product_categories;"`
}

func (c *Category) CreateCategory() (*Category, error) {

	err := configs.DB.Create(c).Error

	if err != nil {
		return nil, err
	}

	return c, nil
}
