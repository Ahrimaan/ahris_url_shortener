package server

import (
	"ahris_url_shortener/cmd/data"
	"github.com/labstack/echo"
	"net/http"
	"os"
)

type UrlModel struct {
	Url string `json:"url"`
}

func GetWithId(c echo.Context) error {
	var urlId = c.Param("id")
	req := new(UrlModel)
	err := c.Bind(req)
	if err != nil {
		return c.String(http.StatusBadRequest, "The Model you have given is not correct: "+err.Error())
	}
	url, err := data.GetUrl(urlId)
	if err != nil {
		return c.String(http.StatusNotFound, "Could not find data")
	}
	if url == "" {
		return c.String(http.StatusNotFound, "Could not find data")
	}

	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func GetShortenedUrl(c echo.Context) error {

	req := new(UrlModel)
	err := c.Bind(req)
	if err != nil {
		return c.String(http.StatusBadRequest, "The Model you have given is not correct: "+err.Error())
	}

	newId, err := data.SetNewRecord(req.Url)
	if err != nil {
		return c.String(http.StatusInternalServerError, "OOOPS Something went wrong, sorry for that ")
	}
	newUri := getCurrentUrl(c, newId)
	resp := &UrlModel{
		Url: newUri,
	}

	return c.JSON(http.StatusOK, resp)
}

func getCurrentUrl(c echo.Context, id string) string {
	currentURL := os.Getenv("API_URL")
	if currentURL == "" {
		r := c.Request()
		currentURL = c.Scheme() + "://" + r.Host + "/api/" + id
	}
	return currentURL
}
