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
	return &Service{
		storage: st,
	}
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
