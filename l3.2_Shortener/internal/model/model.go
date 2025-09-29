package model

import "time"

type URL struct {
	OriginalURL string    `json:"oroginal_url"`
	ShortURL    string    `json:"short_url"`
	CreateAt    time.Time `json:"create_at"`
}

