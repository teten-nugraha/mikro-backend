package domain

import "time"

type Kategori struct {
	Base
	Kategori string `gorm:"type:varchar(30);NOT NULL" json:"kategori" binding:"required"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"update_at"`
	DeletedAt *time.Time `sql:"index"`
}
