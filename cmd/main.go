package main

import (
	"ahris_url_shortener/cmd/data"
	"ahris_url_shortener/cmd/server"
	"log"
	"os"
)

func main() {
	conString := os.Getenv("POSTGRE_SQL_CONNECTION")
	log.Printf("Connectinstring is %v", conString)
	data.InitDB(conString)
	server.StartServer("80")
}
