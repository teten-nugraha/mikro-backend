package mapper

import (
	"github.com/teten-nugraha/mikro-backend/domain"
	"github.com/teten-nugraha/mikro-backend/dto"
)

func ToKategoriDTO(k domain.Kategori) dto.KategoriDTO {
	return dto.KategoriDTO{
		ID: k.ID,
		Kategori: k.Kategori,
	}
}

func ToKategoriListDTO(kategoris []domain.Kategori) []dto.KategoriDTO {
	kategoridtos := make([]dto.KategoriDTO, len(kategoris))

	for i, itm := range kategoris {
		kategoridtos[i] = ToKategoriDTO(itm)
	}

	return kategoridtos
}

func ToKategoriEntity(dto dto.KategoriDTO) domain.Kategori {
	return domain.Kategori{
		Kategori: dto.Kategori,
	}
}
