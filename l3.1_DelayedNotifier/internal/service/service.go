package service

import (
	"fmt"

	"github.com/AugustSerenity/go-contest-L3/l3.1/internal/model"
	"github.com/AugustSerenity/go-contest-L3/l3.1/internal/queue/producer"
	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/zlog"
)

type Service struct {
	producer *producer.Producer
	storage  Storage
}

func New(p *producer.Producer, st Storage) *Service {
	return &Service{
		producer: p,
		storage:  st,
	}
}

func (s *Service) CreateNotification(ctx *ginext.Context, notification model.Notification) error {
	s.storage.Set(notification)
	err := s.producer.Publish(notification)
	if err != nil {
		zlog.Logger.Error().
			Err(err).
			Str("notification_id", notification.ID).
			Msg("failed to publish notification to queue")
		return err
	}
	return nil
}

func (s *Service) ProcessNotification(notification model.Notification) {
	notification.Status = "processed"
	s.storage.Set(notification)
}

func (s *Service) GetStatusByID(ctx *ginext.Context, id string) (model.Notification, error) {
	res, ok := s.storage.Get(id)
	if !ok {
		return model.Notification{}, fmt.Errorf("id not found")
	}

	return res, nil
}
