package model

import "time"

type URL struct {
	OriginalURL string    `json:"original_url"`
	ShortURL    string    `json:"short_url"`
	CreateAt    time.Time `json:"create_at"`
}
