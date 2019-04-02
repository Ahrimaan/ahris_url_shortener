package server

import (
	"github.com/labstack/echo"
)

func registerRoutes(e *echo.Echo) {
	e.GET("/api/:id", GetWithID)
	e.POST("/api", GetShortenedURL)
}
