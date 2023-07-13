package main
import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "l6988hhffPOOOOOggggggOOOFFFFFFFFFFFFFNG!")
	})
	e.GET("/pong", func(c echo.Context) error {
		return c.String(http.StatusOK, "l6988hhffPIIIIIggggggIFFFFFFFFFFFFNG!")
	})
	e.GET("/bum", func(c echo.Context) error {
		return c.String(http.StatusOK, "l6988hhffBOOOOggggggOFFFFFFFFFFM!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

