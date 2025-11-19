package handler

import (
	"github.com/AugustSerenity/go-contest-L3/l3.6_SalesTracker/internal/model"
	"github.com/wb-go/wbf/ginext"
)

type Service interface {
	CreateItem(*ginext.Context, model.Item) (model.Item, error)
}
