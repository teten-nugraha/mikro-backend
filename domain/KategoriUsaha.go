package domain

type Kategori struct {
	BaseUUID

	Kategori string `gorm:"type:varchar(30);NOT NULL" json:"kategori" binding:"required"`

	BaseTime
}
