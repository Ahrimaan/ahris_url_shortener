package main

import (
	"database/sql"
	"log"
	"time"
)

func SetNewRecord(url string) (id string, err error) {
	var newId string
	done := false
	for done == false {
		newId = GetRandomId(5)
		tempId := ""
		row := db.QueryRow("SELECT id from public.urls where id = $1", newId)
		switch err := row.Scan(tempId); err {
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
	err = db.QueryRow(sqlStatement, newId, url, time.Now()).Scan(&newId)
	return newId, err
}

func GetUrl(id string) (string, error) {
	row := db.QueryRow("SELECT url from public.urls where id = $1", id)
	var returnUrl string
	switch err := row.Scan(&returnUrl); err {
	case sql.ErrNoRows:
		log.Printf("No Rows with id $1", id)
		return "", nil
	case nil:
		return returnUrl, nil
	default:
		panic(err)
	}
}
