package storage

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/AugustSerenity/go-contest-L3/l3.5_EventBooker/internal/model"
	"github.com/wb-go/wbf/dbpg"
	"github.com/wb-go/wbf/zlog"
)

type Storage struct {
	db *dbpg.DB
}

func New(db *dbpg.DB) *Storage {
	return &Storage{db: db}
}

func (st *Storage) CreateEvent(ctx context.Context, event *model.Event) error {
	tx, err := st.db.Master.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `
		INSERT INTO events (name, date, capacity, free_seats, payment_ttl, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id;
	`

	row := tx.QueryRowContext(ctx, query,
		event.Name, event.Date, event.Capacity, event.FreeSeats,
		fmt.Sprintf("%ds", int(event.PaymentTTL.Seconds())),
		event.CreatedAt, event.UpdatedAt,
	)

	if err := row.Scan(&event.ID); err != nil {
		return err
	}

	return tx.Commit()
}

func (st *Storage) BookEvent(ctx context.Context, eventID, seats int, paymentTTL time.Duration) (*model.Booking, error) {
	tx, err := st.db.Master.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var freeSeats int
	err = tx.QueryRowContext(ctx, "SELECT free_seats FROM events WHERE id=$1 FOR UPDATE", eventID).Scan(&freeSeats)
	if err != nil {
		return nil, err
	}
	if freeSeats < seats {
		return nil, fmt.Errorf("not enough free seats")
	}

	now := time.Now()
	booking := &model.Booking{
		EventID:   eventID,
		Seats:     seats,
		Paid:      false,
		CreatedAt: now,
		ExpiresAt: now.Add(paymentTTL),
	}

	query := `INSERT INTO bookings (event_id, seats, paid, created_at, expires_at)
			  VALUES ($1,$2,$3,$4,$5) RETURNING id`
	if err := tx.QueryRowContext(ctx, query, booking.EventID, booking.Seats, booking.Paid, booking.CreatedAt, booking.ExpiresAt).Scan(&booking.ID); err != nil {
		return nil, err
	}

	_, err = tx.ExecContext(ctx, "UPDATE events SET free_seats = free_seats - $1, updated_at=now() WHERE id = $2", seats, eventID)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return booking, nil
}

func (st *Storage) ConfirmBooking(ctx context.Context, bookingID int) error {
	tx, err := st.db.Master.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, "UPDATE bookings SET paid=true WHERE id=$1", bookingID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (st *Storage) CancelBooking(ctx context.Context, bookingID int) error {
	tx, err := st.db.Master.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var eventID, seats int
	err = tx.QueryRowContext(ctx, "SELECT event_id, seats FROM bookings WHERE id=$1 AND paid=false", bookingID).Scan(&eventID, &seats)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, "DELETE FROM bookings WHERE id=$1", bookingID)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, "UPDATE events SET free_seats = free_seats + $1, updated_at=now() WHERE id=$2", seats, eventID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (st *Storage) GetEvent(ctx context.Context, eventID int) (*model.Event, error) {
	var event model.Event
	err := st.db.Master.QueryRowContext(ctx, "SELECT id, name, date, capacity, free_seats, payment_ttl, created_at, updated_at FROM events WHERE id=$1", eventID).
		Scan(&event.ID, &event.Name, &event.Date, &event.Capacity, &event.FreeSeats, &event.PaymentTTL, &event.CreatedAt, &event.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (st *Storage) CancelExpiredBookings(ctx context.Context) error {
	tx, err := st.db.Master.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(ctx, `
		SELECT id 
		FROM bookings 
		WHERE paid=false AND expires_at <= now()
	`)
	if err != nil {
		return err
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return err
		}
		ids = append(ids, id)
	}

	for _, id := range ids {
		if err := st.CancelBooking(ctx, id); err != nil {
			zlog.Logger.Error().Err(err).Int("booking_id", id).Msg("failed to cancel expired booking")
		}
	}

	return tx.Commit()
}

func (st *Storage) GetEvents(ctx context.Context) ([]model.Event, error) {
	rows, err := st.db.Master.QueryContext(ctx,
		`SELECT id, name, capacity, free_seats FROM events ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []model.Event
	for rows.Next() {
		var e model.Event
		if err := rows.Scan(&e.ID, &e.Name, &e.Capacity, &e.FreeSeats); err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	return events, nil
}

func (st *Storage) GetEventBookings(ctx context.Context, eventID int) ([]model.Booking, error) {
	rows, err := st.db.Master.QueryContext(ctx,
		`SELECT id, event_id, seats, paid, created_at, expires_at 
		 FROM bookings WHERE event_id=$1 ORDER BY id`, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.Booking
	for rows.Next() {
		var b model.Booking
		if err := rows.Scan(&b.ID, &b.EventID, &b.Seats, &b.Paid, &b.CreatedAt, &b.ExpiresAt); err != nil {
			return nil, err
		}
		list = append(list, b)
	}
	return list, nil
}
