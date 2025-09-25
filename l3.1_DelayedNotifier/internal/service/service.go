package service

import (
	"github.com/AugustSerenity/go-contest-L3/l3.1/internal/model"
	"github.com/AugustSerenity/go-contest-L3/l3.1/internal/queue/producer"
	"github.com/wb-go/wbf/ginext"
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
	err := s.producer.Publish(notification)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) ProcessNotification(notification model.Notification) {
	notification.Status = "processed"
	s.storage.Set(notification)
}
