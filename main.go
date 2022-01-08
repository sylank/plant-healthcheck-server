package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/plant-healthcheck-server/model"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type Application struct {
	History model.History
}

func main() {
	app := &Application{
		History: *model.CreateHistory(5000),
	}

	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	e := echo.New()
	e.Renderer = t

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", app.HandleHomeTemplate)
	e.POST("/insert", app.InsertSensorData)

	e.Logger.Fatal(e.Start(":3000"))
}
