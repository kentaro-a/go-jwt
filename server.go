package main

import (
	"fmt"
	"github.com/jwt/jwt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func main() {
	fmt.Println("start")
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/login", LoginHandler)
	require_auth_group := e.Group("", AuthenticateMiddleware())
	require_auth_group.GET("/mypage", MyPageHandler)
	e.Start(":9999")
}

func AuthenticateMiddleware() echo.MiddlewareFunc {
	return func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Authentication
			if _jwt, err := jwt.Decode(c.QueryParam("auth")); err == nil {
				if ok, _ := _jwt.Authenticate(); ok {
					// When authentication is valid
					return handler(c)
				}
			}
			// return echo.NewHTTPError(http.StatusUnauthorized)
			return c.Redirect(http.StatusMovedPermanently, "/login")
		}
	}
}

func LoginHandler(c echo.Context) error {
	if c.QueryParam("user_id") == "2" {
		j, _ := jwt.Publish(map[string]string{"user_id": c.QueryParam("user_id")})
		return c.String(http.StatusOK, j.Encode())
	} else {
		return c.String(http.StatusOK, "Login Failed")
	}
}

func MyPageHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Success!")
}
