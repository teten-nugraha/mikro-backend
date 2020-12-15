package route

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/teten-nugraha/mikro-backend/api"
	"github.com/teten-nugraha/mikro-backend/util"
)

func UserRoute(routes *echo.Echo, userAPI api.UserAPI) {

	// auth route
	user := routes.Group("/user")
	user.Use(middleware.JWT([]byte(util.SECRET)))
	user.GET("/authenticated", userAPI.FindById)
}
