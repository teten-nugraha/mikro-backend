package middleware

import (
	"github.com/labstack/echo/middleware"
	"github.com/teten-nugraha/mikro-backend/util"
)

var IsAuthenticated = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte(util.SECRET),
})