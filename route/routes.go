package route

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	db2 "github.com/teten-nugraha/mikro-backend/db"
	"github.com/teten-nugraha/mikro-backend/injection"
	"net/http"
)



func Init() *echo.Echo {

	db := db2.InitDB()
	//defer db.Close()

	authAPI := injection.InitAuthAPI(db)
	userAPI := injection.InitUserAPI(db)

	routes := echo.New()

	routes.Use(middleware.Logger())

	routes.GET("/", func(e echo.Context) error {
		return e.String(http.StatusOK, "Hello from mikro backend API")
	})

	// User Route
	routes.GET("/users/:id", userAPI.FindById)

	// login route
	routes.GET("/generate-hash/:password", authAPI.GenerateHashPassword)
	routes.GET("/login", authAPI.CheckLogin)

	return routes
}
