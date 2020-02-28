package main

import (
	"net/http"
	"os"

	"github.com/polaris1119/go-echo-example/pkg/logger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()

	e.Logger = logger.New(os.Stdout)
	e.Logger.SetLevel(log.INFO)

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.GET("/", func(ctx echo.Context) error {
		zerolog := ctx.Logger().(*logger.Logger).ZeroLog
		zerolog.Debug().Str("path", ctx.Path()).Msg("This is Debug msg!")

		return ctx.HTML(http.StatusOK, "Hello World!")
	})

	e.Logger.Fatal(e.Start(":2020"))
}
