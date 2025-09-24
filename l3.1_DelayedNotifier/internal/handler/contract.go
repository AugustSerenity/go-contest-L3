package handler

import (
	"github.com/AugustSerenity/go-contest-L3/l3.1/internal/model"
	"github.com/wb-go/wbf/ginext"
)

type Service interface {
	CreateNotification(*ginext.Context, model.Notification) error
}
