package main
import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "xxxxhelado")
	})
	e.GET("/pong", func(c echo.Context) error {
		return c.String(http.StatusOK, "xxxxxjamon")
	})
	e.GET("/bum", func(c echo.Context) error {
		return c.String(http.StatusOK, "xxxxxpizza")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

