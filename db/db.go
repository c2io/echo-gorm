package db

import (
	"echo-gorm/config"
	"echo-gorm/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Init() {
	c := config.GetConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		c.Database.User, c.Database.Password,
		c.Database.Host, c.Database.Port,
		c.Database.Name)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("DB Connection Error: %s", err.Error()))
	}

	db.AutoMigrate(&model.User{})
}

func DbManager() *gorm.DB {
	return db
}
