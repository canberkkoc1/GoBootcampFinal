package models

import (
	"ck/configs"

	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       uint    `json:"stock"`
}

func (p *Products) CreateProduct() (*Products, error) {

	err := configs.DB.Create(p).Error

	if err != nil {
		return nil, err
	}

	return p, nil

}
