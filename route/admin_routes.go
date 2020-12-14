package route

import (
	"github.com/labstack/echo"
	"github.com/teten-nugraha/mikro-backend/api"
)

func AdminRoute(routes *echo.Echo, kategoriAPI api.KategoriAPI) {

	admin := routes.Group("/admin")
	{
		// Kategori Route

		admin.POST("/kategori", kategoriAPI.SaveOrUpdate)
		admin.GET("/kategori/:id", kategoriAPI.FindById)
		admin.GET("/kategori/findByNama", kategoriAPI.FindByName)
	}

}
