package route

import (
	"github.com/labstack/echo"
	"github.com/teten-nugraha/mikro-backend/api"
)

func AuthRoute(routes *echo.Echo, authAPI api.AuthAPI) {

	// auth route
	auth := routes.Group("/auth")
	{
		auth.GET("/generate-hash/:password", authAPI.GenerateHashPassword)
		auth.POST("/login", authAPI.CheckLogin)
		auth.POST("/signup", authAPI.SignUp)
	}

}