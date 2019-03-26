package server

import (
	"github.com/labstack/echo"
	"net/http"
)

type Response struct {
	Url          string `json:"url"`
	ShortenedUrl string `json:"shortenedUrl"`
}

type Request struct {
	Url string `json:"url"`
}

func getWithId(c echo.Context) error {
	var urlId = c.Param("id")
	return c.String(http.StatusOK, urlId)
}

func getShortenedUrl(c echo.Context) error {
	req := new(Request)
	err := c.Bind(req)
	if err != nil {
		return c.String(http.StatusBadRequest, "The Model you have given is not correct: "+err.Error())
	}
	resp := &Response{
		Url:          req.Url,
		ShortenedUrl: "newURL",
	}

	return c.JSON(http.StatusOK, resp)
}
