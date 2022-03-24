package webstuff

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HelloWebCall() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

func IDWebcall() {
	e := echo.New()
	e.GET("/users/:id", func(c echo.Context) error {
		// User ID from path `users/:id`
		id := c.Param("id")
		return c.String(http.StatusOK, "Hello "+id)
	})
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "server running")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
