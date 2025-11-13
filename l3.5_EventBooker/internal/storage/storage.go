package storage

import (
	"context"
	"time"

	"github.com/AugustSerenity/go-contest-L3/l3.5_EventBooker/internal/model"
	"github.com/wb-go/wbf/dbpg"
)

type Storage struct {
	db *dbpg.DB
}

func New(db *dbpg.DB) *Storage {
	return &Storage{db: db}
}

func (st *Storage) CreateEvent(ctx context.Context, event *model.Event) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `
		INSERT INTO events (name, date, capacity, free_seats, payment_ttl, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id;
	`

	row := st.db.Master.QueryRowContext(ctx, query,
		event.Name, event.Date, event.Capacity, event.FreeSeats,
		event.PaymentTTL, event.CreatedAt, event.UpdatedAt,
	)

	if err := row.Scan(&event.ID); err != nil {
		return err
	}

	return nil
}
