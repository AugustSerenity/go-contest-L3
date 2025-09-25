package service

import (
	"github.com/AugustSerenity/go-contest-L3/l3.1/internal/model"
	"github.com/AugustSerenity/go-contest-L3/l3.1/internal/queue/producer"
	"github.com/wb-go/wbf/ginext"
)

type Service struct {
	producer producer.ProducerService
}

func New(p producer.ProducerService) *Service {
	return &Service{
		producer: p,
	}
}

func (s *Service) CreateNotification(ctx *ginext.Context, notification model.Notification) error {
	err := s.producer.Publish(notification)
	if err != nil {
		return err
	}

	return nil
}
