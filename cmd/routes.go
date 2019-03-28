package main

import (
	"github.com/labstack/echo"
)

func registerRoutes(e *echo.Echo) {
	e.GET("/api/:id", GetWithId)
	e.POST("/api", GetShortenedUrl)
}
