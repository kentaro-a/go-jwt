package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func main() {
	fmt.Println("start")
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/mypage", MyPageHandler)
	e.GET("/login", LoginHandler)
	e.Start(":9999")
}

func LoginHandler(c echo.Context) error {
	if c.QueryParam("user_id") == "2" {
		return c.JSONPretty(http.StatusOK, map[string]string{"user_id": "1", "expire_at": "2019-05-01"}, "\t")
	} else {
		return c.JSONPretty(http.StatusOK, nil, "\t")
	}
}

func MyPageHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Success!")
}
