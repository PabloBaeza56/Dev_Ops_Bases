package main
import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "zzzzzzzz00000555555xxxxxxxxxxxxxbbbbb")
	})
	e.GET("/pong", func(c echo.Context) error {
		return c.String(http.StatusOK, "zzzzzzz0000000555555xxxxxxxxxxvvvvv")
	})
	e.GET("/bum", func(c echo.Context) error {
		return c.String(http.StatusOK, "zzzzzz00000005555555555xxxxxxxxxxxiiiii")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

