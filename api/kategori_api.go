package api

import (
	"github.com/labstack/echo"
	"github.com/teten-nugraha/mikro-backend/service"
	"net/http"
)

type KategoriAPI struct {
	KategoriService service.KategoryService
}

func ProviderKategoryAPI(k service.KategoryService) KategoriAPI {
	return KategoriAPI{KategoriService: k}
}

func (k *KategoriAPI) FindById(e echo.Context) error {
	id := e.Param("id")

	kategori := k.KategoriService.FindById(id)

	return SuccessResponse(e, http.StatusOK, kategori)
}

func (k *KategoriAPI) SaveOrUpdate(c echo.Context) error {

}
