package domain

import "time"

type User struct {
	Base
	Username string `gorm:"type:varchar(255);NOT NULL" json:"username" binding:"required"`
	Password string `gorm:"type:varchar(255)" json:"password"`
	Fullname string `gorm:"type:varchar(255)" json:"fullname"`
	Role     string `gorm:"type:varchar(20)" json:"role"`
	IsActive bool   `gorm:"type:bool" json:"is_active"`
	Email    string `gorm:"type:varchar(255)" json:"email"`
	Phone    string `gorm:"type:varchar(255)" json:"phone"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"update_at"`
	DeletedAt *time.Time `sql:"index"`
}
