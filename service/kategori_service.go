package service

import (
	"github.com/teten-nugraha/mikro-backend/domain"
	"github.com/teten-nugraha/mikro-backend/dto"
	"github.com/teten-nugraha/mikro-backend/mapper"
	"github.com/teten-nugraha/mikro-backend/respository"
)

type KategoriServiceContract interface {
	FindById(id string) domain.Kategori
	FindByNama(nama string) domain.Kategori
	SaveOrUpdate(dto dto.KategoriDTO) (dto.KategoriDTO, error)
}

type KategoryService struct {
	KategoriRepository respository.KategoriRepository
}

func ProviderKategoriService(k respository.KategoriRepository) KategoryService {
	return KategoryService{
		KategoriRepository: k,
	}
}

func (k *KategoryService) FindById(id string) dto.KategoriDTO {
	return mapper.ToKategoriDTO(k.KategoriRepository.FindById(id))
}

func (k *KategoryService) FindByNama(nama string) dto.KategoriDTO {
	return mapper.ToKategoriDTO(k.KategoriRepository.FindByNama(nama))
}

func (k *KategoryService) SaveOrUpdate(dto dto.KategoriDTO) (dto.KategoriDTO, error) {

	kategori, _ := k.KategoriRepository.SaveOrUpdate(dto)

	return mapper.ToKategoriDTO(kategori), nil

}
