package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func helloWorld(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, World!")
}

func ApiHandler(e *echo.Echo) {
	api := e.Group("/api")

	api.GET("/", helloWorld)

	api.GET("/cat", func(c echo.Context) error {
		cat, err := GetCat(c)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, cat)
	})
}
