package configs

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewMySQLDB(configSting string) *gorm.DB {

	var err error
	DB, err = gorm.Open(mysql.Open(configSting), &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		panic(err.Error())
	}

	return DB

}
