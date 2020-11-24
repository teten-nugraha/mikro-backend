package api

import (
	"github.com/labstack/echo"
	"github.com/teten-nugraha/mikro-backend/helpers"
	"github.com/teten-nugraha/mikro-backend/service"
	"net/http"
)

type AuthAPI struct {
	UserService service.UserService
}

func ProvideAuthAPI(p service.UserService) AuthAPI {
	return AuthAPI{
		UserService: p,
	}
}

func (p *AuthAPI)GenerateHashPassword(c echo.Context) error {
	password := c.Param("password")

	hash, _ := helpers.HashPassword(password)

	return c.JSON(http.StatusOK, hash)

}
