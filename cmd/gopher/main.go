package main

import (
	"net/http"
	"os"

	"github.com/polaris1119/go-echo-example/pkg/logger"
	"github.com/polaris1119/go-echo-example/pkg/render"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

// type Template struct {
// 	templates *template.Template
// }

// func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
// 	return t.templates.ExecuteTemplate(w, name, data)
// }

func main() {
	e := echo.New()

	e.Logger = logger.New(os.Stdout)
	e.Logger.SetLevel(log.DEBUG)

	// tpl := &Template{
	// 	templates: template.Must(template.ParseGlob("template/*.html")),
	// }
	// e.Renderer = t

	e.Renderer = render.LayoutTemplate

	e.Use(middleware.Recover())

	e.Static("/static", "static")

	e.GET("/", func(ctx echo.Context) error {
		return ctx.Render(http.StatusOK, "index.html", nil)
	})

	e.Logger.Fatal(e.Start(":2020"))
}
