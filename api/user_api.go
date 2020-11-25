package api

import (
	"github.com/labstack/echo"
	"github.com/teten-nugraha/mikro-backend/service"
	"net/http"
	"strconv"
)

type UserAPI struct {
	UserService service.UserService
}

func ProviderUserAPI(u service.UserService) UserAPI {
	return UserAPI{UserService: u}
}

func (u *UserAPI) FindById(c echo.Context) error {
	id := c.Param("id")

	var idInt,_ = strconv.Atoi(id)

	user := u.UserService.FindById(uint(idInt))

	return SuccessResponse(c, http.StatusOK, user)

}