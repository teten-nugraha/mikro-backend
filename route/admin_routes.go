package route

import (
	"github.com/labstack/echo"
	"github.com/teten-nugraha/mikro-backend/api"
	mdl "github.com/teten-nugraha/mikro-backend/middleware"
)

func AdminRoute(routes *echo.Echo, kategoriAPI api.KategoriAPI) {

	admin := routes.Group("/admin",  mdl.IsAuthenticated, mdl.IsAdmin)
	{
		// Kategori Route

		admin.GET("/kategoris", kategoriAPI.FindAll)
		admin.POST("/kategori", kategoriAPI.SaveOrUpdate)
		admin.GET("/kategori/:id", kategoriAPI.FindById)
		admin.GET("/kategori/findByNama", kategoriAPI.FindByName)
		admin.DELETE("/kategori/:id", kategoriAPI.DeleteKategori)
	}

}
