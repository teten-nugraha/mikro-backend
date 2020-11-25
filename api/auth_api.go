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

	return SuccessResponse(c, http.StatusOK, hash)

}

func (p *AuthAPI) CheckLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	res, err := p.UserService.CheckLogin(username, password)
	if err != nil {

		return ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	if !res {
		return ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
	}

	t, err := p.UserService.GenerateToken(username)

	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return SuccessResponse(c, http.StatusOK, map[string]string{
		"token":t,
	})

}
