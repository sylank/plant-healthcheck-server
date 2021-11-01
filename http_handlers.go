package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/plant-healthcheck-server/model"
)

// ...
type HomeDto struct {
	Current model.HistoryElement
	History []model.HistoryElement
}

func (app *Application) HandleHomeTemplate(c echo.Context) error {
	dto := HomeDto{
		Current: app.History.GetLatest(),
		History: app.History.GetElements(),
	}

	return c.Render(http.StatusOK, "hello", dto)
}
