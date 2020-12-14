package mapper

import (
	"github.com/teten-nugraha/mikro-backend/domain"
	"github.com/teten-nugraha/mikro-backend/dto"
)

func ToKategoriDTO(k domain.Kategori) dto.KategoriDTO {
	return dto.KategoriDTO{
		Kategori: k.Kategori,
	}
}

func ToKategoriEntity(dto dto.KategoriDTO) domain.Kategori {
	return domain.Kategori{
		Kategori: dto.Kategori,
	}
}
