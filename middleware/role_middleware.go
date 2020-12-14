package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/teten-nugraha/mikro-backend/util"
	"strings"
)

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		role := claims["role"].(string)
		if strings.Compare(util.MIKRO_ADMIN, role) != 0 {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}

