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
	e.Logger.SetLevel(log.DEBUG)

	e.Use(middleware.Recover())

	e.GET("/", func(ctx echo.Context) error {
		ctx.Logger().Debugf("This is echo logger debug msg!")

		zerolog := ctx.Logger().(*logger.Logger).ZeroLog
		zerolog.Debug().Str("path", ctx.Path()).Msg("This is zerolog Debug msg!")

		// return ctx.HTML(http.StatusOK, "Hello World!")
		return ctx.Render(http.StatusOK, "")
	})

	e.Logger.Fatal(e.Start(":2020"))
}
