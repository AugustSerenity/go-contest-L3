package dto

import "time"

type CreateEventRequest struct {
	Name       string        `json:"name" binding:"required"`
	Date       time.Time     `json:"date"`
	Capacity   int           `json:"capacity"`
	PaymentTTL time.Duration `json:"payment_ttl"`
}

type EventResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Date      time.Time `json:"date"`
	Capacity  int       `json:"capacity"`
	FreeSeats int       `json:"free_seats"`
}
