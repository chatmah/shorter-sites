package model

import "time"

type UrlData struct {
	ID             int       `json:"id"`
	LongURL        string    `json:"long_url"`
	ShortCode      string    `json:"short_code"`
	CreationDate   time.Time `json:"creation_date"`
	ExpirationDate time.Time `json:"expiration_date"`
}
