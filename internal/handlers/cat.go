package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/untitled-discord-star-project/backend/internal/db"
)

func GetCat(ctx echo.Context) ([]db.Cat, error) {
	cats, err := db.Select()
	if err != nil {
		return nil, err
	}

	return cats, nil
}
