package domain

import "time"

type User struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Username  string `gorm:"type:varchar(255);NOT NULL" json:"username" binding:"required"`
	Password  string `gorm:"type:varchar(255)" json:"password"`
	Fullname  string `gorm:"type:varchar(255)" json:"fullname"`
	Role      string `gorm:"type:varchar(20)" json:"role"`
	IsActive  bool   `gorm:"type:bool" json:"is_active"`
	Email     string `gorm:"type:varchar(255)" json:"email"`
	Phone     string `gorm:"type:varchar(255)" json:"phone"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
