package model

import "time"

type Event struct {
	ID         int           `db:"id" json:"id"`
	Name       string        `db:"name" json:"name"`
	Date       time.Time     `db:"date" json:"date"`
	Capacity   int           `db:"capacity" json:"capacity"`
	FreeSeats  int           `db:"free_seats" json:"free_seats"`
	PaymentTTL time.Duration `db:"payment_ttl" json:"payment_ttl"`
	CreatedAt  time.Time     `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time     `db:"updated_at" json:"updated_at"`
}

type Booking struct {
	ID        int       `json:"id"`
	EventID   int       `json:"event_id"`
	UserID    int       `json:"user_id"`
	Seats     int       `json:"seats"`
	Paid      bool      `json:"paid"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

