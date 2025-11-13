package handler

import "github.com/AugustSerenity/go-contest-L3/l3.5_EventBooker/internal/model"

type Service interface {
	CreateEvent(model.Event) error
}
