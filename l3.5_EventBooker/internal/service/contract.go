package service

import (
	"context"

	"github.com/AugustSerenity/go-contest-L3/l3.5_EventBooker/internal/model"
)

type Storage interface {
	CreateEvent(ctx context.Context, event *model.Event) error
}
