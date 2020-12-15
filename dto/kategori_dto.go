package dto

import uuid "github.com/satori/go.uuid"

type KategoriDTO struct {
	ID       uuid.UUID `json:"id,string,omitempty"`
	Kategori string    `json:"kategori" validate:"required"`
}
