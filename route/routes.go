package route

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	db2 "github.com/teten-nugraha/mikro-backend/db"
	"github.com/teten-nugraha/mikro-backend/injection"
)



func Init(arg string) *echo.Echo {

	db := db2.InitDB(arg)
	//defer db.Close()


	userAPI := injection.InitUserAPI(db)
	authAPI := injection.InitAuthAPI(db)

	routes := echo.New()

	// set logger
	routes.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, time:${time_unix}, uri=${uri}, status=${status}, error=${error}, latency_human=${latency}, bytes_in=${bytes_in}, bytes_out=${bytes_out} \n",
	}))

	// Gzip Compression
	routes.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	// Allow CORS
	routes.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// User Route
	routes.GET("/users/:id", userAPI.FindById)

	// auth route
	auth := routes.Group("/auth")
	{
		auth.GET("/generate-hash/:password", authAPI.GenerateHashPassword)
		auth.GET("/login", authAPI.CheckLogin)
		auth.POST("/signup", authAPI.SignUp)
	}

	return routes
}
