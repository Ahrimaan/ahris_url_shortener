package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func StartServer(port string) {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{os.Getenv("ALLOW_ORIGIN")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPost},
	}))
	registerRoutes(e)
	if os.Getenv("DEBUG") == "true" {
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: `{"time":"${time_rfc3339_nano}","id":"${id}","path":"${path}" ,"remote_ip":"${remote_ip}","host":"${host}",` +
				`"method":"${method}","uri":"${uri}","status":${status},"error":"${error}","latency":${latency},` +
				`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
				`"bytes_out":${bytes_out}}` + "\n",
		}))
		e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
			fmt.Printf("%s\n", resBody)
		}))
	}

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "./public",
		Browse: true,
		Index:  "index.html",
	}))
	e.Logger.Fatal(e.Start(":" + port))
}
