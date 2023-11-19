package handler

import (
	"log"
	"shorter-sites/db"
)

func CountHandler() (int, error) {
	var count int

	err := db.DB.QueryRow("SELECT COUNT(*) FROM url_shortener").Scan(&count)
	if err != nil {
		log.Println("Error retrieving count:", err)
		return -1, err
	}

	return count, nil
}
