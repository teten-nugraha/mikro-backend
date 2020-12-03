package domain

type User struct {
	BaseUUID

	Username string `gorm:"type:varchar(255);NOT NULL" json:"username" binding:"required"`
	Password string `gorm:"type:varchar(255)" json:"password"`
	Fullname string `gorm:"type:varchar(255)" json:"fullname"`
	Role     string `gorm:"type:varchar(20)" json:"role"`
	IsActive bool   `gorm:"type:bool" json:"is_active"`
	Email    string `gorm:"type:varchar(255)" json:"email"`
	Phone    string `gorm:"type:varchar(255)" json:"phone"`

	BaseTime

}
