package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "POoofffNG!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
