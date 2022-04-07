package models

import (
	"ck/configs"

	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	UserID    uint
	ProductID uint
	Quantity  uint
	Price     uint `json:"price"`
	Status    string
}

func (o *Orders) CreateOrder() (*Orders, error) {

	err := configs.DB.Create(o).Error

	if err != nil {
		return nil, err
	}

	return o, nil

}
