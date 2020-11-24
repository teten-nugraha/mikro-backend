package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/teten-nugraha/mikro-backend/domain"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/mikro_backend")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&domain.User{})

	return db
}
