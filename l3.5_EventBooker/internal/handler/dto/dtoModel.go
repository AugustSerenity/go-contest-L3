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

type CreateBookingRequest struct {
	Seats int `json:"seats" binding:"required"`
}

type BookingResponse struct {
	ID        int       `json:"id"`
	EventID   int       `json:"event_id"`
	Seats     int       `json:"seats"`
	Paid      bool      `json:"paid"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}
