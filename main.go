package main

import (
	"fmt"
	"html/template"
	"io"
	"net"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/plant-healthcheck-server/model"
)

const (
	CONN_HOST = "0.0.0.0"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
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

	initTcpServer(CONN_TYPE, CONN_HOST, CONN_PORT, app)

	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	e := echo.New()
	e.Renderer = t

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", app.HandleHomeTemplate)

	e.Logger.Fatal(e.Start(":1323"))
}

func initTcpServer(cType, host, port string, app *Application) {
	go func(app *Application) {
		// Listen for incoming connections.
		l, err := net.Listen(cType, host+":"+port)
		if err != nil {
			fmt.Println("Error listening:", err.Error())
			os.Exit(1)
		}
		// Close the listener when the application closes.
		defer l.Close()
		fmt.Println("TCP server listening on " + host + ":" + port)
		for {
			// Listen for an incoming connection.
			conn, err := l.Accept()
			if err != nil {
				fmt.Println("Error accepting: ", err.Error())
				os.Exit(1)
			}
			// Handle connections in a new goroutine.
			go app.HandleTCPRequest(conn)
		}
	}(app)
}
