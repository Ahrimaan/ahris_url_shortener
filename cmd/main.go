package main

import (
	"log"
	"os"
)

func main() {
	httpPort := os.Getenv("HTTP_PORT")
	conString := os.Getenv("POSTGRE_SQL_CONNECTION")
	log.Printf("Connectinstring is %v", conString)
	InitDB(conString)
	if httpPort == "" {
		httpPort = "80"
	}
	StartServer(httpPort)
}
