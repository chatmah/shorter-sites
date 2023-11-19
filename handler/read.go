package handler

import (
	"shorter-sites/db"
	"shorter-sites/model"
)

func ReadHandler(encodedLink string) (*model.UrlData, error) {

	// err := db.DB.QueryRow("SELECT COUNT(*) FROM url_shortener").Scan(&count)

	query := "SELECT id, long_url, short_code FROM url_shortener WHERE short_code = $1"
	row := db.DB.QueryRow(query, encodedLink)

	var url model.UrlData
	err := row.Scan(&url.ID, &url.LongURL, &url.ShortCode)
	if err != nil {
		return nil, err
	}

	return &url, nil
}
