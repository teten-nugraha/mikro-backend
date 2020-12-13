package service

import (
	"github.com/teten-nugraha/mikro-backend/domain"
)

type KategoriServiceContract interface {
	FindById(id string) domain.Kategori
	FindByNama(nama string) domain.Kategori
}

type KategoriService struct {
	KategoriRepository KategoriRepository
}
