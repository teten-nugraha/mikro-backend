package domain

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
	Fullname string
	Role string
	IsActive bool
	Email string
	Phone string
}
