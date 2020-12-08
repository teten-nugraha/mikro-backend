package domain

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

// Base contains common columns for all tables.
type BaseUUID struct {
	ID        uuid.UUID  `gorm:"primaryKey"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *BaseUUID) BeforeCreate(scope *gorm.Scope) error {
	uuidVal := uuid.NewV4()
	return scope.SetColumn("ID", uuidVal)
}
