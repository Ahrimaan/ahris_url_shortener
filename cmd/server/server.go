package server

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func StartServer(port string) {
	e := echo.New()
	registerRoutes(e)

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","id":"${id}","path":"${path}" ,"remote_ip":"${remote_ip}","host":"${host}",` +
			`"method":"${method}","uri":"${uri}","status":${status},"error":"${error}","latency":${latency},` +
			`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
			`"bytes_out":${bytes_out}}` + "\n",
	}))
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		fmt.Printf("%s\n", resBody)
	}))
	e.Logger.Fatal(e.Start(":" + port))
}
