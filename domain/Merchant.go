package domain

import "time"

type Merchant struct {
	ID                  uint   `gorm:"primary_key" json:"id"`
	UserID              uint   `gorm:"type:int;column:user_id;NOT NULL"`
	KategoriID          uint   `gorm:"type:int;column:kategori_id;NOT NULL"`
	Name                string `gorm:"type:varchar(30);NOT NULL" json:"name"`
	BusinessStartedDate uint   `gorm:"type:date;NOT NULL" json:"business_started_date"`
	Address             uint   `gorm:"type:varchar(255);NULL" json:"address"`
	BusinessAddress     uint   `gorm:"type:varchar(255);NULL" json:"business_address"`
	Omzet               uint   `gorm:"type:double;NULL" json:"omzet"`
	Profit              uint   `gorm:"type:double;NULL" json:"profit"`
	HasLegal            uint   `gorm:"type:tinyint(1);NULL" json:"has_legal"`
	BusinessProblem     uint   `gorm:"type:text;NULL" json:"business_problem"`
	EmployeeCount       uint   `gorm:"type:int(11);NULL" json:"employee_count"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
}
