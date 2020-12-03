package domain

import "time"

type Merchant struct {
	ID         uint   `gorm:"primary_key" json:"id"`
	UserID     uint   `gorm:"type:int;column:user_id;NOT NULL"`
	KategoriID uint   `gorm:"type:int;column:kategori_id;NOT NULL"`
	Nama       string `gorm:"type:varchar(30);NOT NULL" json:"nama" binding:"required"`
	Usia       uint   `gorm:"type:int;NOT NULL" json:"born" binding:"required"`
	Alamat     uint   `gorm:"type:varchar(255);NOT NULL" json:"born" binding:"required"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
