package main

import (
	"net/http"

	echo "github.com/labstack/echo"
)

func main() {
	// instance echo
	e := echo.New()

	// path
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
