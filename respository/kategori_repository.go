package respository

import (
	"github.com/jinzhu/gorm"
	"github.com/teten-nugraha/mikro-backend/domain"
	"github.com/teten-nugraha/mikro-backend/dto"
	"github.com/teten-nugraha/mikro-backend/mapper"
)

type KategoriRepositoryContract interface {
	FindById(id string) domain.Kategori
	FindByNama(nama string) domain.Kategori
	SaveOrUpdate(dto dto.KategoriDTO) domain.Kategori
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

func (k *KategoriRepository) SaveOrUpdate(dto dto.KategoriDTO) (domain.Kategori, error) {

	kategori := mapper.ToKategoriEntity(dto)

	if err := k.DB.Create(&kategori).Error; err != nil {
		return kategori, err
	}

	return kategori, nil
}
