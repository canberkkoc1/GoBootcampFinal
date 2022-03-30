package models

import (
	"ck/configs"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  string `json:"roles"`
}

func (u *User) CreateUser() (*User, error) {
	var err error

	err = configs.DB.Create(&u).Error

	if err != nil {
		return &User{}, err
	}

	return u, nil

}
