package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/untitled-discord-star-project/backend/internal/handlers"
	"github.com/untitled-discord-star-project/backend/pkg/utils"
	"github.com/untitled-discord-star-project/backend/templates"
)

func main() {
	e := echo.New()

	e.GET("/", func(ctx echo.Context) error {
		return utils.RenderComponent(ctx, http.StatusOK, templates.Index("2so cool!!"))
	})

	e.Static("/static", "static")

	handlers.ApiHandler(e)

	e.Logger.Fatal(e.Start(":1323"))
}
