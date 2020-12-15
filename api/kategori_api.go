package api

import (
	"github.com/labstack/echo"
	"github.com/teten-nugraha/mikro-backend/dto"
	"github.com/teten-nugraha/mikro-backend/service"
	"net/http"
)

type KategoriAPI struct {
	KategoriService service.KategoryService
}

func ProviderKategoryAPI(k service.KategoryService) KategoriAPI {
	return KategoriAPI{KategoriService: k}
}

func (k *KategoriAPI) FindAll(e echo.Context) error {
	kategoris := k.KategoriService.FindAll()

	if(len(kategoris) > 0) {
		return SuccessResponse(e, http.StatusOK, kategoris)
	}

	return SuccessResponse(e, http.StatusNoContent, kategoris)

}

func (k *KategoriAPI) FindById(e echo.Context) error {
	id := e.Param("id")

	kategori := k.KategoriService.FindById(id)

	return SuccessResponse(e, http.StatusOK, kategori)
}

func (k *KategoriAPI) FindByName(e echo.Context) error {
	nama := e.FormValue("nama")

	kategori := k.KategoriService.FindByNama(nama)

	return SuccessResponse(e, http.StatusOK, kategori)
}

func (k *KategoriAPI) SaveOrUpdate(c echo.Context) error {

	var newDto dto.KategoriDTO

	form := new(dto.KategoriDTO)
	c.Bind(form)
	if err := c.Validate(form); err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	kategori := c.FormValue("kategori")
	newDto.Kategori = kategori

	res, err := k.KategoriService.SaveOrUpdate(newDto)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return SuccessResponse(c, http.StatusOK, res)
}

func (k *KategoriAPI) DeleteKategori(c echo.Context) error{
	id := c.Param("id")

	err := k.KategoriService.DeleteKategori(id)
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return SuccessResponse(c, http.StatusOK, "Delete Success")
}
