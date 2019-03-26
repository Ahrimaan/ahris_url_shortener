package server

import (
	"github.com/labstack/echo"
	"net/http"
)

func registerRoutes(e *echo.Echo) {
	e.GET("/:id", getWithId)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "API is up and running !")
	})
	e.POST("/", getShortenedUrl)
}
