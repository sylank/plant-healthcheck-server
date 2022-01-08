package main

import (
	"net/http"
	"time"

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

func (app *Application) InsertSensorData(c echo.Context) error {
	sensorData := new(model.SensorData)

	if err := c.Bind(sensorData); err != nil {
		return err
	}

	element := model.HistoryElement{
		DateStr:      time.Now().Format("2006-01-02 15:04:05"),
		Temperature:  sensorData.Temperature,
		Humidity:     sensorData.Humidity,
		SoilMoisture: sensorData.SoilMoisture,
	}

	app.History.Push(element)

	return c.JSONBlob(
		http.StatusOK,
		[]byte(
			"inserted",
		),
	)
}
