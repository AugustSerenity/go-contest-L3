package service

import (
	"time"

	"github.com/AugustSerenity/go-contest-L3/l3.6_SalesTracker/internal/model"
	"github.com/wb-go/wbf/ginext"
)

type Service struct {
	storage Storage
}

func New(st Storage) *Service {
	return &Service{storage: st}
}

func (s *Service) CreateItem(c *ginext.Context, item model.Item) (model.Item, error) {
	item.CreatedAt = time.Now()

	data, err := s.storage.SaveItem(c, item)
	if err != nil {
		return model.Item{}, err
	}

	return data, nil
}

func (s *Service) GetAnalytics(c *ginext.Context, filter model.ItemsFilter) (model.AnalyticsResponse, error) {
	return s.storage.AnalyticsCalculate(c, filter)
}
