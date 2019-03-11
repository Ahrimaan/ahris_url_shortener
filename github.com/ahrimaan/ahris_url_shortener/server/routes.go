package server

import (
	"github.com/labstack/echo"
	"net/http"
)

func registerRoutes(e *echo.Echo) {
	e.GET("/:id", handleGetUrl)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}
