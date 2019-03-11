package server

import (
	"github.com/labstack/echo"
	"net/http"
)

func handleGetUrl(c echo.Context) error {
	var urlId = c.Param("id")
	return c.String(http.StatusOK, urlId)
}
