package api

import (
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/teten-nugraha/mikro-backend/service"
)

type UserAPI struct {
	UserService service.UserService
}

func ProviderUserAPI(u service.UserService) UserAPI {
	return UserAPI{UserService: u}
}

func (u *UserAPI) FindById(c echo.Context) error {
	id := c.Param("id")

	var idInt, _ = strconv.Atoi(id)

	user := u.UserService.FindById(uint(idInt))

	return SuccessResponse(c, http.StatusOK, user)
}

func (u *UserAPI) Authenticated(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)

	return SuccessResponse(c, http.StatusOK, user.Claims)
}
