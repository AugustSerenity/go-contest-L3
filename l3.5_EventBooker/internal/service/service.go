package service

import (
	"context"
	"time"

	"github.com/AugustSerenity/go-contest-L3/l3.5_EventBooker/internal/model"
)

type Service struct {
	storage Storage
}

func New(st Storage) *Service {
	return &Service{storage: st}
}

func (s *Service) CreateEvent(ctx context.Context, event *model.Event) error {
	event.FreeSeats = event.Capacity
	event.CreatedAt = time.Now()
	event.UpdatedAt = event.CreatedAt

	if err := s.storage.CreateEvent(ctx, event); err != nil {
		return err
	}
	return nil
}

func (s *Service) BookEvent(eventID, seats int) (*model.Booking, error) {
	event, err := s.GetEvent(eventID)
	if err != nil {
		return nil, err
	}

	return s.storage.BookEvent(context.Background(), eventID, seats, event.PaymentTTL)
}

func (s *Service) ConfirmBooking(bookingID int) error {
	return s.storage.ConfirmBooking(context.Background(), bookingID)
}

func (s *Service) GetEvent(eventID int) (*model.Event, error) {
	return s.storage.GetEvent(context.Background(), eventID)
}

func (s *Service) CancelExpiredBookings() error {
	return s.storage.CancelExpiredBookings(context.Background())
}

func (s *Service) GetEvents() ([]model.Event, error) {
	return s.storage.GetEvents(context.Background())
}

func (s *Service) GetEventBookings(eventID int) ([]model.Booking, error) {
	return s.storage.GetEventBookings(context.Background(), eventID)
}
