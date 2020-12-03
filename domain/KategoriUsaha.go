package domain

import "time"

type Kategori struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Kategori  string `gorm:"type:varchar(30);NOT NULL" json:"kategori" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
