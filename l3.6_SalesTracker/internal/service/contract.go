package service

import (
	"github.com/AugustSerenity/go-contest-L3/l3.6_SalesTracker/internal/model"
	"github.com/wb-go/wbf/ginext"
)

type Storage interface {
	SaveItem(*ginext.Context, model.Item) (model.Item, error)
	AnalyticsCalculate(c *ginext.Context, filter model.ItemsFilter) (model.AnalyticsResponse, error)
}
