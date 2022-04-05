package models

import (
	"ck/configs"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"primary_key;auto_increment"  json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  string `json:"roles"`
	Carts    []Cart `gorm:"ForeignKey:UserID"  json:"carts"`
}

func (u *User) CreateUser() (*User, error) {
	var err error

	err = configs.DB.Create(&u).Error

	if err != nil {
		return &User{}, err
	}

	return u, nil

}
