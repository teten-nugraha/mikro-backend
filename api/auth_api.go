package api

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/teten-nugraha/mikro-backend/dto"
	"github.com/teten-nugraha/mikro-backend/helpers"
	"github.com/teten-nugraha/mikro-backend/service"
)

type AuthAPI struct {
	UserService service.UserService
}

func ProvideAuthAPI(p service.UserService) AuthAPI {
	return AuthAPI{
		UserService: p,
	}
}

func (p *AuthAPI) GenerateHashPassword(c echo.Context) error {
	password := c.Param("password")

	hash, _ := helpers.HashPassword(password)

	return SuccessResponse(c, http.StatusOK, hash)

}

func (p *AuthAPI) SignUp(c echo.Context) error {

	username := c.FormValue("username")
	password := c.FormValue("password")
	fullname := c.FormValue("fullname")
	email := c.FormValue("email")
	phone := c.FormValue("phone")

	signupDto := dto.SignupDTO{
		Username: username,
		Password: password,
		Fullname: fullname,
		Email:    email,
		Phone:    phone,
	}

	userDto := p.UserService.SignUp(signupDto)

	return SuccessResponse(c, http.StatusOK, userDto)
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
		"token": t,
	})

}
