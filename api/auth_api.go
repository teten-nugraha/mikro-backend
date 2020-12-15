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
	form := new(dto.SignupDTO)
	c.Bind(form)
	if err := c.Validate(form); err != nil {
		return err
	}

	signupDto := dto.SignupDTO{
		Username: form.Username,
		Password: form.Password,
		Fullname: form.Fullname,
		Email:    form.Email,
		Phone:    form.Phone,
	}

	userDto, err := p.UserService.SignUp(signupDto)
	if err != nil {
		return ErrorResponse(c, http.StatusOK, err.Error())
	}

	return SuccessResponse(c, http.StatusOK, userDto)
}

func (p *AuthAPI) CheckLogin(c echo.Context) error {
	form := new(dto.LoginDTO)
	c.Bind(form)
	if err := c.Validate(form); err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	username := form.Username
	password := form.Password

	user, err := p.UserService.CheckLogin(username, password)
	if err != nil {
		return ErrorResponse(c, http.StatusOK, err.Error())
	}

	t, _ := p.UserService.GenerateToken(user)

	return SuccessResponse(c, http.StatusOK, map[string]string{
		"token": t,
	})

}
