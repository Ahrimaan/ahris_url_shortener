package main

import (
	db "ahris_url_shortener/cmd/database"
	server "ahris_url_shortener/cmd/server"
	"log"
	"os"
)

func main() {
	httpPort := os.Getenv("HTTP_PORT")
	conString := os.Getenv("POSTGRE_SQL_CONNECTION")
	log.Printf("Connectinstring is %v", conString)
	db.InitDB(conString)
	if httpPort == "" {
		httpPort = "80"
	}
	server.StartServer(httpPort)
}
