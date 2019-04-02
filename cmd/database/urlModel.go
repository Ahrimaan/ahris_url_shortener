package database

import (
	"database/sql"
	"log"
	"time"
)

//SetNewRecord Writes a new Record to the Database
func SetNewRecord(url string) (id string, err error) {
	var newID string
	done := false
	for done == false {
		newID = GetRandomID(5)
		tempID := ""
		row := db.QueryRow("SELECT id from public.urls where id = $1", newID)
		switch err := row.Scan(tempID); err {
		case nil:
		case sql.ErrNoRows:
			done = true
		default:
			done = true
		}
	}
	sqlStatement := `
INSERT INTO public.urls (id,url,timestamp)
VALUES ($1, $2, $3)
RETURNING id`
	err = db.QueryRow(sqlStatement, newID, url, time.Now()).Scan(&newID)
	return newID, err
}

//GetURL Returns the URL specified bye the ID
func GetURL(id string) (string, error) {
	row := db.QueryRow("SELECT url from public.urls where id = $1", id)
	var returnURL string
	switch err := row.Scan(&returnURL); err {
	case sql.ErrNoRows:
		log.Printf("No Rows with id %v", id)
		return "", nil
	case nil:
		return returnURL, nil
	default:
		panic(err)
	}
}
