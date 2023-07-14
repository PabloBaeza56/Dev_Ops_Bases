package main
import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "55555vvvvvxxxxheladommmmmm")
	})
	e.GET("/pong", func(c echo.Context) error {
		return c.String(http.StatusOK, "55555vvvvvxxxxxjamonmmmmm")
	})
	e.GET("/bum", func(c echo.Context) error {
		return c.String(http.StatusOK, "55555vvvvvxxxxxpizzammmmm")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

