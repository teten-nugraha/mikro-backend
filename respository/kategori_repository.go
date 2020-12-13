package respository

import (
	"github.com/jinzhu/gorm"
	"github.com/teten-nugraha/mikro-backend/domain"
)

type KategoriRepositoryContract interface {
	FindById(id string) domain.Kategori
	FindByNama(nama string) domain.Kategori
}

type KategoriRepository struct {
	DB *gorm.DB
}

func ProviderKategoriRepository(DB *gorm.DB) KategoriRepository {
	return KategoriRepository{DB: DB}
}

func (k *KategoriRepository) FindById(id string) domain.Kategori {
	var kategori domain.Kategori

	k.DB.First(&kategori, id)

	return kategori
}

func (k *KategoriRepository) FindByNama(nama string) domain.Kategori {
	var kategori domain.Kategori

	k.DB.Where("kategori = ?", nama).Find(&kategori)

	return kategori
}
