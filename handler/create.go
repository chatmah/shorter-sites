package handler

import (
	"log"
	"net/http"
	"shorter-sites/db"
	"shorter-sites/model"
	"time"
)

func CreateHandler(w http.ResponseWriter, r *http.Request, obj model.UrlData) error {
	obj.CreationDate = time.Now()
	obj.ExpirationDate = time.Now().Add(30 * 24 * time.Hour)

	_, err := db.DB.Exec("INSERT INTO url_shortener (long_url, short_code, creation_date, expiration_date) VALUES ($1, $2, $3, $4)",
		obj.LongURL, obj.ShortCode, obj.CreationDate, obj.ExpirationDate)

	if err != nil {
		log.Println("Error inserting URL shortener:", err)
		return err
	}

	return nil
}
