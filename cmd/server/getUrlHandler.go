package server

import (
	"net/http"
	"os"

	db "ahris_url_shortener/cmd/database"

	"github.com/labstack/echo"
)

//URLModel the Model
type URLModel struct {
	URL string `json:"url"`
}

//GetWithID Returns the Model specified by the ID
func GetWithID(c echo.Context) error {
	var urlID = c.Param("id")
	req := new(URLModel)
	err := c.Bind(req)
	if err != nil {
		return c.String(http.StatusBadRequest, "The Model you have given is not correct: "+err.Error())
	}
	url, err := db.GetURL(urlID)
	if err != nil {
		return c.String(http.StatusNotFound, "Could not find data")
	}
	if url == "" {
		return c.String(http.StatusNotFound, "Could not find data")
	}

	return c.Redirect(http.StatusTemporaryRedirect, url)
}

//GetShortenedURL Returns the shortened URL
func GetShortenedURL(c echo.Context) error {

	req := new(URLModel)
	err := c.Bind(req)
	if err != nil {
		return c.String(http.StatusBadRequest, "The Model you have given is not correct: "+err.Error())
	}

	newID, err := db.SetNewRecord(req.URL)
	if err != nil {
		return c.String(http.StatusInternalServerError, "OOOPS Something went wrong, sorry for that ")
	}
	newURI := getCurrentURL(c, newID)
	resp := &URLModel{
		URL: newURI,
	}

	return c.JSON(http.StatusOK, resp)
}

func getCurrentURL(c echo.Context, id string) string {
	currentURL := os.Getenv("API_URL")
	if currentURL == "" {
		r := c.Request()
		currentURL = c.Scheme() + "://" + r.Host + "/api/" + id
	}
	return currentURL
}
