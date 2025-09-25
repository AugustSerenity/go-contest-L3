package service

import (
	"github.com/AugustSerenity/go-contest-L3/l3.1/internal/model"
	"github.com/AugustSerenity/go-contest-L3/l3.1/internal/queue/producer"
	"github.com/wb-go/wbf/ginext"
)

type Service struct {
	producer producer.ProducerService
	storage  Storage
}

func New(p producer.ProducerService, st Storage) *Service {
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
