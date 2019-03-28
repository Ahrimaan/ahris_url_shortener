package main

import (
	"log"
	"os"
)

func main() {
	conString := os.Getenv("POSTGRE_SQL_CONNECTION")
	log.Printf("Connectinstring is %v", conString)
	InitDB(conString)
	StartServer("80")
}
